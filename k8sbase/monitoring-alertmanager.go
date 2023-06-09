package k8sbase

import (
	_ "embed"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8simports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

//go:embed alertmanager-templates.tpl
var alertmanagerTemplates string

type AlertmanagerProps struct {
	Enabled bool             `yaml:"enabled"`
	Image   k8sapp.ImageInfo `yaml:"image"`
	Ingress struct {
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

	chart := k8sapp.NewApplicationChart(scope, "alertmanager", &k8sapp.ApplicationProps{
		Kind:                         "StatefulSet",
		Name:                         "alertmanager",
		AutomountServiceAccountToken: true,
		PodSecurityContext: &k8s.PodSecurityContext{
			FsGroup: jsii.Number(65534),
		},
		ServiceAccountName:     "alertmanager",
		CreateHeadlessService:  true,
		StatefulSetServiceName: "alertmanager-headless",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "alertmanager",
			Image: props.Image,
			Args: []string{
				"--storage.path=/alertmanager",
				"--config.file=/etc/alertmanager/alertmanager.yml",
			},
			ExtraEnvs: []*k8s.EnvVar{{Name: jsii.String("POD_IP"), ValueFrom: &k8s.EnvVarSource{FieldRef: &k8s.ObjectFieldSelector{ApiVersion: jsii.String("v1"), FieldPath: jsii.String("status.podIP")}}}},
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 9093, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
			},
			LivenessProbe:  &k8s.Probe{HttpGet: &k8s.HttpGetAction{Port: k8s.IntOrString_FromString(jsii.String("http")), Path: jsii.String("/")}},
			ReadinessProbe: &k8s.Probe{HttpGet: &k8s.HttpGetAction{Port: k8s.IntOrString_FromString(jsii.String("http")), Path: jsii.String("/")}},
			SecurityContext: &k8s.SecurityContext{
				RunAsGroup:   jsii.Number(65534),
				RunAsUser:    jsii.Number(65534),
				RunAsNonRoot: jsii.Bool(true),
			},
		}},
		Secrets: []k8sapp.ApplicationSecret{{
			Name: "alertmanager-templates",
			Data: map[string]string{
				"templates.tmpl": alertmanagerTemplates,
			},
			MountPath: "/etc/alertmanager/templates",
			MountName: "templates",
			ReadOnly:  true,
		}},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{{
			Name: "alertmanager-config",
			Template: map[string]string{
				"alertmanager.yml": strings.TrimSpace(dedent.String(`
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
				`)),
			},
			RemoteRefs: map[string]string{
				"telegramBotToken":         "TELEGRAM_BOT_TOKEN",
				"telegramChatID":           "TELEGRAM_CHAT_ID",
				"slackApiUrl":              "SLACK_API_URL",
				"alertmanagerWatchdogUrls": "ALERTMANAGER_WATCHDOG_URLS",
			},
			MountPath: "/etc/alertmanager",
			MountName: "config",
			ReadOnly:  true,
		}},
		StatefulSetVolumeClaimTemplates: []k8sapp.ApplicationPersistentVolume{{
			Name:            "storage",
			RequestsStorage: "50Mi",
			MountPath:       "/alertmanager",
			MountName:       "storage",
		}},
	})

	k8s.NewKubeServiceAccount(scope, jsii.String("alertmanager-sa"), &k8s.KubeServiceAccountProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("alertmanager"),
		},
		AutomountServiceAccountToken: jsii.Bool(true),
	})

	return chart
}
