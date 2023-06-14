package k8sbase

import (
	"strings"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	"github.com/muesli/reflow/dedent"
)

type SecretsDockerCredsProps struct {
	Enabled   bool   `json:"enabled"`
	KeyPrefix string `json:"keyPrefix"`
	Namespace string `json:"namespace"`
}

func NewSecretsDockerCreds(scope packager.Construct, props SecretsDockerCredsProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := packager.ChartProps{
		Namespace: props.Namespace,
	}
	chart := scope.Chart("docker-creds", cprops)
	k8sapp.NewExternalSecret(chart, jsii.String("externalsecret"), &k8sapp.ExternalSecretProps{
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
