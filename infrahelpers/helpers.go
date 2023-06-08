package infrahelpers

import (
	"bytes"

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

func MapKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0)
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
