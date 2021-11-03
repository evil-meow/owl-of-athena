package github_api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-github/v39/github"
)

type FilesToCommit struct {
	Files []FileToCommit
}

type FileToCommit struct {
	FilePath string
	Content  string
}

func ReadFile(fileUrl *string) (string, error) {
	req, err := http.NewRequest("GET", *fileUrl, nil)
	if err != nil {
		return "", err
	}

	client, _ := GetClient()

	resp, err := client.Client().Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	configFile, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("wrong status code received %s:\n%s", resp.Status, resp.Body)
	}

	return string(configFile), nil
}

// Commits the array of files sent to the github repo
func CommitFilesToMain(sourceRepo *string, files FilesToCommit) error {
	client, ctx := GetClient()

	entries := []*github.TreeEntry{}

	// Load each file into the tree.
	for _, file := range files.Files {
		entries = append(entries, &github.TreeEntry{Path: github.String(file.FilePath), Type: github.String("blob"), Content: github.String(file.Content), Mode: github.String("100644")})
	}

	main, err := GetMainRef(client, ctx, sourceRepo)
	if err != nil {
		return err
	}

	tree, _, err := client.Git.CreateTree(*ctx, "evil-meow", *sourceRepo, *main.Object.SHA, entries)
	if err != nil {
		return err
	}

	parent, _, err := client.Repositories.GetCommit(*ctx, "evil-meow", *sourceRepo, *main.Object.SHA, nil)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	author_name := "owl-of-athena"
	author_email := "owl-of-athena@evilmeow.com"
	commit_message := "Automated commit"

	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: &author_name, Email: &author_email}
	commit := &github.Commit{Author: author, Message: &commit_message, Tree: tree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := client.Git.CreateCommit(*ctx, "evil-meow", *sourceRepo, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	main.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(*ctx, "evil-meow", *sourceRepo, main, false)
	if err != nil {
		return err
	}

	return nil
}
