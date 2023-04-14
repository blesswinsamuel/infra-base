//go:build no_runtime_type_checking

// generatorsexternal-secretsio
package generatorsexternalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateFake_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateFake_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFake_ManifestParameters(props *FakeProps) error {
	return nil
}

func validateFake_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewFakeParameters(scope constructs.Construct, id *string, props *FakeProps) error {
	return nil
}

