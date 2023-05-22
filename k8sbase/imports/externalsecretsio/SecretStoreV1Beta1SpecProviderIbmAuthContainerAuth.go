package externalsecretsio


// IBM Container-based auth with IAM Trusted Profile.
type SecretStoreV1Beta1SpecProviderIbmAuthContainerAuth struct {
	// the IBM Trusted Profile.
	Profile *string `field:"required" json:"profile" yaml:"profile"`
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// Location the token is mounted on the pod.
	TokenLocation *string `field:"optional" json:"tokenLocation" yaml:"tokenLocation"`
}

