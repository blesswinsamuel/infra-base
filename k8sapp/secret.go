package k8sapp

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretProps struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
	StringData  map[string]string
}

func NewSecret(scope constructs.Construct, id *string, props *SecretProps) cdk8s.ApiObject {
	configMap := v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:        props.Name,
			Labels:      props.Labels,
			Annotations: props.Annotations,
		},
		StringData: props.StringData,
	}

	return NewK8sObject(scope, id, &configMap)
}
