package resourcesbase

import (
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8s-base/imports/externalsecretsio"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

type SecretsDockerCredsProps struct {
	Enabled   bool   `yaml:"enabled"`
	KeyPrefix string `yaml:"keyPrefix"`
	Namespace string `yaml:"namespace"`
}

func NewSecretsDockerCreds(scope constructs.Construct, props SecretsDockerCredsProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: &props.Namespace,
	}
	chart := cdk8s.NewChart(scope, jsii.String("docker-creds"), &cprops)
	NewExternalSecret(chart, jsii.String("externalsecret"), &ExternalSecretProps{
		Name:            jsii.String("regcred"),
		RefreshInterval: jsii.String("2m"),
		Template: &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
			Type: jsii.String("kubernetes.io/dockerconfigjson"),
			Data: &map[string]*string{
				".dockerconfigjson": jsii.String(strings.TrimSpace(dedent.String(`
					{
					  "auths": {
					    "{{ .registry }}": {
					      "auth": "{{ (printf "%s:%s" .username .password) | b64enc }}"
					    }
					  }
					}
				`))),
			},
		},
		Secrets: map[string]string{
			"registry": props.KeyPrefix + "CONTAINER_REGISTRY_URL",
			"username": props.KeyPrefix + "CONTAINER_REGISTRY_USERNAME",
			"password": props.KeyPrefix + "CONTAINER_REGISTRY_PASSWORD",
		},
	})
	return chart
}
