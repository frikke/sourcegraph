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

	repoPath := filepath.Join(tempDir, "repo")

	if err := os.Mkdir(repoPath, os.ModePerm); err != nil {
		return "", err
	}

	if repositoryName != "" {
		cloneURL, err := makeRelativeURL(
			h.options.ClientOptions.EndpointOptions.URL,
			h.options.GitServicePath,
			repositoryName,
		)
		if err != nil {
			return "", err
		}
		// cloneURL, err = url.Parse("https://github.com/sourcegraph/sourcegraph")
		// if err != nil {
		// 	return "", err
		// }

		authorizationOption := fmt.Sprintf(
			"http.extraHeader=Authorization: %s %s",
			SchemeExecutorToken,
			h.options.ClientOptions.EndpointOptions.Token,
		)

		// TODO: Can we do something here to improve caching?
		// Like keeping the same tmp dir for the .git but checking out into multiple workspaces?
		// This might help with performance on configurations with fewer executors,
		// but maybe it doesn't matter beyond a certain point.
		fetchCommand := []string{}
		if sparseCheckout != "" {
			fetchCommand = []string{"git", "-C", repoPath, "-c", "protocol.version=2",
				// "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal",
				"clone", cloneURL.String(), "--depth=1", "--filter=blob:none", "--no-checkout", repoPath}
		} else if shallowClone {
			fetchCommand = []string{"git", "-C", repoPath, "-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal", "fetch", cloneURL.String(), "--depth=1", commit}
		} else {
			fetchCommand = []string{"git", "-C", repoPath, "-c", "protocol.version=2", "-c", authorizationOption, "-c", "http.extraHeader=X-Sourcegraph-Actor-UID: internal", "fetch", cloneURL.String(), "-t", commit}
		}

		gitCommands := []command.CommandSpec{
			{Key: "setup.git.init", Command: []string{"git", "-C", repoPath, "init"}, Operation: h.operations.SetupGitInit},
			{Key: "setup.git.fetch", Command: fetchCommand, Operation: h.operations.SetupGitFetch},
			{Key: "setup.git.add-remote", Command: []string{"git", "-C", repoPath, "remote", "add", "origin", repositoryName}, Operation: h.operations.SetupAddRemote},
		}
		if sparseCheckout != "" {
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.checkoutsparsecone", Command: []string{"git", "-C", repoPath, "sparse-checkout", "init", "--cone"}, Operation: h.operations.SetupGitCheckout})
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.checkout", Command: []string{"git", "-C", repoPath, "checkout", commit}, Operation: h.operations.SetupGitCheckout})
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.checkoutsparse", Command: []string{"git", "-C", repoPath, "sparse-checkout", "set", sparseCheckout}, Operation: h.operations.SetupGitCheckout})
		} else {
			gitCommands = append(gitCommands, command.CommandSpec{Key: "setup.git.checkout", Command: []string{"git", "-C", repoPath, "checkout", commit}, Operation: h.operations.SetupGitCheckout})
		}
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
