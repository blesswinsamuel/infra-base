package kbaseresources

import (
	_ "embed"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog/log"
)

//go:embed values-default.yaml
var defaultValuesBytes []byte

var defaultValues map[string]ast.Node

func init() {
	if err := yaml.UnmarshalWithOptions(defaultValuesBytes, &defaultValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		log.Panic().Err(err).Msg("Unmarshal default values")
	}
}
