//go:build no_runtime_type_checking

package externalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterExternalSecret_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterExternalSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterExternalSecret_ManifestParameters(props *ClusterExternalSecretProps) error {
	return nil
}

func validateClusterExternalSecret_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterExternalSecretParameters(scope constructs.Construct, id *string, props *ClusterExternalSecretProps) error {
	return nil
}

