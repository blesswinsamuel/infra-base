package k8sapp

import (
	"path"

	"github.com/blesswinsamuel/kgen"
	"github.com/blesswinsamuel/kgen/kaddons"
	"github.com/rs/zerolog/log"
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
	PatchObject         func(obj runtime.Object)
}

func NewHelm(scope kgen.Scope, props *HelmProps) {
	config := GetConfig(scope)
	objects, err := kaddons.ExecHelmTemplateAndGetObjects(kaddons.HelmChartProps{
		ChartInfo: kaddons.HelmChartInfo{
			Repo:    *props.ChartInfo.Repo,
			Chart:   *props.ChartInfo.Chart,
			Version: *props.ChartInfo.Version,
		},
		Namespace:           scope.Namespace(),
		ChartFileNamePrefix: props.ChartFileNamePrefix,
		ReleaseName:         props.ReleaseName,
		Values:              props.Values,
		CacheDir:            path.Join(config.CacheDir, "charts"),
		HelmKubeVersion:     config.HelmKubeVersion,
		Logger:              kaddons.CustomLogger{InfoFn: log.Info().Msgf, WarnFn: log.Warn().Msgf},
	})
	if err != nil {
		log.Panic().Err(err).Msg("failed to execute helm template")
	}
	for _, object := range objects {
		if props.PatchObject != nil {
			props.PatchObject(object)
		}
		scope.AddApiObject(object)
	}
}
