package k8sapp

import (
	"os"
	"time"

	"github.com/blesswinsamuel/kgen"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
}

func NewApp(props kgen.AppProps) kgen.App {
	app := kgen.NewApp(props)

	var cacheDir = os.Getenv("CACHE_DIR")
	if cacheDir == "" {
		cacheDir = "./cache"
	}

	SetConfig(app, Config{
		CacheDir: cacheDir,
	})
	return app
}

func Render(scope kgen.Scope, values Values) {
	startTime := time.Now()
	log.Info().Msg("Starting render...")

	getModuleName := func(node ast.Node, defaultValue string) string {
		if node == nil {
			return defaultValue
		}
		moduleCommons := struct {
			Module string `json:"_module"`
		}{}
		if err := yaml.NodeToValue(node, &moduleCommons, yaml.UseJSONUnmarshaler()); err != nil {
			printErrIfPretty(err)
			log.Panic().Err(err).Msg("NodeToValue (module)")
		}
		if moduleCommons.Module == "" {
			return defaultValue
		}
		return moduleCommons.Module
	}
	for _, key := range values.Services.keyOrder {
		namespace, services := key, values.Services.Map[key]
		t := logModuleTiming(namespace, 0)
		namespaceScope := scope.CreateScope(namespace, kgen.ScopeProps{})
		namespaceScope.SetContext("namespace", namespace)
		if namespace != "default" {
			NewNamespace(namespaceScope, namespace)
		}

		for _, key := range services.keyOrder {
			serviceName, serviceProps := key, services.Map[key]
			moduleName := getModuleName(serviceProps, serviceName)
			t := logModuleTiming(serviceName, 1)
			module := registeredModules[moduleName]
			if module == nil {
				log.Panic().Str("module", moduleName).Msgf("Module is not registered")
			}
			// fmt.Println(namespace, serviceName, service, reflect.TypeOf(module))
			if defaultValues, ok := registeredDefaultValues[moduleName]; ok && defaultValues != nil {
				if err := yaml.NodeToValue(defaultValues, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
					printErrIfPretty(err)
					log.Panic().Err(err).Msg("NodeToValue(defaults)")
				}
			}
			if serviceProps != nil {
				if err := yaml.NodeToValue(serviceProps, module, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
					printErrIfPretty(err)
					log.Panic().Err(err).Msg("NodeToValue(Module)")
				}
			}
			// unmarshal(module, service)
			scopeProps := kgen.ScopeProps{}
			serviceScope := namespaceScope.CreateScope(serviceName, scopeProps)
			serviceScope.SetContext("name", serviceName)
			module.Render(serviceScope)
			NewGrafanaDashboards(serviceScope, module.GetGrafanaDashboards())
			NewAlertingRules(serviceScope, module.GetAlertingRules())
			t()
		}
		t()
	}
	log.Info().Msgf("Render done in %s.", time.Since(startTime))
}

func Synth(app kgen.App) {
	startTime := time.Now()
	log.Info().Msg("Starting synth (writing YAMLs to disk)...")
	NewKappConfig(app)

	app.WriteYAMLsToDisk()

	log.Info().Msgf("Synth done in %s.", time.Since(startTime))
}
