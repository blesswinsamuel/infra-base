//go:build no_runtime_type_checking

// external-secretsio
package externalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validatePushSecret_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validatePushSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePushSecret_ManifestParameters(props *PushSecretProps) error {
	return nil
}

func validatePushSecret_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewPushSecretParameters(scope constructs.Construct, id *string, props *PushSecretProps) error {
	return nil
}

