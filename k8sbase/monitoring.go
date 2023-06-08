package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type MonitoringProps struct {
	Enabled                bool                        `yaml:"enabled"`
	Grafana                GrafanaProps                `yaml:"grafana"`
	GrafanaDashboards      GrafanaDashboardsProps      `yaml:"grafanaDashboards"`
	AlertingRules          AlertingRulesProps          `yaml:"alertingRules"`
	KubeStateMetrics       KubeStateMetricsProps       `yaml:"kubeStateMetrics"`
	NodeExporter           NodeExporterProps           `yaml:"nodeExporter"`
	Vector                 VectorProps                 `yaml:"vector"`
	Victoriametrics        VictoriametricsProps        `yaml:"victoriametrics"`
	Alertmanager           AlertmanagerProps           `yaml:"alertmanager"`
	Vmagent                VmagentProps                `yaml:"vmagent"`
	Vmalert                VmalertProps                `yaml:"vmalert"`
	Loki                   LokiProps                   `yaml:"loki"`
	Crowdsec               CrowdsecProps               `yaml:"crowdsec"`
	CrowdsecTraefikBouncer CrowdsecTraefikBouncerProps `yaml:"crowdsec-traefik-bouncer"`
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
