package certmanagerio


// Use the Cloudflare API to manage DNS01 challenge records.
type IssuerSpecAcmeSolversDns01Cloudflare struct {
	// API key to use to authenticate with Cloudflare.
	//
	// Note: using an API token to authenticate is now the recommended method as it allows greater control of permissions.
	ApiKeySecretRef *IssuerSpecAcmeSolversDns01CloudflareApiKeySecretRef `field:"optional" json:"apiKeySecretRef" yaml:"apiKeySecretRef"`
	// API token used to authenticate with Cloudflare.
	ApiTokenSecretRef *IssuerSpecAcmeSolversDns01CloudflareApiTokenSecretRef `field:"optional" json:"apiTokenSecretRef" yaml:"apiTokenSecretRef"`
	// Email of the account, only required when using API key based authentication.
	Email *string `field:"optional" json:"email" yaml:"email"`
}

