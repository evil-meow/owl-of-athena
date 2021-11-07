package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/config"
	"html/template"
)

func BuildDeploymentYaml(config *config.Config) (string, error) {

	templateText := `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  selector:
    matchLabels:
      app: {{.Name}}
  replicas: {{.Replicas}}
  template:
    metadata:
      labels:
        app: {{.Name}}
    spec:
      containers:
      - name: {{.Name}}
        image: {{.Image}}:latest
      imagePullSecrets:
      - name: evilmeow-registry-secret
`

	t, err := template.New("deployment").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "deployment", config)
	return buf.String(), err
}
