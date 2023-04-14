// acmecert-managerio
package acmecertmanagerio


// managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID.
type ChallengeSpecSolverDns01AzureDnsManagedIdentity struct {
	// client ID of the managed identity, can not be used at the same time as resourceID.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// resource ID of the managed identity, can not be used at the same time as clientID.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
}

