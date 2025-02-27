package git

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nxpkg/scip-go/internal/command"
)

// InferRepo gets a human-readable repository name from the git clone enclosing
// the given directory.
func InferRepo(dir string) (string, error) {
	remoteURL, err := command.Run(dir, "git", "remote", "get-url", "origin")
	if err != nil {
		return "", err
	}

	return parseRemote(remoteURL)
}

// parseRemote converts a git origin url into a Nxpkg-friendly repo name.
func parseRemote(remoteURL string) (string, error) {
	// e.g., git@github.com:nxpkg/scip-go.git
	if strings.HasPrefix(remoteURL, "git@") {
		if parts := strings.Split(remoteURL, ":"); len(parts) == 2 {
			return strings.Join([]string{
				strings.TrimPrefix(parts[0], "git@"),
				strings.TrimSuffix(parts[1], ".git"),
			}, "/"), nil
		}
	}

	// e.g., https://github.com/nxpkg/scip-go.git
	if url, err := url.Parse(remoteURL); err == nil {
		return url.Hostname() + strings.TrimSuffix(url.Path, ".git"), nil
	}

	return "", fmt.Errorf("unrecognized remote URL: %s", remoteURL)
}
