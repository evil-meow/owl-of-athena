package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/config"
	"html/template"
)

func BuildSecretsProdYaml(config *config.Config) (string, error) {

	templateText := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  template:
  spec:
    containers:
    - name: {{.Name}}
    envFrom:
    - secretRef:
      name: {{.Name}}-production-secrets
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
