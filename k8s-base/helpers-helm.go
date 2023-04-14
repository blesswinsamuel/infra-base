package resourcesbase

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

type ImageInfo struct {
	Repository *string `yaml:"repository"`
	Tag        *string `yaml:"tag"`
}

func (i *ImageInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"repository": i.Repository,
		"tag":        i.Tag,
	}
}

func (i *ImageInfo) ToString() *string {
	return jsii.String(fmt.Sprintf("%s:%s", *i.Repository, *i.Tag))
}

type ChartInfo struct {
	Repo    *string `yaml:"repo"`
	Chart   *string `yaml:"chart"`
	Version *string `yaml:"version"`
}

type HelmProps struct {
	ChartInfo   ChartInfo
	ReleaseName *string
	Namespace   *string
	Values      *map[string]interface{}
}

func NewHelmCached(scope constructs.Construct, id *string, props *HelmProps) cdk8s.Helm {
	chartsCacheDir := fmt.Sprintf("%s/%s", cacheDir, "charts")
	if err := os.MkdirAll(chartsCacheDir, os.ModePerm); err != nil {
		log.Fatalln("MkdirAll failed", err)
	}
	if _, err := exec.LookPath("helm"); err != nil {
		log.Fatalln("helm LookPath failed", err)
	}
	chartPath := fmt.Sprintf("%s/%s-%s.tgz", chartsCacheDir, *props.ChartInfo.Chart, *props.ChartInfo.Version)
	if _, err := os.Stat(chartPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cmd := exec.Command("helm", "pull", *props.ChartInfo.Chart, "--repo", *props.ChartInfo.Repo, "--destination", chartsCacheDir, "--version", *props.ChartInfo.Version)
			if out, err := cmd.CombinedOutput(); err != nil {
				log.Println("Error occured during helm pull command", string(out))
				log.Fatalln("Error occured during helm pull command", err)
			} else {
				log.Printf("Fetching chart '%s' from repo '%s' version '%s' ...", *props.ChartInfo.Chart, *props.ChartInfo.Repo, *props.ChartInfo.Version)
				log.Println(string(out))
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
		HelmFlags:   jsii.Strings("--include-crds"),
	})
}
