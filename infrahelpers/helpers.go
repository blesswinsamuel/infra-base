package infrahelpers

import (
	"bytes"
	"encoding/json"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"text/template"

	"github.com/aws/jsii-runtime-go"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

func FromYamlString[T any](s string) T {
	var v T
	err := yaml.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}

func ToYamlString(v any) string {
	buf := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(buf)
	enc.SetIndent(2)
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func PtrMap[K comparable, V any](m map[K]V) *map[K]*V {
	if len(m) == 0 {
		return nil
	}
	out := make(map[K]*V)
	for k, v := range m {
		v := v
		out[k] = &v
	}
	return &out
}

func PtrSlice[V any](ss ...V) *[]*V {
	out := make([]*V, 0, len(ss))
	for _, s := range ss {
		s := s
		out = append(out, &s)
	}
	return &out
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	out := make(map[K]V)
	for k, v := range m {
		out[k] = v
	}
	return out
}

func PtrIfNonEmpty[V comparable](s V) *V {
	var empty V
	if s != empty {
		return &s
	}
	return nil
}

func PtrIfLenGt0[T any](s []T) *[]T {
	if len(s) > 0 {
		return &s
	}
	return nil
}

func If[V any](cond bool, trueVal V, falseVal V) V {
	if cond {
		return trueVal
	}
	return falseVal
}

func UseOrDefault[V comparable](val V, def V) V {
	var empty V
	if val == empty {
		return def
	}
	return val
}

func UseOrDefaultPtr[V any](val *V, def V) V {
	if val == nil {
		return def
	}
	return *val
}

func MapKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func ToJSONString(v any) string {
	out, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func GoTemplate(s string, data interface{}) *string {
	tmpl, err := template.New("tmpl").Parse(s)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return jsii.String(buf.String())
}

func Ptr[T any](v T) *T {
	return &v
}

func MapToEnvVars(m map[string]string) *[]*k8s.EnvVar {
	envVars := make([]*k8s.EnvVar, 0)
	for k, v := range m {
		envVars = append(envVars, &k8s.EnvVar{
			Name:  jsii.String(k),
			Value: jsii.String(v),
		})
	}
	slices.SortFunc(envVars, func(i, j *k8s.EnvVar) bool {
		return *i.Name < *j.Name
	})
	return &envVars
}
