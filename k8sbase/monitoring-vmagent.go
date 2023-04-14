package k8sbase

import (
	_ "embed"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type Resources struct {
	Requests struct {
		CPU    string `yaml:"cpu"`
		Memory string `yaml:"memory"`
	}
	Limits struct {
		CPU    string `yaml:"cpu"`
		Memory string `yaml:"memory"`
	}
}

type VmagentProps struct {
	Enabled            bool             `yaml:"enabled"`
	HelmChartInfo      ChartInfo        `yaml:"helm"`
	ExtraScrapeConfigs []map[string]any `yaml:"extraScrapeConfigs"`
	Resources          *Resources       `yaml:"resources"`
	Ingress            struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
}

//go:embed vmagent-config.yaml
var vmagentConfig string

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func NewVmagent(scope constructs.Construct, props VmagentProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("vmagent"), &cprops)

	vmagentConfig := FromYamlString[map[string]any](vmagentConfig)

	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("vmagent"),
		Namespace:   chart.Namespace(),
		Values: &map[string]any{
			"remoteWriteUrls": []string{
				"http://victoriametrics-victoria-metrics-single-server:8428/api/v1/write",
			},
			"extraScrapeConfigs": props.ExtraScrapeConfigs,
			"resources":          props.Resources,
			// # extraArgs:
			// #  http.pathPrefix: /vmagent
			"service": map[string]any{
				"enabled": true,
			},
			"ingress": map[string]any{
				"enabled":     props.Ingress.Enabled,
				"annotations": GetCertIssuerAnnotation(scope),
				"hosts": []map[string]any{
					{
						"host": props.Ingress.SubDomain + "." + GetDomain(scope),
						"path": "/",
						"port": "http",
					},
				},
				"tls": []map[string]any{
					{
						"secretName": "vmagent-tls",
						"hosts": []string{
							props.Ingress.SubDomain + "." + GetDomain(scope),
						},
						"pathType": "Prefix",
					},
				},
			},
			"config": vmagentConfig,
		},
	})

	return chart
}

// #     extraArgs:
// #       remoteWrite.relabelConfig: /relabel-config/relabel-config.yaml
// #     extraVolumeMounts:
// #     - mountPath: /relabel-config/relabel-config.yaml
// #       subPath: relabel-config.yaml
// #       name: relabel-config
// #       readOnly: true
// #     extraVolumes:
// #     - name: relabel-config
// #       configMap:
// #         name: vmagent-relabel-config
// # ---
// # apiVersion: v1
// # kind: ConfigMap
// # metadata:
// #   name: vmagent-relabel-config
// #   namespace: monitoring
// # data:
// #   relabel-config.yaml: |
// #     - action: labeldrop
// #       regex: "(node_role_kubernetes_io_|node_kubernetes_io_|beta_kubernetes_io_|kubernetes_io_|app_kubernetes_io_|helm_sh_).+"
// #     - action: labeldrop
// #       regex: "(chart|heritage|release|pod_template_hash|objectset_rio_cattle_io_hash)"
// #     - source_labels: job
// #       action: replace
// #       regex: kubernetes-nodes
// #       replacement: apiserver