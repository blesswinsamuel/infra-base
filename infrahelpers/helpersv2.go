package infrahelpers

import (
	"bytes"
	"encoding/json"
	"time"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var Scheme = runtime.NewScheme()
var yamlSerializer *k8sjson.Serializer
var jsonSerializer *k8sjson.Serializer

func init() {
	if err := corev1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := networkingv1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := externalsecretsv1beta1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := appsv1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := batchv1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := corev1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := externalsecretsv1beta1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := certmanagerv1.AddToScheme(Scheme); err != nil {
		panic(err)
	}
	if err := traefikv1alpha1.AddToScheme(Scheme); err != nil {
		panic(err)
	}

	yamlSerializer = k8sjson.NewSerializerWithOptions(
		k8sjson.DefaultMetaFactory, Scheme, Scheme,
		k8sjson.SerializerOptions{
			Pretty: true,
			Yaml:   true,
			Strict: true,
		},
	)
	jsonSerializer = k8sjson.NewSerializerWithOptions(
		k8sjson.DefaultMetaFactory, Scheme, Scheme,
		k8sjson.SerializerOptions{
			Pretty: true,
			Yaml:   false,
			Strict: true,
		},
	)
}

func K8sObjectToMap(obj runtime.Object) map[string]any {
	//serializer := json.NewSerializerWithOptions(
	//	json.DefaultMetaFactory, nil, nil,
	//	json.SerializerOptions{
	//		Yaml:   true,
	//		Pretty: true,
	//		Strict: true,
	//	},
	//)
	//b := bytes.NewBuffer(nil)
	//err := serializer.Encode(obj, b)
	//if err != nil {
	//	panic(err)
	//}

	//codec := serializer.NewCodecFactory(scheme).LegacyCodec(
	//	corev1.SchemeGroupVersion,
	//	externalsecretsv1beta1.SchemeGroupVersion,
	//)
	//output, err := runtime.Encode(codec, obj)
	//if err != nil {
	//	panic(err)
	//}

	b := bytes.NewBuffer(nil)
	if err := jsonSerializer.Encode(obj, b); err != nil {
		panic(err)
	}
	// fmt.Println(string(b.Bytes()))

	var out map[string]any
	if err := json.Unmarshal(b.Bytes(), &out); err != nil {
		panic(err)
	}
	return out
}

func ToDuration(duration string) *metav1.Duration {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return &metav1.Duration{Duration: dur}
}
