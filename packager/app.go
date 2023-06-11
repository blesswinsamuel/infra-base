package packager

import (
	"log"
	"os"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type App interface {
	Construct
	OutDir() string
	Synth()
}

type AppProps struct {
	Outdir       string
	CacheDir     string
	DeleteOutDir bool
}

type app struct {
	construct
	props AppProps
}

func NewApp(props AppProps) App {
	return &app{
		props: props,
	}
}

func (a *app) Synth() {
	if a.props.DeleteOutDir {
		if err := os.RemoveAll(a.props.Outdir); err != nil {
			log.Fatalf("RemoveAll: %v", err)
		}
	}
}

func (a *app) OutDir() string {
	return a.props.Outdir
}

type cdk8sapp struct {
	cdk8sConstruct
	app cdk8s.App

	props AppProps
}

func NewCdk8sApp(props AppProps) App {
	appProps := &cdk8s.AppProps{
		Outdir:         jsii.String(props.Outdir),
		YamlOutputType: cdk8s.YamlOutputType_FILE_PER_CHART,
		// YamlOutputType: cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
		// RecordConstructMetadata: jsii.Bool(true),
	}
	app := cdk8s.NewApp(appProps)
	return &cdk8sapp{
		cdk8sConstruct: cdk8sConstruct{app},
		app:            app,
		props:          props,
	}
}

func (a *cdk8sapp) Synth() {
	if a.props.DeleteOutDir {
		if err := os.RemoveAll(a.props.Outdir); err != nil {
			log.Fatalf("RemoveAll: %v", err)
		}
	}

	a.app.Synth()
}

func (a *cdk8sapp) OutDir() string {
	return *a.app.Outdir()
}
