package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildSecretsProdYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: v1
kind: Secret
metadata:
  name: {{.Name}}-production-secrets
type: Opaque
data:
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
