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
func (h *handler) prepareWorkspace(ctx context.Context, commandRunner command.Runner, repositoryName, commit string, shallowClone bool, sparseCheckout []string) (_, _ string, err error) {
	tempDir, err := makeTempDir()
	if err != nil {
		return "", "", err
	}
	defer func() {
		if err != nil {
			_ = os.RemoveAll(tempDir)
		}
	}()

	// Store the repo in a subdirectory. We don't want to clutter the repo with our scripts and additional files
	// that may be mounted.
	// TODO: This doesn't need to happen when no repo cloning is requested.
	repoPath := filepath.Join(tempDir, "repo")

	if err := os.Mkdir(repoPath, os.ModePerm); err != nil {
		return "", "", err
	}

	if repositoryName != "" {
		cloneURL, err := makeRelativeURL(
			h.options.ClientOptions.EndpointOptions.URL,
			h.options.GitServicePath,
			repositoryName,
		)
		if err != nil {
			return "", "", err
		}

		// These env vars should be set for git commands. We want to make sure it never hangs on interactive input.
		gitStdEnv := []string{"GIT_TERMINAL_PROMPT=0"}

		authorizationOption := fmt.Sprintf(
			"http.extraHeader=Authorization: %s %s",
			SchemeExecutorToken,
			h.options.ClientOptions.EndpointOptions.Token,
		)

		fetchCommand := []string{
			"git",
			"-C", repoPath,
			"-c", "protocol.version=2",
			"-c", authorizationOption,
			"-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal",
			"fetch",
			"--prune",
			"--progress",
			"--no-recurse-submodules",
			"origin",
			commit,
		}

		if shallowClone {
			fetchCommand = append(fetchCommand[:11], fetchCommand[10:]...)
			fetchCommand[10] = "--no-tags"
			fetchCommand = append(fetchCommand[:11], fetchCommand[10:]...)
			fetchCommand[10] = "--depth=1"
		}

		// For a sparse checkout, we want to add a blob filter so we only fetch the minimum set of files initially.
		if len(sparseCheckout) > 0 {
			fetchCommand = append(fetchCommand[:11], fetchCommand[10:]...)
			fetchCommand[10] = "--filter=blob:none"
		}

		gitCommands := []command.CommandSpec{
			// First, mark the repo as a safe directory. This will allow git commands to be run in the repo even if the UID changes.
			// TODO: This will clutter the global config.
			// TODO: This just failed in local dev with `error: could not lock config file /Users/erik/.gitconfig: File exists`.
			{Key: "setup.git.safe-directory", Env: gitStdEnv, Command: []string{"git", "-C", repoPath, "config", "--global", "safe.directory", repoPath}, Operation: h.operations.SetupGitSafeDirectory},
			{Key: "setup.git.init", Env: gitStdEnv, Command: []string{"git", "-C", repoPath, "init"}, Operation: h.operations.SetupGitInit},
			{Key: "setup.git.add-remote", Env: gitStdEnv, Command: []string{"git", "-C", repoPath, "remote", "add", "origin", cloneURL.String()}, Operation: h.operations.SetupAddRemote},
			// Disable gc, this can improve performance and should never run for executor clones.
			{Key: "setup.git.disable-gc", Env: gitStdEnv, Command: []string{"git", "-C", repoPath, "config", "--local", "gc.auto", "0"}, Operation: h.operations.SetupGitDisableGC},
			{Key: "setup.git.fetch", Env: gitStdEnv, Command: fetchCommand, Operation: h.operations.SetupGitFetch},
		}

		if len(sparseCheckout) > 0 {
			// TODO: Is this required?
			gitCommands = append(gitCommands, command.CommandSpec{
				Key:       "setup.git.sparse-checkout-config",
				Env:       gitStdEnv,
				Command:   []string{"git", "-C", repoPath, "config", "--local", "core.sparseCheckoutCone", "1"},
				Operation: h.operations.SetupGitSparseCheckoutConfig,
			})
			// TODO: `set` or `init`?
			gitCommands = append(gitCommands, command.CommandSpec{
				Key:       "setup.git.sparse-checkout-init",
				Env:       gitStdEnv,
				Command:   []string{"git", "-C", repoPath, "sparse-checkout", "set", "--cone"},
				Operation: h.operations.SetupGitSparseCheckoutInit,
			})
		}

		gitCommands = append(gitCommands, command.CommandSpec{
			Key: "setup.git.checkout",
			Env: gitStdEnv,
			Command: []string{
				"git",
				"-C", repoPath,
				// TODO: This is only required for sparse checkouts:
				"-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal",
				"checkout",
				"--progress",
				"--force",
				// TODO: the commit needs to be `+commit:ref`, instead of just commit. Maybe we don't need a proper branch though.
				//  "-B",
				commit,
			},
			Operation: h.operations.SetupGitCheckout,
		})

		// Now checkout the sparse checkout patterns. This will require some network activity to fetch the blobs for the patterns, so we
		// need to include the auth part here as well.
		if len(sparseCheckout) > 0 {
			// TODO: git seems to always check out the root dir contents. Is this a problem?
			gitCommands = append(gitCommands, command.CommandSpec{
				Key: "setup.git.sparse-checkout-set",
				Env: gitStdEnv,
				Command: append([]string{
					"git",
					"-C", repoPath,
					"-c", "protocol.version=2",
					"-c", authorizationOption,
					"-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal",
					"sparse-checkout",
					"set",
					"--",
				}, sparseCheckout...,
				),
				Operation: h.operations.SetupGitSparseCheckoutSet,
			})
		}

		// This is for LSIF, they rely on the origin being set to the upstream repo
		// for indexing.
		gitCommands = append(gitCommands, command.CommandSpec{
			Key: "setup.git.set-remote",
			Env: gitStdEnv,
			Command: []string{
				"git",
				"-C", repoPath,
				"remote",
				"set-url",
				"origin",
				repositoryName,
			},
			Operation: h.operations.SetupGitCheckout,
		})

		// TODO: Remove, this just validates the files are there.
		gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.ls-files", Env: gitStdEnv, Command: []string{"du", "-ahc", tempDir}, Operation: h.operations.SetupGitFetch})

		for _, spec := range gitCommands {
			if err := commandRunner.Run(ctx, spec); err != nil {
				return "", "", errors.Wrap(err, fmt.Sprintf("failed %s", spec.Key))
			}
		}
	}

	// Create the scripts path.
	if err := os.MkdirAll(filepath.Join(tempDir, command.ScriptsPath), os.ModePerm); err != nil {
		return "", "", err
	}

	return tempDir, repoPath, nil
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
