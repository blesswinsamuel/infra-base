package certmanagerio


// A reference to a service account that will be used to request a bound token (also known as "projected token").
//
// Compared to using "secretRef", using this field means that you don't rely on statically bound tokens. To use this field, you must configure an RBAC rule to let cert-manager request a token.
type IssuerSpecVaultAuthKubernetesServiceAccountRef struct {
	// Name of the ServiceAccount used to request a token.
	Name *string `field:"required" json:"name" yaml:"name"`
}

