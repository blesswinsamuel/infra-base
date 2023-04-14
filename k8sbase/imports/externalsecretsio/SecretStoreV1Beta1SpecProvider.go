// external-secretsio
package externalsecretsio


// Used to configure the provider.
//
// Only one provider may be set.
type SecretStoreV1Beta1SpecProvider struct {
	// Akeyless configures this store to sync secrets using Akeyless Vault provider.
	Akeyless *SecretStoreV1Beta1SpecProviderAkeyless `field:"optional" json:"akeyless" yaml:"akeyless"`
	// Alibaba configures this store to sync secrets using Alibaba Cloud provider.
	Alibaba *SecretStoreV1Beta1SpecProviderAlibaba `field:"optional" json:"alibaba" yaml:"alibaba"`
	// AWS configures this store to sync secrets using AWS Secret Manager provider.
	Aws *SecretStoreV1Beta1SpecProviderAws `field:"optional" json:"aws" yaml:"aws"`
	// AzureKV configures this store to sync secrets using Azure Key Vault provider.
	Azurekv *SecretStoreV1Beta1SpecProviderAzurekv `field:"optional" json:"azurekv" yaml:"azurekv"`
	// Doppler configures this store to sync secrets using the Doppler provider.
	Doppler *SecretStoreV1Beta1SpecProviderDoppler `field:"optional" json:"doppler" yaml:"doppler"`
	// Fake configures a store with static key/value pairs.
	Fake *SecretStoreV1Beta1SpecProviderFake `field:"optional" json:"fake" yaml:"fake"`
	// GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider.
	Gcpsm *SecretStoreV1Beta1SpecProviderGcpsm `field:"optional" json:"gcpsm" yaml:"gcpsm"`
	// Gitlab configures this store to sync secrets using Gitlab Variables provider.
	Gitlab *SecretStoreV1Beta1SpecProviderGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// IBM configures this store to sync secrets using IBM Cloud provider.
	Ibm *SecretStoreV1Beta1SpecProviderIbm `field:"optional" json:"ibm" yaml:"ibm"`
	// KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider.
	Keepersecurity *SecretStoreV1Beta1SpecProviderKeepersecurity `field:"optional" json:"keepersecurity" yaml:"keepersecurity"`
	// Kubernetes configures this store to sync secrets using a Kubernetes cluster provider.
	Kubernetes *SecretStoreV1Beta1SpecProviderKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// OnePassword configures this store to sync secrets using the 1Password Cloud provider.
	Onepassword *SecretStoreV1Beta1SpecProviderOnepassword `field:"optional" json:"onepassword" yaml:"onepassword"`
	// Oracle configures this store to sync secrets using Oracle Vault provider.
	Oracle *SecretStoreV1Beta1SpecProviderOracle `field:"optional" json:"oracle" yaml:"oracle"`
	// Scaleway.
	Scaleway *SecretStoreV1Beta1SpecProviderScaleway `field:"optional" json:"scaleway" yaml:"scaleway"`
	// Senhasegura configures this store to sync secrets using senhasegura provider.
	Senhasegura *SecretStoreV1Beta1SpecProviderSenhasegura `field:"optional" json:"senhasegura" yaml:"senhasegura"`
	// Vault configures this store to sync secrets using Hashi provider.
	Vault *SecretStoreV1Beta1SpecProviderVault `field:"optional" json:"vault" yaml:"vault"`
	// Webhook configures this store to sync secrets using a generic templated webhook.
	Webhook *SecretStoreV1Beta1SpecProviderWebhook `field:"optional" json:"webhook" yaml:"webhook"`
	// YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider.
	Yandexcertificatemanager *SecretStoreV1Beta1SpecProviderYandexcertificatemanager `field:"optional" json:"yandexcertificatemanager" yaml:"yandexcertificatemanager"`
	// YandexLockbox configures this store to sync secrets using Yandex Lockbox provider.
	Yandexlockbox *SecretStoreV1Beta1SpecProviderYandexlockbox `field:"optional" json:"yandexlockbox" yaml:"yandexlockbox"`
}

