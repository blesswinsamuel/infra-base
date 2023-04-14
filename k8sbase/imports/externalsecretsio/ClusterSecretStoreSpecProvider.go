// external-secretsio
package externalsecretsio


// Used to configure the provider.
//
// Only one provider may be set.
type ClusterSecretStoreSpecProvider struct {
	// Akeyless configures this store to sync secrets using Akeyless Vault provider.
	Akeyless *ClusterSecretStoreSpecProviderAkeyless `field:"optional" json:"akeyless" yaml:"akeyless"`
	// Alibaba configures this store to sync secrets using Alibaba Cloud provider.
	Alibaba *ClusterSecretStoreSpecProviderAlibaba `field:"optional" json:"alibaba" yaml:"alibaba"`
	// AWS configures this store to sync secrets using AWS Secret Manager provider.
	Aws *ClusterSecretStoreSpecProviderAws `field:"optional" json:"aws" yaml:"aws"`
	// AzureKV configures this store to sync secrets using Azure Key Vault provider.
	Azurekv *ClusterSecretStoreSpecProviderAzurekv `field:"optional" json:"azurekv" yaml:"azurekv"`
	// Fake configures a store with static key/value pairs.
	Fake *ClusterSecretStoreSpecProviderFake `field:"optional" json:"fake" yaml:"fake"`
	// GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider.
	Gcpsm *ClusterSecretStoreSpecProviderGcpsm `field:"optional" json:"gcpsm" yaml:"gcpsm"`
	// Gitlab configures this store to sync secrets using Gitlab Variables provider.
	Gitlab *ClusterSecretStoreSpecProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// IBM configures this store to sync secrets using IBM Cloud provider.
	Ibm *ClusterSecretStoreSpecProviderIbm `field:"optional" json:"ibm" yaml:"ibm"`
	// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
	Kubernetes *ClusterSecretStoreSpecProviderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Oracle configures this store to sync secrets using Oracle Vault provider.
	Oracle *ClusterSecretStoreSpecProviderOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// Vault configures this store to sync secrets using Hashi provider.
	Vault *ClusterSecretStoreSpecProviderVault `field:"optional" json:"vault" yaml:"vault"`
	// Webhook configures this store to sync secrets using a generic templated webhook.
	Webhook *ClusterSecretStoreSpecProviderWebhook `field:"optional" json:"webhook" yaml:"webhook"`
	// YandexLockbox configures this store to sync secrets using Yandex Lockbox provider.
	Yandexlockbox *ClusterSecretStoreSpecProviderYandexlockbox `field:"optional" json:"yandexlockbox" yaml:"yandexlockbox"`
}

