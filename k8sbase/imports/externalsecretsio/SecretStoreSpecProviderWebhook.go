// external-secretsio
package externalsecretsio


// Webhook configures this store to sync secrets using a generic templated webhook.
type SecretStoreSpecProviderWebhook struct {
	// Result formatting.
	Result *SecretStoreSpecProviderWebhookResult `field:"required" json:"result" yaml:"result"`
	// Webhook url to call.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Body.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// PEM encoded CA bundle used to validate webhook server certificate.
	//
	// Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.
	CaBundle *string `field:"optional" json:"caBundle" yaml:"caBundle"`
	// The provider for the CA bundle to use to validate webhook server certificate.
	CaProvider *SecretStoreSpecProviderWebhookCaProvider `field:"optional" json:"caProvider" yaml:"caProvider"`
	// Headers.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Webhook Method.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name.
	Secrets *[]*SecretStoreSpecProviderWebhookSecrets `field:"optional" json:"secrets" yaml:"secrets"`
	// Timeout.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

