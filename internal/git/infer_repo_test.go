package git

import (
	"fmt"
	"testing"
)

func TestInferRepo(t *testing.T) {
	repo, err := InferRepo("")
	if err != nil {
		t.Fatalf("unexpected error inferring repo: %s", err)
	}

	if repo != "github.com/nxpkg/scip-go" {
		t.Errorf("unexpected remote repo. want=%q have=%q", "github.com/nxpkg/scip-go", repo)
	}
}

func TestParseRemote(t *testing.T) {
	testCases := map[string]string{
		"git@github.com:nxpkg/scip-go.git":                                "github.com/nxpkg/scip-go",
		"https://github.com/nxpkg/scip-go":                                "github.com/nxpkg/scip-go",
		"ssh://git@phabricator.company.com:2222/diffusion/COMPANY/companay.git": "phabricator.company.com/diffusion/COMPANY/companay",
	}

	for input, expectedOutput := range testCases {
		t.Run(fmt.Sprintf("input=%q", input), func(t *testing.T) {
			output, err := parseRemote(input)
			if err != nil {
				t.Fatalf("unexpected error parsing remote: %s", err)
			}

			if output != expectedOutput {
				t.Errorf("unexpected repo name. want=%q have=%q", expectedOutput, output)
			}
		})
	}
}
