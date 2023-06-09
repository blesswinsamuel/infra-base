//go:build no_runtime_type_checking

package serverstransports_traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromNumberParameters(value *float64) error {
	return nil
}

func validateServersTransportSpecForwardingTimeoutsIdleConnTimeout_FromStringParameters(value *string) error {
	return nil
}

