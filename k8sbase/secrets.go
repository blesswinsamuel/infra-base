package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
)

type SecretsProps struct {
	ExternalSecrets    ExternalSecretsProps    `json:"externalSecrets"`
	DockerCreds        SecretsDockerCredsProps `json:"dockerCreds"`
	ClusterSecretStore ClusterSecretStoreProps `json:"clusterSecretStore"`
}

func NewSecrets(scope packager.Construct, props SecretsProps) packager.Construct {
	defer logModuleTiming("secrets")()

	chart := k8sapp.NewNamespaceChart(scope, "secrets")

	NewExternalSecrets(chart, props.ExternalSecrets)
	NewSecretsDockerCreds(chart, props.DockerCreds)
	NewClusterSecretStore(chart, props.ClusterSecretStore)
	return chart
}
