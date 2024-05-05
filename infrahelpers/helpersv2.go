package infrahelpers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	externalsecretsv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"github.com/goccy/go-yaml"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/apimachinery/pkg/runtime"
	k8sjson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var yamlSerializer *k8sjson.Serializer
var jsonSerializer *k8sjson.Serializer
var localSchemeBuilder = runtime.SchemeBuilder{
	scheme.AddToScheme,
	apiextensionsv1.AddToScheme,
	externalsecretsv1beta1.AddToScheme,
	certmanagerv1.AddToScheme,
	traefikv1alpha1.AddToScheme,
}

func init() {
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(localSchemeBuilder.AddToScheme(Scheme))

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

func YamlToK8sObject(data []byte) runtime.Object {
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

	obj, _, err := Codecs.UniversalDeserializer().Decode(data, nil, nil)
	if err != nil {
		// log.Println(string(data))
		panic(err)
	}
	return obj
}

func K8sObjectToYaml(obj runtime.Object) []byte {
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
	err := yamlSerializer.Encode(obj, b)
	if err != nil {
		// log.Println(string(data))
		panic(err)
	}
	return b.Bytes()
}

func ToDuration(duration string) *metav1.Duration {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return &metav1.Duration{Duration: dur}
}

type MergeableMap[K comparable, V any] map[K]V

func (m *MergeableMap[K, V]) MergeMap(other map[K]V) {
	if *m == nil {
		*m = make(map[K]V)
	}
	for k, v := range other {
		(*m)[k] = v
	}
}

func (m *MergeableMap[K, V]) UnmarshalYAML(ctx context.Context, data []byte) error {
	// workaround to allow merging of maps
	var goMap map[K]V
	if err := yaml.UnmarshalContext(ctx, data, &goMap, yaml.Strict()); err != nil {
		return err
	}
	m.MergeMap(goMap)
	return nil
}

// https://github.com/goccy/go-yaml/issues/425
type YAMLRawMessage []byte

func (m YAMLRawMessage) MarshalJSON() ([]byte, error) { return m.marshal() }
func (m YAMLRawMessage) MarshalYAML() ([]byte, error) { return m.marshal() }

func (m YAMLRawMessage) marshal() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *YAMLRawMessage) UnmarshalJSON(data []byte) error { return m.unmarshal(data) }
func (m *YAMLRawMessage) UnmarshalYAML(data []byte) error { return m.unmarshal(data) }

func (m *YAMLRawMessage) unmarshal(data []byte) error {
	if m == nil {
		return errors.New("RawMessage: unmarshal on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

type YAMLAllowInclude[T any] struct{ V T }

// // UnmarshalYAML is a custom unmarshaler for go-yaml.
// func (m *YAMLAllowInclude[T]) UnmarshalYAML(ctx context.Context, unmarshal func(interface{}) error) error {
// 	var raw string
// 	if err := unmarshal(&raw); err != nil {
// 		// not !include
// 		// fmt.Println("err", err)
// 	}
// 	// fmt.Println("raw", raw)
// 	if strings.HasPrefix(raw, "^include") {
// 		// !include
// 		filePath := strings.TrimSpace(strings.TrimPrefix(raw, "^include"))
// 		fileBytes, err := os.Open(filePath)
// 		if err != nil {
// 			return err
// 		}
// 		var v T
// 		decoder := yaml.NewDecoder(fileBytes, yaml.Strict())
// 		if err = decoder.Decode(&v); err != nil {
// 			return err
// 		}
// 		m.value = v
// 	} else {
// 		// not !include
// 		var v T
// 		if err := unmarshal(&v); err != nil {
// 			return err
// 		}
// 		m.value = v
// 	}
// 	return nil
// }

func (m *YAMLAllowInclude[T]) UnmarshalYAML(ctx context.Context, data []byte) error {
	if m == nil {
		return fmt.Errorf("UnmarshalYAML on nil pointer")
	}
	if bytes.HasPrefix(data, []byte("^include")) {
		// !include
		filePath := strings.TrimSpace(strings.TrimPrefix(string(data), "^include"))
		fileBytes, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		data = fileBytes
	} else {
		// not !include
	}
	var v T
	decoder := yaml.NewDecoder(bytes.NewReader(data), yaml.Strict())
	if err := decoder.DecodeContext(ctx, &v); err != nil {
		return err
	}
	(*m).V = v
	return nil
}

func (m YAMLAllowInclude[T]) MarshalYAML() (interface{}, error) {
	return m.V, nil
}
