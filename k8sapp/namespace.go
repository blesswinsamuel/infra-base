package k8sapp

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewNamespaceChart(scope constructs.Construct, namespaceName string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(namespaceName), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Namespace:                 jsii.String(namespaceName),
	})
	SetNamespaceContext(chart, namespaceName)
	NewNamespace(chart, jsii.String("namespace"), namespaceName)
	return chart
}

func NewNamespaceChartSkipCreate(scope constructs.Construct, namespaceName string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(namespaceName), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Namespace:                 jsii.String(namespaceName),
	})
	SetNamespaceContext(chart, namespaceName)
	return chart
}

func NewNamespace(scope constructs.Construct, id *string, namespaceName string) constructs.Construct {
	return k8s.NewKubeNamespace(scope, id, &k8s.KubeNamespaceProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(namespaceName),
			Labels: &map[string]*string{
				"name": jsii.String(namespaceName),
			},
		},
	})
}

func SetNamespaceContext(scope constructs.Construct, namespaceName string) constructs.Construct {
	scope.Node().SetContext(jsii.String("namespace"), namespaceName)
	return nil
}

func GetNamespaceContextPtr(scope constructs.Construct) *string {
	return jsii.String(GetNamespaceContext(scope))
}

func GetNamespaceContext(scope constructs.Construct) string {
	return scope.Node().TryGetContext(jsii.String("namespace")).(string)
}
