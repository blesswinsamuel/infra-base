// external-secretsio
package externalsecretsio


type ClusterExternalSecretSpecExternalSecretSpecDataFromRewrite struct {
	// Used to rewrite with regular expressions.
	//
	// The resulting key will be the output of a regexp.ReplaceAll operation.
	Regexp *ClusterExternalSecretSpecExternalSecretSpecDataFromRewriteRegexp `field:"optional" json:"regexp" yaml:"regexp"`
}

