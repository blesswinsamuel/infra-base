package k8sbase

import (
	_ "embed"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

//go:embed alertmanager-templates.tpl
var alertmanagerTemplates string

type AlertmanagerProps struct {
	Enabled       bool      `yaml:"enabled"`
	HelmChartInfo ChartInfo `yaml:"helm"`
	Ingress       struct {
		Enabled   bool   `yaml:"enabled"`
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
	Config struct {
		Slack struct {
			Channel string `yaml:"channel"`
		} `yaml:"slack"`
		Telegram struct {
			ParseMode string `yaml:"parseMode"`
		} `yaml:"telegram"`
	} `yaml:"config"`
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/alertmanager
// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert

func NewAlertmanager(scope constructs.Construct, props AlertmanagerProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("alertmanager"), &cprops)

	helmOutput := NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.HelmChartInfo,
		ReleaseName: jsii.String("alertmanager"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"service": map[string]interface{}{
				"annotations": map[string]string{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "9093",
				},
			},
			"statefulSet": map[string]interface{}{
				"annotations": map[string]string{
					"secret.reloader.stakater.com/reload": "alertmanager-config",
				},
			},
			"ingress": map[string]interface{}{
				"enabled":     props.Ingress.Enabled,
				"annotations": GetCertIssuerAnnotation(scope),
				"hosts": []map[string]interface{}{
					{
						"host": props.Ingress.SubDomain + "." + GetDomain(scope),
						"paths": []map[string]interface{}{
							{
								"path":     "/",
								"pathType": "Prefix",
							},
						},
					},
				},
				"tls": []map[string]interface{}{
					{
						"hosts": []string{
							props.Ingress.SubDomain + "." + GetDomain(scope),
						},
						"secretName": "alertmanager-tls",
					},
				},
			},
			// # extraArgs:
			// #   web.external-url: "https://{{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}"
			"config": []map[string]interface{}{},
			"extraSecretMounts": []map[string]interface{}{
				{
					"name":       "config",
					"secretName": "alertmanager-config",
					"mountPath":  "/etc/alertmanager",
					"readOnly":   true,
				},
				{
					"name":       "templates",
					"secretName": "alertmanager-templates",
					"mountPath":  "/etc/alertmanager/templates",
					"readOnly":   true,
				},
			},
		},
	})

	for _, obj := range *helmOutput.ApiObjects() {
		if *obj.Metadata().Name() == "alertmanager-test-connection" {
			helmOutput.Node().TryRemoveChild(obj.Node().Id())
		}
	}

	k8s.NewKubeSecret(chart, jsii.String("alertmanager-templates"), &k8s.KubeSecretProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("alertmanager-templates"),
		},
		StringData: &map[string]*string{
			"templates.tmpl": jsii.String(alertmanagerTemplates),
		},
	})

	NewExternalSecret(chart, jsii.String("alertmanager-config"), &ExternalSecretProps{
		Name:            jsii.String("alertmanager-config"),
		RefreshInterval: jsii.String("2m"),
		Template: &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
			Data: &map[string]*string{
				"alertmanager.yml": jsii.String(strings.TrimSpace(dedent.String(`
					global:
					  resolve_timeout: 1m
					
					templates:
					  - '/etc/alertmanager/templates/*.tmpl'
					
					route:
					  receiver: notify-main
					  # tag to group by
					  group_by: ["alertname"]
					  # How long to initially wait to send a notification for a group of alerts
					  group_wait: 30s
					  # How long to wait before sending a notification about new alerts that are added to a group
					  group_interval: 10s
					  # How long to wait before sending a notification again if it has already been sent successfully for an alert
					  repeat_interval: 2h
					  routes:
					  - matchers:
					    - "alertname = Watchdog"
					    receiver: watchdog
					    repeat_interval: 1m
					  - matchers:
					    - "alertname = InfoInhibitor"
					    receiver: devnull
					
					receivers:
					- name: devnull
					- name: watchdog
					  webhook_configs:
					  {{- range (.alertmanagerWatchdogUrls | fromJson) }}
					  - url: {{ . }}
					  {{- end }}
					- name: notify-main
					  slack_configs:
					  - channel: "` + props.Config.Slack.Channel + `"
					    send_resolved: true
					    api_url: {{ .slackApiUrl | quote }}
					    icon_url: 'https://avatars3.githubusercontent.com/u/3380462'
					    title: {{ ` + "`" + `'{{ template "slack.title" . }}'` + "`" + ` }}
					    text: {{ ` + "`" + `'{{ template "slack.text" . }}'` + "`" + ` }}
					  telegram_configs:
					  - api_url: https://api.telegram.org
					    bot_token: {{ .telegramBotToken | quote }}
					    chat_id: {{ .telegramChatID }}
					    message: {{ ` + "`" + `'{{ template "telegram.message" . }}'` + "`" + ` }}
					    parse_mode: "` + props.Config.Telegram.ParseMode + `"
				`))),
			},
		},
		Secrets: map[string]string{
			"telegramBotToken":         "TELEGRAM_BOT_TOKEN",
			"telegramChatID":           "TELEGRAM_CHAT_ID",
			"slackApiUrl":              "SLACK_API_URL",
			"alertmanagerWatchdogUrls": "ALERTMANAGER_WATCHDOG_URLS",
		},
	})

	return chart
}
