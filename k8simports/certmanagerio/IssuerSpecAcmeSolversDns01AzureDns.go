package certmanagerio


// Use the Microsoft Azure DNS API to manage DNS01 challenge records.
type IssuerSpecAcmeSolversDns01AzureDns struct {
	// resource group the DNS zone is located in.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// ID of the Azure subscription.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// if both this and ClientSecret are left unset MSI will be used.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// if both this and ClientID are left unset MSI will be used.
	ClientSecretSecretRef *IssuerSpecAcmeSolversDns01AzureDnsClientSecretSecretRef `field:"optional" json:"clientSecretSecretRef" yaml:"clientSecretSecretRef"`
	// name of the Azure environment (default AzurePublicCloud).
	Environment IssuerSpecAcmeSolversDns01AzureDnsEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// name of the DNS zone that should be used.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// managed identity configuration, can not be used at the same time as clientID, clientSecretSecretRef or tenantID.
	ManagedIdentity *IssuerSpecAcmeSolversDns01AzureDnsManagedIdentity `field:"optional" json:"managedIdentity" yaml:"managedIdentity"`
	// when specifying ClientID and ClientSecret then this field is also needed.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

