package resourcesbase

import (
	_ "embed"
	"log"
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gopkg.in/yaml.v3"
)

//go:embed values.yaml
var BaseValues []byte

type BaseProps struct {
	Global     GlobalProps     `yaml:"global"`
	Ingress    IngressProps    `yaml:"ingress"`
	System     SystemProps     `yaml:"system"`
	Secrets    SecretsProps    `yaml:"secrets"`
	Auth       AuthProps       `yaml:"auth"`
	Monitoring MonitoringProps `yaml:"monitoring"`
	Databases  DatabaseProps   `yaml:"databases"`
}

func NewBase(scope constructs.Construct, props BaseProps) constructs.Construct {
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
		if err := yaml.Unmarshal(valuesFileBytes, &values); err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}
