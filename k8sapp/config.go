package k8sapp

import "github.com/blesswinsamuel/kgen"

type Config struct {
	CacheDir string
}

func GetConfig(scope kgen.Scope) Config {
	return scope.GetContext("config").(Config)
}

func SetConfig(scope kgen.Scope, props Config) {
	scope.SetContext("config", props)
}
