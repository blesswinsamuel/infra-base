package externalsecretsio


// Used to configure the provider.
//
// Only one provider may be set.
type ClusterSecretStoreV1Beta1SpecProvider struct {
	// Akeyless configures this store to sync secrets using Akeyless Vault provider.
	Akeyless *ClusterSecretStoreV1Beta1SpecProviderAkeyless `field:"optional" json:"akeyless" yaml:"akeyless"`
	// Alibaba configures this store to sync secrets using Alibaba Cloud provider.
	Alibaba *ClusterSecretStoreV1Beta1SpecProviderAlibaba `field:"optional" json:"alibaba" yaml:"alibaba"`
	// AWS configures this store to sync secrets using AWS Secret Manager provider.
	Aws *ClusterSecretStoreV1Beta1SpecProviderAws `field:"optional" json:"aws" yaml:"aws"`
	// AzureKV configures this store to sync secrets using Azure Key Vault provider.
	Azurekv *ClusterSecretStoreV1Beta1SpecProviderAzurekv `field:"optional" json:"azurekv" yaml:"azurekv"`
	// Doppler configures this store to sync secrets using the Doppler provider.
	Doppler *ClusterSecretStoreV1Beta1SpecProviderDoppler `field:"optional" json:"doppler" yaml:"doppler"`
	// Fake configures a store with static key/value pairs.
	Fake *ClusterSecretStoreV1Beta1SpecProviderFake `field:"optional" json:"fake" yaml:"fake"`
	// GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider.
	Gcpsm *ClusterSecretStoreV1Beta1SpecProviderGcpsm `field:"optional" json:"gcpsm" yaml:"gcpsm"`
	// Gitlab configures this store to sync secrets using Gitlab Variables provider.
	Gitlab *ClusterSecretStoreV1Beta1SpecProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// IBM configures this store to sync secrets using IBM Cloud provider.
	Ibm *ClusterSecretStoreV1Beta1SpecProviderIbm `field:"optional" json:"ibm" yaml:"ibm"`
	// KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider.
	Keepersecurity *ClusterSecretStoreV1Beta1SpecProviderKeepersecurity `field:"optional" json:"keepersecurity" yaml:"keepersecurity"`
	// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
	Kubernetes *ClusterSecretStoreV1Beta1SpecProviderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// OnePassword configures this store to sync secrets using the 1Password Cloud provider.
	Onepassword *ClusterSecretStoreV1Beta1SpecProviderOnepassword `field:"optional" json:"onepassword" yaml:"onepassword"`
	// Oracle configures this store to sync secrets using Oracle Vault provider.
	Oracle *ClusterSecretStoreV1Beta1SpecProviderOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// Scaleway.
	Scaleway *ClusterSecretStoreV1Beta1SpecProviderScaleway `field:"optional" json:"scaleway" yaml:"scaleway"`
	// Senhasegura configures this store to sync secrets using senhasegura provider.
	Senhasegura *ClusterSecretStoreV1Beta1SpecProviderSenhasegura `field:"optional" json:"senhasegura" yaml:"senhasegura"`
	// Vault configures this store to sync secrets using Hashi provider.
	Vault *ClusterSecretStoreV1Beta1SpecProviderVault `field:"optional" json:"vault" yaml:"vault"`
	// Webhook configures this store to sync secrets using a generic templated webhook.
	Webhook *ClusterSecretStoreV1Beta1SpecProviderWebhook `field:"optional" json:"webhook" yaml:"webhook"`
	// YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider.
	Yandexcertificatemanager *ClusterSecretStoreV1Beta1SpecProviderYandexcertificatemanager `field:"optional" json:"yandexcertificatemanager" yaml:"yandexcertificatemanager"`
	// YandexLockbox configures this store to sync secrets using Yandex Lockbox provider.
	Yandexlockbox *ClusterSecretStoreV1Beta1SpecProviderYandexlockbox `field:"optional" json:"yandexlockbox" yaml:"yandexlockbox"`
}

