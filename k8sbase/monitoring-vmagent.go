package k8sbase

import (
	_ "embed"

	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
)

type VmagentProps struct {
	Enabled            bool                        `json:"enabled"`
	HelmChartInfo      k8sapp.ChartInfo            `json:"helm"`
	ExtraScrapeConfigs []map[string]any            `json:"extraScrapeConfigs"`
	Resources          corev1.ResourceRequirements `json:"resources"`
	Ingress            struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

//go:embed vmagent-config.yaml
var vmagentConfig string

// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
func NewVmagent(scope packager.Construct, props VmagentProps) packager.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("vmagent", cprops)

	vmagentConfig := infrahelpers.FromYamlString[map[string]any](vmagentConfig)

	k8sapp.NewHelm(chart, jsii.String("helm"), &k8sapp.HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("vmagent"),
		Namespace:   chart.Namespace(),
		Values: map[string]any{
			"remoteWriteUrls": []string{
				"http://victoriametrics-victoria-metrics-single-server:8428/api/v1/write",
			},
			"extraScrapeConfigs": props.ExtraScrapeConfigs,
			"resources":          props.Resources,
			"extraArgs": map[string]any{
				"promscrape.streamParse": "true",
			},
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
