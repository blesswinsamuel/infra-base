//go:build no_runtime_type_checking

// generatorsexternal-secretsio
package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateEcrAuthorizationToken_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateEcrAuthorizationToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcrAuthorizationToken_ManifestParameters(props *EcrAuthorizationTokenProps) error {
	return nil
}

func validateEcrAuthorizationToken_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewEcrAuthorizationTokenParameters(scope constructs.Construct, id *string, props *EcrAuthorizationTokenProps) error {
	return nil
}

