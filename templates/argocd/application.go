package argocd

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildApplicationYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: {{.Name}}
  namespace: argocd
  annotations:
    argocd-image-updater.argoproj.io/image-list: '{{.Name}}=rg.fr-par.scw.cloud/evilmeow/{{.Name}}'
    argocd-image-updater.argoproj.io/owl-of-athena.update-strategy: latest

spec:
  project: default

  source:
    repoURL: git@github.com:evil-meow/{{.RepoName}}.git
    targetRevision: HEAD
    path: kustomize/overlays/production

  destination:
    server: https://kubernetes.default.svc
    namespace: {{.Name}}

  syncPolicy:
    automated: # automated sync by default retries failed attempts 5 times with following delays between attempts ( 5s, 10s, 20s, 40s, 80s ); retry controlled using retry field.
      prune: true # Specifies if resources should be pruned during auto-syncing ( false by default ).
      selfHeal: true # Specifies if partial app sync should be executed when resources are changed only in target Kubernetes cluster and no git change detected ( false by default ).
      allowEmpty: false # Allows deleting all application resources during automatic syncing ( false by default ).

    syncOptions:     # Sync options which modifies sync behavior
    - CreateNamespace=true # Namespace Auto-Creation ensures that namespace specified as the application destination exists in the destination cluster.
    - PrunePropagationPolicy=foreground # Supported policies are background, foreground and orphan.
    - PruneLast=true # Allow the ability for resource pruning to happen as a final, implicit wave of a sync operation

    retry:
      limit: 5 # number of failed sync attempt retries; unlimited number of attempts if less than 0
      backoff:
        duration: 5s # the amount to back off. Default unit is seconds, but could also be a duration (e.g. "2m", "1h")
        factor: 2 # a factor to multiply the base duration after each failed retry
        maxDuration: 3m # the maximum amount of time allowed for the backoff strategy
`

	t, err := template.New("kustomization").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomization", config)
	return buf.String(), err
}
