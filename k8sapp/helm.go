package k8sapp

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ChartInfo struct {
	Repo    *string `yaml:"repo"`
	Chart   *string `yaml:"chart"`
	Version *string `yaml:"version"`
}

type HelmProps struct {
	ChartInfo           ChartInfo
	ChartFileNamePrefix *string
	ReleaseName         *string
	Namespace           *string
	Values              *map[string]interface{}
}

func NewHelmCached(scope constructs.Construct, id *string, props *HelmProps) cdk8s.Helm {
	globals := GetGlobalContext(scope)
	chartsCacheDir := fmt.Sprintf("%s/%s", globals.CacheDir, "charts")
	if err := os.MkdirAll(chartsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("MkdirAll failed", err)
	}
	if _, err := exec.LookPath("helm"); err != nil {
		log.Fatalln("helm LookPath failed", err)
	}
	if props.ChartInfo.Repo == nil {
		log.Fatalf("props.ChartInfo is nil for %s", *props.ReleaseName)
	}
	chartFileName := *props.ChartInfo.Chart + "-" + *props.ChartInfo.Version + ".tgz"
	if props.ChartFileNamePrefix != nil {
		chartFileName = *props.ChartFileNamePrefix + *props.ChartInfo.Version + ".tgz"
	}
	chartPath := fmt.Sprintf("%s/%s", chartsCacheDir, chartFileName)
	if _, err := os.Stat(chartPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cmd := exec.Command("helm", "pull", *props.ChartInfo.Chart, "--repo", *props.ChartInfo.Repo, "--destination", chartsCacheDir, "--version", *props.ChartInfo.Version)
			if out, err := cmd.CombinedOutput(); err != nil {
				log.Println("Error occured during helm pull command", string(out))
				log.Fatalln("Error occured during helm pull command", err)
			} else {
				log.Printf("Fetching chart '%s' from repo '%s' version '%s' ...", *props.ChartInfo.Chart, *props.ChartInfo.Repo, *props.ChartInfo.Version)
				if len(out) > 0 {
					log.Println(string(out))
				}
			}
		} else {
			log.Fatalln("helm Stat failed", err)
		}
	}

	return cdk8s.NewHelm(scope, id, &cdk8s.HelmProps{
		Chart:       jsii.String(chartPath),
		ReleaseName: props.ReleaseName,
		Namespace:   props.Namespace,
		Values:      props.Values,
		HelmFlags:   jsii.Strings("--include-crds", "--skip-tests", "--no-hooks"),
	})
}