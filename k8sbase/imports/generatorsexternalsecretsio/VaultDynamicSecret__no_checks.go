//go:build no_runtime_type_checking

// generatorsexternal-secretsio
package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateVaultDynamicSecret_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateVaultDynamicSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVaultDynamicSecret_ManifestParameters(props *VaultDynamicSecretProps) error {
	return nil
}

func validateVaultDynamicSecret_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewVaultDynamicSecretParameters(scope constructs.Construct, id *string, props *VaultDynamicSecretProps) error {
	return nil
}

