// external-secretsio
package externalsecretsio


// Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.
type ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuth struct {
	// the Akeyless Kubernetes auth-method access-id.
	AccessId *string `field:"required" json:"accessId" yaml:"accessId"`
	// Kubernetes-auth configuration name in Akeyless-Gateway.
	K8SConfName *string `field:"required" json:"k8SConfName" yaml:"k8SConfName"`
	// Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless.
	//
	// If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.
	SecretRef *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuthSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// Optional service account field containing the name of a kubernetes ServiceAccount.
	//
	// If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.
	ServiceAccountRef *ClusterSecretStoreV1Beta1SpecProviderAkeylessAuthSecretRefKubernetesAuthServiceAccountRef `field:"optional" json:"serviceAccountRef" yaml:"serviceAccountRef"`
}

