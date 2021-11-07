package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildNamespaceYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: v1
kind: Namespace
metadata:
  name: {{.Name}}
  labels:
    istio-injection: enabled
`

	t, err := template.New("namespace").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "namespace", config)
	return buf.String(), err
}
