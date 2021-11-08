package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildKustomizeYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: {{.Name}}

resources:
  - deployment.yaml
  - namespace.yaml
  - service.yaml
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
