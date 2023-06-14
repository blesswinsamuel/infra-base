package k8sbase

import (
	_ "embed"
	"log"
	"os"
	"time"

	"github.com/blesswinsamuel/infra-base/packager"
	"github.com/goccy/go-yaml"
)

//go:embed values.yaml
var BaseValues []byte

type BaseProps struct {
	Global     GlobalProps     `json:"global"`
	Ingress    IngressProps    `json:"ingress"`
	System     SystemProps     `json:"system"`
	Secrets    SecretsProps    `json:"secrets"`
	Auth       AuthProps       `json:"auth"`
	Monitoring MonitoringProps `json:"monitoring"`
	Databases  DatabaseProps   `json:"databases"`
}

func logModuleTiming(moduleName string) func() {
	startTime := time.Now()
	log.Printf("Starting %q..", moduleName)
	return func() {
		log.Printf("Done %q in %s", moduleName, time.Since(startTime))
	}
}

func NewBase(scope packager.Construct, props BaseProps) packager.Construct {
	defer logModuleTiming("base")()
	construct := scope.Construct("base")

	// secrets
	NewSecrets(construct, props.Secrets)

	// ingress
	NewIngress(construct, props.Ingress)

	// system
	NewSystem(construct, props.System)

	// monitoring
	NewMonitoring(construct, props.Monitoring)

	// database
	NewDatabase(construct, props.Databases)

	// auth
	NewAuth(construct, props.Auth)

	return construct
}

func GetBaseValues() BaseProps {
	v := BaseProps{}
	if err := yaml.UnmarshalWithOptions(BaseValues, &v, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		panic(err)
	}
	return v
}

func LoadValues[T any](values *T, valuesFiles []string) {
	for _, valuesFile := range valuesFiles {
		valuesFileBytes, err := os.ReadFile(valuesFile)
		if err != nil {
			log.Fatalf("ReadFile: %v", err)
		}
		// TODO: maybe go back to gopkg.in/yaml.v3 library
		// "github.com/goccy/go-yaml" library overwrites the map (example: grafana dashboards)
		// mergo approach doesn't override boolean values because it considers them as zero values
		// validation
		// var values T
		customUnmarshaller := yaml.CustomUnmarshaler(func(dst *map[string]GrafanaDashboardsConfigProps, b []byte) error {
			// workaround to allow merging of maps
			v := map[string]GrafanaDashboardsConfigProps{}
			if err := yaml.UnmarshalWithOptions(b, &v, yaml.Strict()); err != nil {
				return err
			}
			for k, v := range v {
				(*dst)[k] = v
			}
			return nil
		})
		if err := yaml.UnmarshalWithOptions(valuesFileBytes, values, yaml.Strict(), yaml.UseJSONUnmarshaler(), customUnmarshaller); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}
