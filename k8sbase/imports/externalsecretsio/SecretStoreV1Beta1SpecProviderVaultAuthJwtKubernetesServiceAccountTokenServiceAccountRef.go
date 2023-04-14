// external-secretsio
package externalsecretsio


// Service account field containing the name of a kubernetes ServiceAccount.
type SecretStoreV1Beta1SpecProviderVaultAuthJwtKubernetesServiceAccountTokenServiceAccountRef struct {
	// The name of the ServiceAccount resource being referred to.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list.
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Namespace of the resource being referred to.
	//
	// Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

