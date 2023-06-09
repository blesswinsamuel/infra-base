//go:build no_runtime_type_checking

package traefikservices_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateTraefikService_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateTraefikService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTraefikService_ManifestParameters(props *TraefikServiceProps) error {
	return nil
}

func validateTraefikService_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewTraefikServiceParameters(scope constructs.Construct, id *string, props *TraefikServiceProps) error {
	return nil
}

