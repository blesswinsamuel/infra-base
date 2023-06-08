package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type SecretsProps struct {
	ExternalSecrets    ExternalSecretsProps    `yaml:"externalSecrets"`
	DockerCreds        SecretsDockerCredsProps `yaml:"dockerCreds"`
	ClusterSecretStore ClusterSecretStoreProps `yaml:"clusterSecretStore"`
}

func NewSecrets(scope constructs.Construct, props SecretsProps) constructs.Construct {
	defer logModuleTiming("secrets")()

	chart := k8sapp.NewNamespaceChart(scope, "secrets")

	NewExternalSecrets(chart, props.ExternalSecrets)
	NewSecretsDockerCreds(chart, props.DockerCreds)
	NewClusterSecretStore(chart, props.ClusterSecretStore)
	return chart
}
