package certmanagerio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CertificateRequest is used to request a signed certificate from one of the configured issuers.
//
// All fields within the CertificateRequest's `spec` are immutable after creation. A CertificateRequest will either succeed or fail, as denoted by its `status.state` field.
// A CertificateRequest is a one-shot resource, meaning it represents a single point in time request for a certificate and cannot be re-used.
type CertificateRequestProps struct {
	// Desired state of the CertificateRequest resource.
	Spec *CertificateRequestSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

