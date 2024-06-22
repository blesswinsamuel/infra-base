package k8sapp

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"github.com/rs/zerolog/log"
	"golang.org/x/xerrors"
)

type ValuesGlobalDefaults struct {
	SecretStoreName               string `json:"secretStoreName"`
	SecretStoreKind               string `json:"secretStoreKind"`
	ExternalSecretRefreshInterval string `json:"externalSecretRefreshInterval"`

	CertIssuerName string `json:"certIssuerName"`
	CertIssuerKind string `json:"certIssuerKind"`
}

type ValuesGlobal struct {
	Domain          string               `json:"domain"`
	ClusterName     string               `json:"clusterName"`
	DisableTls      bool                 `json:"disableTls"`
	SecretsProvider string               `json:"secretsProvider"`
	Defaults        ValuesGlobalDefaults `json:"defaults"`
}

type Values struct {
	Global   ValuesGlobal                                     `json:"global"`
	Services OrderedMap[string, OrderedMap[string, ast.Node]] `json:"services"`
}

func GetGlobals(scope kubegogen.Scope) ValuesGlobal {
	return scope.GetContext("global").(ValuesGlobal)
}

func LoadValues(valuesFiles []string, templateMap map[string]any) Values {
	var values Values
	// default global values
	values.Global = ValuesGlobal{
		Domain:          "",
		SecretsProvider: "doppler", // or 1password
		Defaults: ValuesGlobalDefaults{
			SecretStoreName:               "secretstore",
			SecretStoreKind:               "ClusterSecretStore",
			ExternalSecretRefreshInterval: "10m",
			CertIssuerName:                "letsencrypt-prod",
			CertIssuerKind:                "ClusterIssuer",
		},
	}

	valuesMerged := map[string]interface{}{}
	for _, valuesFile := range valuesFiles {
		log.Info().Msgf("Loading values from %s", valuesFile)
		valuesFileBytes, err := os.ReadFile(valuesFile)
		if err != nil {
			log.Panic().Err(err).Msg("ReadFile")
		}
		if templateMap != nil {
			out, err := template.New("tpl").Delims("[{@", "@}]").Parse(string(valuesFileBytes))
			if err != nil {
				log.Panic().Err(err).Msg("template Parse")
			}
			w := bytes.NewBuffer([]byte{})
			if err := out.Execute(w, templateMap); err != nil {
				log.Panic().Err(err).Msg("template Execute")
			}
			valuesFileBytes = w.Bytes()
		}
		fileValues := map[string]interface{}{}
		if err := yaml.UnmarshalWithOptions(valuesFileBytes, &fileValues, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
			printErrIfPretty(err)
			log.Panic().Err(err).Msg("Unmarshal")
		}
		valuesMerged = mergeMaps(valuesMerged, fileValues)
	}
	valuesNode, err := yaml.ValueToNode(valuesMerged)
	if err != nil {
		printErrIfPretty(err)
		log.Panic().Err(err).Msg("ValueToNode")
	}
	if err := yaml.NodeToValue(valuesNode, &values, yaml.Strict(), yaml.UseJSONUnmarshaler()); err != nil {
		printErrIfPretty(err)
		log.Panic().Err(err).Msg("NodeToValue")
	}
	return values
}

type Sink struct{ *bytes.Buffer }

func (es *Sink) Print(args ...interface{}) {
	fmt.Fprint(es.Buffer, args...)
}

func (es *Sink) Printf(f string, args ...interface{}) {
	fmt.Fprintf(es.Buffer, f, args...)
}

func (es *Sink) Detail() bool {
	return false
}

func printErrIfPretty(err error) {
	var prettyError interface {
		PrettyPrint(p xerrors.Printer, colored, inclSource bool) error
	}
	if errors.As(err, &prettyError) {
		var buf bytes.Buffer
		if err := prettyError.PrettyPrint(&Sink{&buf}, true, true); err != nil {
			log.Panic().Err(err).Msg("PrettyPrint")
		}
		fmt.Println(buf.String())
	}
}
