package externalsecretsio


// Scaleway.
type ClusterSecretStoreV1Beta1SpecProviderScaleway struct {
	// AccessKey is the non-secret part of the api key.
	AccessKey *ClusterSecretStoreV1Beta1SpecProviderScalewayAccessKey `field:"required" json:"accessKey" yaml:"accessKey"`
	// ProjectID is the id of your project, which you can find in the console: https://console.scaleway.com/project/settings.
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Region where your secrets are located: https://developers.scaleway.com/en/quickstart/#region-and-zone.
	Region *string `field:"required" json:"region" yaml:"region"`
	// SecretKey is the non-secret part of the api key.
	SecretKey *ClusterSecretStoreV1Beta1SpecProviderScalewaySecretKey `field:"required" json:"secretKey" yaml:"secretKey"`
	// APIURL is the url of the api to use.
	//
	// Defaults to https://api.scaleway.com
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
}

