package worker

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/sourcegraph/sourcegraph/enterprise/cmd/executor/internal/command"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

const SchemeExecutorToken = "token-executor"

// prepareWorkspace creates and returns a temporary director in which acts the workspace
// while processing a single job. It is up to the caller to ensure that this directory is
// removed after the job has finished processing. If a repository name is supplied, then
// that repository will be cloned (through the frontend API) into the workspace.
func (h *handler) prepareWorkspace(ctx context.Context, commandRunner command.Runner, repositoryName, commit, sparseCheckout string, shallowClone bool) (_ string, err error) {
	tempDir, err := makeTempDir()
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = os.RemoveAll(tempDir)
		}
	}()

	// // Store the repo in a subdirectory. Why? We want to be able to put other files in there, too and they shouldn't interfere.
	// repoPath := filepath.Join(tempDir, "repo")

	// if err := os.Mkdir(repoPath, os.ModePerm); err != nil {
	// 	return "", err
	// }

	if repositoryName != "" {
		cloneURL, err := makeRelativeURL(
			h.options.ClientOptions.EndpointOptions.URL,
			h.options.GitServicePath,
			repositoryName,
		)
		if err != nil {
			return "", err
		}

		authorizationOption := fmt.Sprintf(
			"http.extraHeader=Authorization: %s %s",
			SchemeExecutorToken,
			h.options.ClientOptions.EndpointOptions.Token,
		)

		// TODO: Can we do something here to improve caching?
		// Like keeping the same tmp dir for the .git but checking out into multiple workspaces?
		// This might help with performance on configurations with fewer executors,
		// but maybe it doesn't matter beyond a certain point.

		fetchCommand := []string{"git", "-C", tempDir, "-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal", "fetch", "--no-tags", "--prune", "--progress", "--no-recurse-submodules",
			"--depth=1", "origin", commit}

		if sparseCheckout != "" {
			fetchCommand = append(fetchCommand[:14], fetchCommand[13:]...)
			fetchCommand[13] = "--filter=blob:none"
		}

		gitStdEnv := []string{"GIT_TERMINAL_PROMPT=0", "GCM_INTERACTIVE=Never"}

		gitCommands := []command.CommandSpec{
			// First, mark the repo as a safe directory. This will allow git commands to be run in the repo even if the UID changes.
			// TODO: This will clutter the global config.
			// TODO: This just failed in local dev with `error: could not lock config file /Users/erik/.gitconfig: File exists`.
			{Key: "setup.git.safe-directory", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "config", "--global", "safe.directory", tempDir}, Operation: h.operations.SetupGitInit},
			{Key: "setup.git.init", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "init"}, Operation: h.operations.SetupGitInit},
			{Key: "setup.git.add-remote", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "remote", "add", "origin", cloneURL.String()}, Operation: h.operations.SetupAddRemote},
			// Disable gc, this can improve performance.
			{Key: "setup.git.disable-gc", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "config", "--local", "gc.auto", "0"}, Operation: h.operations.SetupAddRemote},
			{Key: "setup.git.fetch", Env: gitStdEnv, Command: fetchCommand, Operation: h.operations.SetupGitFetch},
		}

		if sparseCheckout != "" {
			// TODO: Is this required?
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.sparse-checkout-config", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "config", "--local", "core.sparseCheckoutCone", "1"}, Operation: h.operations.SetupGitCheckout})
			// TODO: `set` or `init`?
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.sparse-checkout-init", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "sparse-checkout", "set", "--cone"}, Operation: h.operations.SetupGitCheckout})
		}

		gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.checkout", Env: gitStdEnv, Command: []string{"git", "-C", tempDir,
			// TODO: This is only required for sparse checkouts:
			"-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal",
			"checkout", "--progress", "--force",
			// TODO: the commit needs to be `+commit:ref`, instead of just commit. Maybe we don't need a proper branch though.
			//  "-B",
			commit}, Operation: h.operations.SetupGitCheckout})
		if sparseCheckout != "" {
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.sparse-checkout-set", Env: gitStdEnv, Command: []string{"git", "-C", tempDir, "-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal", "sparse-checkout", "set", "--", sparseCheckout}, Operation: h.operations.SetupGitCheckout})
		}

		// TODO: Remove, this just validates the files are there.
		gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.ls-files", Env: gitStdEnv, Command: []string{"du", "-ahc", tempDir}, Operation: h.operations.SetupGitFetch})

		for _, spec := range gitCommands {
			if err := commandRunner.Run(ctx, spec); err != nil {
				return "", errors.Wrap(err, fmt.Sprintf("failed %s", spec.Key))
			}
		}
	}

	if err := os.MkdirAll(filepath.Join(tempDir, command.ScriptsPath), os.ModePerm); err != nil {
		return "", err
	}

	return tempDir, nil
}

func makeRelativeURL(base string, path ...string) (*url.URL, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	urlx, err := baseURL.ResolveReference(&url.URL{Path: filepath.Join(path...)}), nil
	if err != nil {
		return nil, err
	}

	urlx.User = url.User("executor")
	return urlx, nil
}

// makeTempDir defaults to makeTemporaryDirectory and can be replaced for testing
// with determinstic workspace/scripts directories.
var makeTempDir = makeTemporaryDirectory

func makeTemporaryDirectory() (string, error) {
	// TMPDIR is set in the dev Procfile to avoid requiring developers to explicitly
	// allow bind mounts of the host's /tmp. If this directory doesn't exist,
	// os.MkdirTemp below will fail.
	if tempdir := os.Getenv("TMPDIR"); tempdir != "" {
		if err := os.MkdirAll(tempdir, os.ModePerm); err != nil {
			return "", err
		}
		return os.MkdirTemp(tempdir, "")
	}

	return os.MkdirTemp("", "")
}
