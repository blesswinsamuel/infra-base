package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SecretsProps struct {
	ExternalSecrets    ExternalSecretsProps    `yaml:"externalSecrets"`
	DockerCreds        SecretsDockerCredsProps `yaml:"dockerCreds"`
	ClusterSecretStore ClusterSecretStoreProps `yaml:"clusterSecretStore"`
}

func NewSecrets(scope constructs.Construct, props SecretsProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("secrets"))

	NewNamespace(construct, "secrets")

	NewExternalSecrets(construct, props.ExternalSecrets)

	NewSecretsDockerCreds(construct, props.DockerCreds)
	NewClusterSecretStore(construct, props.ClusterSecretStore)
	return construct
}
