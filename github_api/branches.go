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

	var baseRef *github.Reference
	if baseRef, _, err = client.Git.GetRef(*ctx, "owl-of-athena", *sourceRepo, "refs/heads/main"); err != nil {
		return nil, err
	}

	newRef := &github.Reference{Ref: github.String("refs/heads/main"), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = client.Git.CreateRef(*ctx, "owl-of-athena", *sourceRepo, newRef)
	return ref, err
}
