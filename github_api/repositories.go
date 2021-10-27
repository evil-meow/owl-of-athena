package github_api

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func CreateGitubRepo(name *string) {
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}

	client, ctx := getClient()

	isPrivate := true
	description := "Repo created by owl-of-athena"

	r := &github.Repository{Name: name, Private: &isPrivate, Description: &description}
	repo, _, err := client.Repositories.Create(*ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created new repo: %v\n", repo.GetName())
}

func IsGithubRepoCreated(name string) bool {
	client, ctx := getClient()
	_, _, err := client.Repositories.Get(*ctx, "evil-meow", name)
	return err != nil
}

func getClient() (*github.Client, *context.Context) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), &ctx
}
