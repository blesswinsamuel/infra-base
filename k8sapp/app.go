package k8sapp

import (
	"fmt"
	"os"
	"time"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly})
}

func modifyObj[T any](apiObject kgen.ApiObject, f func(*T)) error {
	var res T
	statefulsetUnstructured := apiObject.GetObject().(*unstructured.Unstructured)
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(statefulsetUnstructured.UnstructuredContent(), &res)
	if err != nil {
		return fmt.Errorf("FromUnstructured: %w", err)
	}
	f(&res)
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&res)
	if err != nil {
		return fmt.Errorf("ToUnstructured: %w", err)
	}
	apiObject.SetObject(&unstructured.Unstructured{Object: unstructuredObj})
	return nil
}

func patchObject(apiObject kgen.ApiObject) error {
	dnsConfig := &corev1.PodDNSConfig{
		Options: []corev1.PodDNSConfigOption{
			{Name: "ndots", Value: infrahelpers.Ptr("1")},
		},
	}
	switch apiObject.GetKind() {
	case "Deployment":
		if err := modifyObj(apiObject, func(deployment *appsv1.Deployment) {
			deployment.Spec.Template.Spec.DNSConfig = dnsConfig
		}); err != nil {
			return err
		}
	case "StatefulSet":
		if err := modifyObj(apiObject, func(statefulset *appsv1.StatefulSet) {
			statefulset.Spec.Template.Spec.DNSConfig = dnsConfig
		}); err != nil {
			return err
		}
	case "DaemonSet":
		if err := modifyObj(apiObject, func(statefulset *appsv1.DaemonSet) {
			statefulset.Spec.Template.Spec.DNSConfig = dnsConfig
		}); err != nil {
			return err
		}
	case "CronJob":
		if err := modifyObj(apiObject, func(cronjob *batchv1.CronJob) {
			cronjob.Spec.JobTemplate.Spec.Template.Spec.DNSConfig = dnsConfig
		}); err != nil {
			return err
		}
	}
	return nil
}

type AppProps struct {
	kgen.BuilderOptions
}

func NewApp(props AppProps) kgen.Builder {
	builder := kgen.NewBuilder(props.BuilderOptions)

	var cacheDir = os.Getenv("CACHE_DIR")
	if cacheDir == "" {
		cacheDir = "./cache"
	}

	SetConfig(builder, Config{
		CacheDir: cacheDir,
	})
	return builder
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

func Synth(app kgen.Builder, patchNdots bool, opts kgen.WriteOpts) {
	startTime := time.Now()
	log.Info().Msg("Starting synth (writing YAMLs to disk)...")
	NewKappConfig(app)

	if patchNdots {
		opts.PatchObject = patchObject
	}

	app.WriteYAMLsToDisk(opts)

	log.Info().Msgf("Synth done in %s.", time.Since(startTime))
}
