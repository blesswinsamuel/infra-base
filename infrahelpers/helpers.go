package infrahelpers

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/goccy/go-yaml"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
	return buf.String()
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

func FromJSONString[T any](s string) T {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}

func GoTemplate(s string, data interface{}) string {
	tmpl, err := template.New("tmpl").Parse(s)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func Ptr[T any](v T) *T {
	return &v
}

func MergeMaps[K comparable, V any](annotations ...map[K]V) map[K]V {
	merged := make(map[K]V)
	for _, annotation := range annotations {
		for k, v := range annotation {
			merged[k] = v
		}
	}
	return merged
}

func MergeAnnotations(annotations ...map[string]string) map[string]string {
	merged := make(map[string]string)
	for _, annotation := range annotations {
		for k, v := range annotation {
			merged[k] = v
		}
	}
	return merged
}

func MergeLists[T any](annotations ...[]T) []T {
	merged := make([]T, 0)
	for _, annotation := range annotations {
		merged = append(merged, annotation...)
	}
	return merged
}

func Ternary[V any](cond bool, trueVal V, falseVal V) V {
	if cond {
		return trueVal
	}
	return falseVal
}
