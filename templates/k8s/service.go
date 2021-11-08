package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildServiceYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}
spec:
  selector:
    app: {{.Name}}
  ports:
  - protocol: TCP
    port: {{.InternalPort}}
    targetPort: {{.InternalPort}}
`

	t, err := template.New("namespace").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "namespace", config)
	return buf.String(), err
}
