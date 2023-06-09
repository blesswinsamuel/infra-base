//go:build no_runtime_type_checking

package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateGcrAccessToken_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateGcrAccessToken_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGcrAccessToken_ManifestParameters(props *GcrAccessTokenProps) error {
	return nil
}

func validateGcrAccessToken_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewGcrAccessTokenParameters(scope constructs.Construct, id *string, props *GcrAccessTokenProps) error {
	return nil
}

