package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CertManagerProps struct {
	Enabled       bool      `yaml:"enabled"`
	HelmChartInfo ChartInfo `yaml:"helm"`
}

// https://github.com/cert-manager/cert-manager/tree/master/deploy/charts/cert-manager
// https://artifacthub.io/packages/helm/cert-manager/cert-manager
func NewCertManager(scope constructs.Construct, props CertManagerProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	construct := constructs.NewConstruct(scope, jsii.String("cert-manager"))
	NewNamespace(construct, "cert-manager")

	chart := cdk8s.NewChart(construct, jsii.String("cert-manager"), &cdk8s.ChartProps{
		Namespace:                 jsii.String("cert-manager"),
		DisableResourceNameHashes: jsii.Bool(true),
	})

	helmContents := NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("cert-manager"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"installCRDs": "true",
		},
	})

	for _, obj := range *helmContents.ApiObjects() {
		if *obj.Metadata().Name() == "cert-manager-startupapicheck" || *obj.Metadata().Name() == "cert-manager-startupapicheck:create-cert" {
			helmContents.Node().TryRemoveChild(obj.Node().Id())
		}
	}

	return construct
}
