package utils

import "gopkg.in/yaml.v3"

func FromYamlString[T any](s string) T {
	var v T
	err := yaml.Unmarshal([]byte(s), &v)
	if err != nil {
		panic(err)
	}
	return v
}
