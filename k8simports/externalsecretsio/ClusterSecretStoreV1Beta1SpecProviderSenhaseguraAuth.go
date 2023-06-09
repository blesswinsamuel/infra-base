package externalsecretsio


// Auth defines parameters to authenticate in senhasegura.
type ClusterSecretStoreV1Beta1SpecProviderSenhaseguraAuth struct {
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.
	ClientSecretSecretRef *ClusterSecretStoreV1Beta1SpecProviderSenhaseguraAuthClientSecretSecretRef `field:"required" json:"clientSecretSecretRef" yaml:"clientSecretSecretRef"`
}

