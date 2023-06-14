package k8sapp

import (
	"github.com/blesswinsamuel/infra-base/packager"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretProps struct {
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	StringData  map[string]string
	Data        map[string][]byte
}

func NewSecret(scope packager.Construct, id string, props *SecretProps) packager.ApiObject {
	configMap := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        props.Name,
			Namespace:   props.Namespace,
			Labels:      props.Labels,
			Annotations: props.Annotations,
		},
		StringData: props.StringData,
		Data:       props.Data,
	}

	return NewK8sObject(scope, id, &configMap)
}
