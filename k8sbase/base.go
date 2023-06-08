package k8sbase

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"time"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"gopkg.in/yaml.v3"
)

//go:embed values.yaml
var BaseValues []byte

type BaseProps struct {
	Global     infraglobal.GlobalProps `yaml:"global"`
	Ingress    IngressProps            `yaml:"ingress"`
	System     SystemProps             `yaml:"system"`
	Secrets    SecretsProps            `yaml:"secrets"`
	Auth       AuthProps               `yaml:"auth"`
	Monitoring MonitoringProps         `yaml:"monitoring"`
	Databases  DatabaseProps           `yaml:"databases"`
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
	err := yaml.Unmarshal(BaseValues, &v)
	if err != nil {
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
		decoder := yaml.NewDecoder(bytes.NewReader(valuesFileBytes))
		decoder.KnownFields(true)
		if err := decoder.Decode(&values); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}
