package main

import (
	"bytes"
	"fmt"

	"github.com/goccy/go-yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GenericObject interface {
	//GetNamespace() string
	//SetNamespace(namespace string)
	//GetName() string
	//GetKind() string
	metav1.Object
	metav1.Type
}

type Scope interface {
	AddObject(obj GenericObject)
}

type App struct {
	charts []*Chart
}

type AppProps struct {
}

func (a *App) Synth() {
	for _, chart := range a.charts {
		chart.Render()
	}
}

func NewApp(props AppProps) *App {
	return &App{}
}

func (a *App) AddChart(chart *Chart) {
	a.charts = append(a.charts, chart)
}

type Chart struct {
	id        string
	namespace string
	objects   []any
}

type ChartProps struct {
	Namespace string
}

func NewChart(a *App, id string, props ChartProps) *Chart {
	chart := &Chart{
		id:        id,
		namespace: props.Namespace,
	}
	a.AddChart(chart)
	return chart
}

func (c *Chart) Render() {
	fmt.Println("Chart: ", c.id)
	for _, obj := range c.objects {
		switch obj := obj.(type) {
		case *Chart:
			obj.Render()
			continue
		}
		b := bytes.NewBuffer(nil)
		enc := yaml.NewEncoder(b)
		err := enc.Encode(obj)
		if err != nil {
			panic(err)
		}
		fmt.Println("---")
		fmt.Print(b.String())
	}
}

func (c *Chart) AddObject(obj any) {
	c.objects = append(c.objects, obj)
}

func main() {

	// -- break --

	// valuesFiles := []string{
	// 	"values.yaml",
	// }
	// values := k8sapp.LoadValues(valuesFiles, nil)
	// props := kgen.AppProps{
	// 	Outdir:       outputDir,
	// 	DeleteOutDir: true,
	// 	// https://pracucci.com/kubernetes-dns-resolution-ndots-options-and-why-it-may-affect-application-performances.html
	// 	// to avoid polluting the dns logs on blocky with search domain home.local. E.g.: postgres.database.svc.cluster.local.home.local
	// 	PatchNdots: false,
	// }
	// app := k8sapp.NewApp(props)
	// k8sbase.SetGlobalContext(app, values.Global)
	// app.SetContext("environment", environment)
	// app.SetContext("SpecialRootDomainSubdomainSeparator", templateVars["SpecialRootDomainSubdomainSeparator"])

	// k8sapp.Render(app, values)

	// k8sapp.Synth(app)
	// log.Printf("Done in %s.", time.Since(startTime))

	// log.Printf("values:\n%v", prettyPrint(values))

	// outputDir := "k8s-generated"
	// props := kgen.AppProps{
	// 	Outdir:       outputDir,
	// 	DeleteOutDir: true,
	// 	PatchNdots:   true,
	// }
	// app := k8sapp.NewApp(props)
	// k8sbase.SetGlobalContext(app, values.Global)

	// k8sapp.Render(app, values)

	// k8sapp.Synth(app)

	// -- break --

	// app := NewApp(AppProps{})
	// chart := NewChart(app, "test", ChartProps{
	// 	Namespace: "test",
	// })
	// //n := corev1.Namespace{
	// //	ObjectMeta: metav1.ObjectMeta{},
	// //	Spec:       corev1.NamespaceSpec{},
	// //}
	// //fmt.Println(n.GetObjectKind())
	// chart.AddObject(corev1.Namespace{
	// 	// TypeMeta:   metav1.TypeMeta{APIVersion: "v1", Kind: "Namespace"},
	// 	ObjectMeta: metav1.ObjectMeta{},
	// 	Spec:       corev1.NamespaceSpec{},
	// })
	// chart.AddObject(v1.DaemonSet{
	// 	// TypeMeta:   metav1.TypeMeta{APIVersion: "apps/v1", Kind: "DaemonSet"},
	// 	ObjectMeta: metav1.ObjectMeta{},
	// 	Spec: v1.DaemonSetSpec{
	// 		Selector: nil,
	// 		Template: corev1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{
	// 				Name:                       "1",
	// 				GenerateName:               "",
	// 				Namespace:                  "",
	// 				UID:                        "",
	// 				ResourceVersion:            "",
	// 				Generation:                 0,
	// 				CreationTimestamp:          metav1.Time{},
	// 				DeletionTimestamp:          nil,
	// 				DeletionGracePeriodSeconds: nil,
	// 				Labels:                     nil,
	// 				Annotations:                nil,
	// 				OwnerReferences:            nil,
	// 				Finalizers:                 nil,
	// 				ManagedFields:              nil,
	// 			},
	// 			Spec: corev1.PodSpec{},
	// 		},
	// 		UpdateStrategy:       v1.DaemonSetUpdateStrategy{},
	// 		MinReadySeconds:      0,
	// 		RevisionHistoryLimit: nil,
	// 	},
	// })

	// app.Synth()
}
