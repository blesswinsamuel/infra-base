package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type MonitoringProps struct {
	Enabled                bool                        `json:"enabled"`
	Grafana                GrafanaProps                `json:"grafana"`
	GrafanaDashboards      GrafanaDashboardsProps      `json:"grafanaDashboards"`
	AlertingRules          AlertingRulesProps          `json:"alertingRules"`
	KubeStateMetrics       KubeStateMetricsProps       `json:"kubeStateMetrics"`
	NodeExporter           NodeExporterProps           `json:"nodeExporter"`
	Vector                 VectorProps                 `json:"vector"`
	Victoriametrics        VictoriametricsProps        `json:"victoriametrics"`
	Alertmanager           AlertmanagerProps           `json:"alertmanager"`
	Vmagent                VmagentProps                `json:"vmagent"`
	Vmalert                VmalertProps                `json:"vmalert"`
	Loki                   LokiProps                   `json:"loki"`
	Crowdsec               CrowdsecProps               `json:"crowdsec"`
	CrowdsecTraefikBouncer CrowdsecTraefikBouncerProps `json:"crowdsec-traefik-bouncer"`
}

func NewMonitoring(scope constructs.Construct, props MonitoringProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	defer logModuleTiming("monitoring")()
	chart := k8sapp.NewNamespaceChart(scope, "monitoring")

	NewGrafana(chart, props.Grafana)
	NewKubeStateMetrics(chart, props.KubeStateMetrics)
	NewGrafanaDashboards(chart, props.GrafanaDashboards)
	NewAlertingRules(chart, props.AlertingRules)
	NewNodeExporter(chart, props.NodeExporter)
	NewVector(chart, props.Vector)
	NewVictoriaMetrics(chart, props.Victoriametrics)
	NewAlertmanager(chart, props.Alertmanager)
	NewVmagent(chart, props.Vmagent)
	NewVmalert(chart, props.Vmalert)
	NewLoki(chart, props.Loki)
	NewCrowdsec(chart, props.Crowdsec)
	NewCrowdsecTraefikBouncer(chart, props.CrowdsecTraefikBouncer)

	NewGrafanaDatasource(chart, GrafanaDatasourceProps{})

	return chart
}
