package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/config"
	"html/template"
)

func BuildGatewayProdYaml(config *config.Config) (string, error) {

	templateText := `---
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: {{.Name}}-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
    - hosts:
      - {{.Url}}
      port:
        name: https
        number: 443
        protocol: HTTPS
      tls:
        mode: SIMPLE
        credentialName: {{.Url}}-cert
  - hosts:
      - {{.Url}}
    port:
      name: http
      number: 80
      protocol: HTTP
    tls:
      httpsRedirect: true
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
