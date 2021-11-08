package k8s

import (
	"bytes"
	"evil-meow/owl-of-athena/service_config"
	"html/template"
)

func BuildVirtualServiceProdYaml(config *service_config.ServiceConfig) (string, error) {

	templateText := `---
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: {{.Name}}
  namespace: {{.Name}}
spec:
  gateways:
    - {{.Name}}-gateway
  hosts:
    - {{.Url}}
  http:
    - match:
      route:
        - destination:
            host: {{.Name}}
            port:
              number: {{.InternalPort}}
`

	t, err := template.New("kustomize").Parse(templateText)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "kustomize", config)
	return buf.String(), err
}
