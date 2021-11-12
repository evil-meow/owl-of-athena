package operations

import (
	"errors"
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/k8s_api"
	"evil-meow/owl-of-athena/service_config"
	"evil-meow/owl-of-athena/templates/argocd"
	"fmt"
	"log"
)

func CommitArgocdDescriptor(config *service_config.ServiceConfig) error {
	argoYaml, err := argocd.BuildApplicationYaml(config)
	if err != nil {
		return err
	}

	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "argocd.yaml",
				Content:  argoYaml,
			},
		},
	}

	err = github_api.CommitFilesToMain(&config.RepoName, files)

	return err
}

func ApplyArgocdDescriptor(config *service_config.ServiceConfig) error {
	argoFileUrl := fmt.Sprintf("https://raw.githubusercontent.com/evil-meow/%s/main/argocd.yaml", config.RepoName)
	argoFileContents, err := github_api.ReadFile(&argoFileUrl)
	if err != nil {
		log.Printf("Could not find argocd.yaml CRD file at: %s", argoFileUrl)
		return errors.New("argocd.yaml file not found")
	}

	err = k8s_api.Apply(argoFileContents)
	return err
}
