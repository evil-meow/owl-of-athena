package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildCertificateProdYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: {{.Url}}-cert
  namespace: istio-system
spec:
  commonName: {{.Url}}
  secretName: {{.Url}}-cert
  dnsNames:
    - {{.Url}}
  issuerRef:
    name: letsencrypt-production
    kind: ClusterIssuer
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
