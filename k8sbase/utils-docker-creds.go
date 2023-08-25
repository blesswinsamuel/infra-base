package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	"github.com/muesli/reflow/dedent"
)

type UtilsDockerCreds struct {
	KeyPrefix string `json:"keyPrefix"`
}

func (props *UtilsDockerCreds) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{}
	chart := scope.Chart("docker-creds", cprops)
	k8sapp.NewExternalSecret(chart, "externalsecret", &k8sapp.ExternalSecretProps{
		Name:       "regcred",
		SecretType: "kubernetes.io/dockerconfigjson",
		Template: map[string]string{
			".dockerconfigjson": strings.TrimSpace(dedent.String(`
				{
				  "auths": {
					"{{ .registry }}": {
					  "auth": "{{ (printf "%s:%s" .username .password) | b64enc }}"
					}
				  }
				}
			`)),
		},
		RemoteRefs: map[string]string{
			"registry": props.KeyPrefix + "CONTAINER_REGISTRY_URL",
			"username": props.KeyPrefix + "CONTAINER_REGISTRY_USERNAME",
			"password": props.KeyPrefix + "CONTAINER_REGISTRY_PASSWORD",
		},
	})
	return chart
}
