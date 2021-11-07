package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildKustomizeProdYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

bases:
- ../../base

nameSuffix: -prod

resources:
- certificate.yaml
- gateway.yaml
- virtual-service.yaml

patchesStrategicMerge:
- deployment-label.yaml
- deployment-secrets.yaml
`

	t, err := template.New("kustomization").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomization", config)
	return buf.String(), err
}
