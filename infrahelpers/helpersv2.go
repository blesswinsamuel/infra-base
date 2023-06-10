package infrahelpers

import (
	"bytes"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"github.com/goccy/go-yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"time"
)

var scheme = runtime.NewScheme()

func init() {
	if err := corev1.AddToScheme(scheme); err != nil {
		panic(err)
	}
	if err := externalsecretsv1beta1.AddToScheme(scheme); err != nil {
		panic(err)
	}
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

	serializer := json.NewSerializerWithOptions(
		json.DefaultMetaFactory, scheme, scheme,
		json.SerializerOptions{
			Pretty: true,
			Yaml:   true,
			Strict: true,
		},
	)
	b := bytes.NewBuffer(nil)
	if err := serializer.Encode(obj, b); err != nil {
		panic(err)
	}
	//fmt.Println(string(b.Bytes()))

	var out map[string]any
	if err := yaml.Unmarshal(b.Bytes(), &out); err != nil {
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
