package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SystemProps struct {
	Reloader            ReloaderProps            `yaml:"reloader"`
	KubernetesDashboard KubernetesDashboardProps `yaml:"kubernetesDashboard"`
}

func NewSystem(scope constructs.Construct, props SystemProps) constructs.Construct {
	construct := constructs.NewConstruct(scope, jsii.String("system"))

	NewNamespace(construct, "system")

	NewReloader(construct, props.Reloader)
	NewKubernetesDashboard(construct, props.KubernetesDashboard)

	return construct
}
