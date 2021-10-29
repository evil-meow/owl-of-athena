package github_api

import (
	"github.com/google/go-github/v39/github"
)

// Gets the main branch of a repository or creates it if not there
func GetMainRef(sourceRepo *string) (ref *github.Reference, err error) {
	client, ctx := GetClient()

	if ref, _, err = client.Git.GetRef(*ctx, "owl-of-athena", *sourceRepo, "refs/heads/main"); err == nil {
		return ref, nil
	}

	newRef := &github.Reference{Ref: github.String("refs/heads/main"), Object: &github.GitObject{}}
	ref, _, err = client.Git.CreateRef(*ctx, "owl-of-athena", *sourceRepo, newRef)
	return ref, err
}
