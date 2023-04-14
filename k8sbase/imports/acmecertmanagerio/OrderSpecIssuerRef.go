// acmecert-managerio
package acmecertmanagerio


// IssuerRef references a properly configured ACME-type Issuer which should be used to create this Order.
//
// If the Issuer does not exist, processing will be retried. If the Issuer is not an 'ACME' Issuer, an error will be returned and the Order will be marked as failed.
type OrderSpecIssuerRef struct {
	// Name of the resource being referred to.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Group of the resource being referred to.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Kind of the resource being referred to.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

