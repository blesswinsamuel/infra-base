package k8sapp

import (
	"github.com/blesswinsamuel/kgen"
	"github.com/blesswinsamuel/kgen/kaddons"
	"k8s.io/apimachinery/pkg/runtime"
)

type ChartInfo struct {
	Repo    *string `json:"repo"`
	Chart   *string `json:"chart"`
	Version *string `json:"version"`
}

type HelmProps struct {
	ChartInfo           ChartInfo
	ChartFileNamePrefix string
	ReleaseName         string
	Values              map[string]interface{}
	PatchObject         func(obj runtime.Object) error
}

func NewHelm(scope kgen.Scope, props *HelmProps) {
	kaddons.AddHelmChart(scope, kaddons.HelmChartProps{
		ChartInfo: kaddons.HelmChartInfo{
			Repo:    *props.ChartInfo.Repo,
			Chart:   *props.ChartInfo.Chart,
			Version: *props.ChartInfo.Version,
		},
		ChartFileNamePrefix: props.ChartFileNamePrefix,
		ReleaseName:         props.ReleaseName,
		Namespace:           scope.Namespace(),
		Values:              props.Values,
		PatchObject:         props.PatchObject,
	})
}
