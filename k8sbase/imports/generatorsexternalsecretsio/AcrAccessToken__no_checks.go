//go:build no_runtime_type_checking

// generatorsexternal-secretsio
package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateAcrAccessToken_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateAcrAccessToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAcrAccessToken_ManifestParameters(props *AcrAccessTokenProps) error {
	return nil
}

func validateAcrAccessToken_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewAcrAccessTokenParameters(scope constructs.Construct, id *string, props *AcrAccessTokenProps) error {
	return nil
}

