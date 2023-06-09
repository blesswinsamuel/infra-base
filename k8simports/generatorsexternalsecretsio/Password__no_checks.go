//go:build no_runtime_type_checking

package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validatePassword_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validatePassword_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePassword_ManifestParameters(props *PasswordProps) error {
	return nil
}

func validatePassword_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewPasswordParameters(scope constructs.Construct, id *string, props *PasswordProps) error {
	return nil
}

