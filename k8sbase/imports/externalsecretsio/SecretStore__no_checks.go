//go:build no_runtime_type_checking

// external-secretsio
package externalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateSecretStore_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateSecretStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretStore_ManifestParameters(props *SecretStoreProps) error {
	return nil
}

func validateSecretStore_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewSecretStoreParameters(scope constructs.Construct, id *string, props *SecretStoreProps) error {
	return nil
}

