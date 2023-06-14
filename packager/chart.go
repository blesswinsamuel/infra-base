package packager

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type Chart interface {
	Construct
	Namespace() string
}

type ChartProps struct {
	Namespace string
}

type cdk8schart struct {
	cdk8sConstruct
	props ChartProps
}

func (c *cdk8schart) Namespace() string {
	return c.props.Namespace
}

func (c *cdk8sConstruct) Chart(id string, props ChartProps) Chart {
	chart := cdk8s.NewChart(c.construct, &id, &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Namespace:                 jsii.String(props.Namespace),
	})
	return &cdk8schart{
		cdk8sConstruct: cdk8sConstruct{chart},
		props:          props,
	}
}

// func newChart(scope Construct, id string, props *ChartProps) Chart {
// 	// TODO
// 	chart := cdk8s.NewChart(getCdk8sConstruct(scope), &id, &cdk8s.ChartProps{
// 		DisableResourceNameHashes: jsii.Bool(true),
// 		Namespace:                 jsii.String(props.Namespace),
// 	})
// 	return &cdk8schart{
// 		cdk8sConstruct: cdk8sConstruct{chart},
// 		props:          props,
// 	}
// }

type chart struct {
	construct
	props ChartProps
}

func (c *chart) Namespace() string {
	return c.props.Namespace
}

func (c *construct) Chart(id string, props ChartProps) Chart {
	return &chart{}
}
