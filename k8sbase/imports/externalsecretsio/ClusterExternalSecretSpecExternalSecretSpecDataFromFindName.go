package externalsecretsio


// Finds secrets based on the name.
type ClusterExternalSecretSpecExternalSecretSpecDataFromFindName struct {
	// Finds secrets base.
	Regexp *string `field:"optional" json:"regexp" yaml:"regexp"`
}

