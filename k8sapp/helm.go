package k8sapp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/kgen"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	Namespace           string
	Values              map[string]interface{}
	PatchResource       func(resource *unstructured.Unstructured)
}

func NewHelm(scope kgen.Scope, props *HelmProps) {
	globals := GetConfig(scope)
	chartsCacheDir := fmt.Sprintf("%s/%s", globals.CacheDir, "charts")
	if err := os.MkdirAll(chartsCacheDir, os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("MkdirAll failed")
	}
	if _, err := exec.LookPath("helm"); err != nil {
		log.Fatal().Err(err).Msg("helm LookPath failed")
	}
	if props.ChartInfo.Repo == nil {
		log.Fatal().Msgf("props.ChartInfo is nil for %s", props.ReleaseName)
	}
	chartFileName := *props.ChartInfo.Chart + "-" + *props.ChartInfo.Version + ".tgz"
	if props.ChartFileNamePrefix != "" {
		chartFileName = props.ChartFileNamePrefix + *props.ChartInfo.Version + ".tgz"
	}
	chartPath := fmt.Sprintf("%s/%s", chartsCacheDir, chartFileName)
	if _, err := os.Stat(chartPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Info().Msgf("Fetching chart '%s' from repo '%s' version '%s' ...", *props.ChartInfo.Chart, *props.ChartInfo.Repo, *props.ChartInfo.Version)
			cmd := exec.Command("helm", "pull", *props.ChartInfo.Chart, "--repo", *props.ChartInfo.Repo, "--destination", chartsCacheDir, "--version", *props.ChartInfo.Version)
			if out, err := cmd.CombinedOutput(); err != nil {
				fmt.Println(string(out))
				log.Fatal().Err(err).Msg("Error occured during helm pull command")
			} else {
				if len(out) > 0 {
					log.Info().Msgf(string(out))
				}
			}
		} else {
			log.Fatal().Err(err).Msg("helm Stat failed")
		}
	}
	namespace := props.Namespace
	if namespace == "" {
		namespace = scope.Namespace()
	}

	cmd := exec.Command(
		"helm",
		"template",
		props.ReleaseName,
		chartPath,
		"--namespace",
		namespace,
		"--kube-version=v1.28.8",
		"--include-crds",
		"--skip-tests",
		"--no-hooks",
		"--values",
		"-",
	)
	cmd.Stdin = strings.NewReader(infrahelpers.ToJSONString(props.Values))
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(ee.Stderr))
		}
		log.Panic().Err(err).Msg("helm template failed")
	}

	dec := yaml.NewDecoder(bytes.NewReader(out))
	i := 0
	for {
		i++
		var obj map[string]any
		err := dec.Decode(&obj)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic().Err(err).Msg("Error decoding yaml")
		}
		if len(obj) == 0 {
			continue
		}
		runtimeObj := &unstructured.Unstructured{Object: obj}
		if props.PatchResource != nil {
			props.PatchResource(runtimeObj)
		}
		scope.AddApiObject(runtimeObj)
	}
}
