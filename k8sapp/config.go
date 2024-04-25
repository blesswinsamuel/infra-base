package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/kubegogen"
)

type Config struct {
	CacheDir string
}

func GetConfig(scope kubegogen.Scope) Config {
	return scope.GetContext("config").(Config)
}

func SetConfig(scope kubegogen.Scope, props Config) {
	scope.SetContext("config", props)
}
