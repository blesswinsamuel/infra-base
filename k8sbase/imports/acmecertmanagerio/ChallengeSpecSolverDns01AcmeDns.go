package acmecertmanagerio


// Use the 'ACME DNS' (https://github.com/joohoi/acme-dns) API to manage DNS01 challenge records.
type ChallengeSpecSolverDns01AcmeDns struct {
	// A reference to a specific 'key' within a Secret resource.
	//
	// In some instances, `key` is a required field.
	AccountSecretRef *ChallengeSpecSolverDns01AcmeDnsAccountSecretRef `field:"required" json:"accountSecretRef" yaml:"accountSecretRef"`
	Host *string `field:"required" json:"host" yaml:"host"`
}
