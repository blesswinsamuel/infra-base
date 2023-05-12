package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type MonitoringProps struct {
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
	construct := constructs.NewConstruct(scope, jsii.String("monitoring"))

	NewNamespace(construct, "monitoring")

	NewGrafana(construct, props.Grafana)
	NewKubeStateMetrics(construct, props.KubeStateMetrics)
	NewGrafanaDashboards(construct, props.GrafanaDashboards)
	NewAlertingRules(construct, props.AlertingRules)
	NewNodeExporter(construct, props.NodeExporter)
	NewVector(construct, props.Vector)
	NewVictoriaMetrics(construct, props.Victoriametrics)
	NewAlertmanager(construct, props.Alertmanager)
	NewVmagent(construct, props.Vmagent)
	NewVmalert(construct, props.Vmalert)
	NewLoki(construct, props.Loki)
	NewCrowdsec(construct, props.Crowdsec)
	NewCrowdsecTraefikBouncer(construct, props.CrowdsecTraefikBouncer)

	NewGrafanaDatasource(construct, GrafanaDatasourceProps{})

	return construct
}
