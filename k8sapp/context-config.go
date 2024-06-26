package k8sapp

import "github.com/blesswinsamuel/kgen"

type Config struct {
	CacheDir        string
	HelmKubeVersion string
}

var configContextKey = kgen.GenerateContextKey()

func GetConfig(scope kgen.Scope) Config {
	return scope.GetContext(configContextKey).(Config)
}

func SetConfig(scope kgen.Scope, props Config) {
	scope.SetContext(configContextKey, props)
}
