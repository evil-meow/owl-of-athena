package github_api

import (
	"context"
	"fmt"

	"github.com/google/go-github/v39/github"
)

// Gets the main branch of a repository or creates it if not there
func GetMainRef(client *github.Client, ctx *context.Context, sourceRepo *string) (ref *github.Reference, err error) {
	ref, _, err = client.Git.GetRef(*ctx, "evil-meow", *sourceRepo, "refs/heads/main")
	if err != nil {
		return nil, fmt.Errorf("error finding main branch: %v", err)
	}

	return ref, err
}
