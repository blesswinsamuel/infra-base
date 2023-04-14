//go:build no_runtime_type_checking

// acmecert-managerio
package acmecertmanagerio

// Building without runtime type checking enabled, so all the below just return nil

func validateOrder_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateOrder_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOrder_ManifestParameters(props *OrderProps) error {
	return nil
}

func validateOrder_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewOrderParameters(scope constructs.Construct, id *string, props *OrderProps) error {
	return nil
}

