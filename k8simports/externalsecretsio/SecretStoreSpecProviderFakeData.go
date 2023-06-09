package externalsecretsio


type SecretStoreSpecProviderFakeData struct {
	Key *string `field:"required" json:"key" yaml:"key"`
	Value *string `field:"optional" json:"value" yaml:"value"`
	ValueMap *map[string]*string `field:"optional" json:"valueMap" yaml:"valueMap"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

