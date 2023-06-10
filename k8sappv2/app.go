package main

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-yaml"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
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
	app := NewApp(AppProps{})
	chart := NewChart(app, "test", ChartProps{
		Namespace: "test",
	})
	//n := corev1.Namespace{
	//	ObjectMeta: metav1.ObjectMeta{},
	//	Spec:       corev1.NamespaceSpec{},
	//}
	//fmt.Println(n.GetObjectKind())
	chart.AddObject(corev1.Namespace{
		TypeMeta:   metav1.TypeMeta{APIVersion: "v1", Kind: "Namespace"},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       corev1.NamespaceSpec{},
	})
	chart.AddObject(v1.DaemonSet{
		TypeMeta:   metav1.TypeMeta{APIVersion: "apps/v1", Kind: "DaemonSet"},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: v1.DaemonSetSpec{
			Selector: nil,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:                       "1",
					GenerateName:               "",
					Namespace:                  "",
					UID:                        "",
					ResourceVersion:            "",
					Generation:                 0,
					CreationTimestamp:          metav1.Time{},
					DeletionTimestamp:          nil,
					DeletionGracePeriodSeconds: nil,
					Labels:                     nil,
					Annotations:                nil,
					OwnerReferences:            nil,
					Finalizers:                 nil,
					ManagedFields:              nil,
				},
				Spec: corev1.PodSpec{},
			},
			UpdateStrategy:       v1.DaemonSetUpdateStrategy{},
			MinReadySeconds:      0,
			RevisionHistoryLimit: nil,
		},
	})

	app.Synth()
}
