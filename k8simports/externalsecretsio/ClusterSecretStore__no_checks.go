//go:build no_runtime_type_checking

package externalsecretsio

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterSecretStore_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterSecretStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterSecretStore_ManifestParameters(props *ClusterSecretStoreProps) error {
	return nil
}

func validateClusterSecretStore_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterSecretStoreParameters(scope constructs.Construct, id *string, props *ClusterSecretStoreProps) error {
	return nil
}

