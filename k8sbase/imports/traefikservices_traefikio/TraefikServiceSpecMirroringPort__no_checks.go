//go:build no_runtime_type_checking

package traefikservices_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateTraefikServiceSpecMirroringPort_FromNumberParameters(value *float64) error {
	return nil
}

func validateTraefikServiceSpecMirroringPort_FromStringParameters(value *string) error {
	return nil
}

