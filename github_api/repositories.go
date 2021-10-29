package github_api

import (
	"log"

	"github.com/google/go-github/v39/github"
)

func CreateGitubRepo(name *string) {
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}

	client, ctx := GetClient()

	isPrivate := true
	description := "Repo created by owl-of-athena"

	r := &github.Repository{Name: name, Private: &isPrivate, Description: &description}
	repo, _, err := client.Repositories.Create(*ctx, "evil-meow", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created new repo: %v\n", repo.GetName())
}

func IsGithubRepoCreated(name string) bool {
	client, ctx := GetClient()
	_, _, err := client.Repositories.Get(*ctx, "evil-meow", name)
	return err == nil
}
