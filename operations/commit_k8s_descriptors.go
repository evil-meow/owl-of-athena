package operations

import (
	"evil-meow/owl-of-athena/github_api"
	"evil-meow/owl-of-athena/service_config"
	"evil-meow/owl-of-athena/templates/k8s"
)

func CommitK8sDescriptors(config *service_config.ServiceConfig) error {
	deploymentYaml, err := k8s.BuildDeploymentYaml(config)
	if err != nil {
		return err
	}

	namespaceYaml, err := k8s.BuildNamespaceYaml(config)
	if err != nil {
		return err
	}

	serviceYaml, err := k8s.BuildServiceYaml(config)
	if err != nil {
		return err
	}

	kustomizeYaml, err := k8s.BuildKustomizeYaml(config)
	if err != nil {
		return err
	}

	kustomizeProdYaml, err := k8s.BuildKustomizeProdYaml(config)
	if err != nil {
		return err
	}

	deploymentSecretsProdYaml, err := k8s.BuildDeploymentSecretsProdYaml(config)
	if err != nil {
		return err
	}

	secretsProdYaml, err := k8s.BuildSecretsProdYaml(config)
	if err != nil {
		return err
	}

	certificateProdYaml, err := k8s.BuildCertificateProdYaml(config)
	if err != nil {
		return err
	}

	gatewayProdYaml, err := k8s.BuildGatewayProdYaml(config)
	if err != nil {
		return err
	}

	labelProdYaml, err := k8s.BuildLabelProdYaml(config)
	if err != nil {
		return err
	}

	virtualServiceProdYaml, err := k8s.BuildVirtualServiceProdYaml(config)
	if err != nil {
		return err
	}

	files := github_api.FilesToCommit{
		Files: []github_api.FileToCommit{
			{
				FilePath: "kustomize/base/deployment.yaml",
				Content:  deploymentYaml,
			},
			{
				FilePath: "kustomize/base/namespace.yaml",
				Content:  namespaceYaml,
			},
			{
				FilePath: "kustomize/base/kustomization.yaml",
				Content:  kustomizeYaml,
			},
			{
				FilePath: "kustomize/base/service.yaml",
				Content:  serviceYaml,
			},
			{
				FilePath: "kustomize/overlays/production/kustomization.yaml",
				Content:  kustomizeProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/certificate.yaml",
				Content:  certificateProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/secrets.yaml",
				Content:  secretsProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/deployment-secrets.yaml",
				Content:  deploymentSecretsProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/gateway.yaml",
				Content:  gatewayProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/deployment-label.yaml",
				Content:  labelProdYaml,
			},
			{
				FilePath: "kustomize/overlays/production/virtual-service.yaml",
				Content:  virtualServiceProdYaml,
			},
		},
	}

	err = github_api.CommitFilesToMain(&config.RepoName, files)

	return err
}
