package k8sapp

import (
	"github.com/blesswinsamuel/kgen"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigmapProps struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
	Data        map[string]string
}

func NewConfigMap(scope kgen.Scope, props *ConfigmapProps) kgen.ApiObject {
	configMap := v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:        props.Name,
			Labels:      props.Labels,
			Annotations: props.Annotations,
		},
		Data: props.Data,
	}

	return scope.AddApiObject(&configMap)
}
