package github_api

import (
	"fmt"
	"io"
	"net/http"

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
func CommitFilesToMain(sourceRepo *string, files FilesToCommit, ref *github.Reference) (tree *github.Tree, err error) {
	client, ctx := GetClient()

	entries := []*github.TreeEntry{}

	// Load each file into the tree.
	for _, file := range files.Files {
		entries = append(entries, &github.TreeEntry{Path: github.String(file.FilePath), Type: github.String("blob"), Content: github.String(file.Content), Mode: github.String("100644")})
	}

	tree, _, err = client.Git.CreateTree(*ctx, "owl-of-athena", *sourceRepo, *ref.Object.SHA, entries)
	return tree, err
}
