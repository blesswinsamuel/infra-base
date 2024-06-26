package infrahelpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ToK8sDuration(duration string) *metav1.Duration {
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
