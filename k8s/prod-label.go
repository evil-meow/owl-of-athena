package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/config"
	"html/template"
)

func BuildLabelProdYaml(config *config.Config) (string, error) {

	templateText := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  selector:
  matchLabels:
    env: production
  template:
  metadata:
    labels:
    env: production
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
