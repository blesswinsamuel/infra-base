package k8sbase

import (
	_ "embed"
	"log"
	"os"
	"time"

	"github.com/Velocidex/yaml"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
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

func NewBase(scope constructs.Construct, props BaseProps) constructs.Construct {
	defer logModuleTiming("base")()
	construct := constructs.NewConstruct(scope, jsii.String("base"))

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
	if err := yaml.UnmarshalStrict(BaseValues, &v); err != nil {
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
		// // TODO: maybe go back to gopkg.in/yaml.v3 library
		// // "github.com/goccy/go-yaml" library overwrites the map (example: grafana dashboards)
		// // mergo approach doesn't override boolean values because it considers them as zero values
		// // validation
		// var values T
		// if err := yaml.UnmarshalWithOptions(valuesFileBytes, &values, yaml.Strict()); err != nil {
		// 	log.Fatalf("Unmarshal: %v", err)
		// }
		// if err := mergo.MapWithOverwrite(&ret, values); err != nil {
		// 	log.Fatalf("Merge: %v", err)
		// }
		// k8syaml.YAMLToJSONStrict()
		if err := yaml.UnmarshalStrict(valuesFileBytes, values); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}
