package k8sbase

import (
	_ "embed"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/muesli/reflow/dedent"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//go:embed alertmanager-templates.tpl
var alertmanagerTemplates string

type AlertmanagerProps struct {
	Image   k8sapp.ImageInfo `json:"image"`
	Ingress struct {
		Enabled   bool   `json:"enabled"`
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	Persistence struct {
		StorageClass string `json:"storageClass"`
		VolumeName   string `json:"volumeName"`
	} `json:"persistence"`
	Config struct {
		Slack struct {
			Channel string `json:"channel"`
		} `json:"slack"`
		Telegram struct {
			ParseMode string `json:"parseMode"`
		} `json:"telegram"`
	} `json:"config"`
}

// https://github.com/prometheus-community/helm-charts/blob/main/charts/alertmanager
// https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert

func (props *AlertmanagerProps) Chart(scope kubegogen.Scope) kubegogen.Scope {
	chart := k8sapp.NewApplicationChart(scope, "alertmanager", &k8sapp.ApplicationProps{
		Kind:                         "StatefulSet",
		Name:                         "alertmanager",
		AutomountServiceAccountToken: true,
		PodSecurityContext: &corev1.PodSecurityContext{
			FSGroup: infrahelpers.Ptr(int64(65534)),
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
				"--web.external-url=https://" + props.Ingress.SubDomain + "." + GetDomain(scope),
				// "--log.level=debug",
			},
			ExtraEnvs: []corev1.EnvVar{{Name: "POD_IP", ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{APIVersion: "v1", FieldPath: "status.podIP"}}}},
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 9093, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
			},
			LivenessProbe:  &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/"}}},
			ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/"}}},
			SecurityContext: &corev1.SecurityContext{
				RunAsGroup:   infrahelpers.Ptr(int64(65534)),
				RunAsUser:    infrahelpers.Ptr(int64(65534)),
				RunAsNonRoot: infrahelpers.Ptr(true),
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
					  {{- if .slackApiUrl }}
					  slack_configs:
					  - channel: "` + props.Config.Slack.Channel + `"
					    send_resolved: true
					    api_url: {{ .slackApiUrl | quote }}
					    icon_url: 'https://avatars3.githubusercontent.com/u/3380462'
					    title: {{ ` + "`" + `'{{ template "slack.title" . }}'` + "`" + ` }}
					    text: {{ ` + "`" + `'{{ template "slack.text" . }}'` + "`" + ` }}
					  {{- end }}
					  {{- if and .telegramBotToken .telegramChatID }}
					  telegram_configs:
					  - api_url: https://api.telegram.org
					    bot_token: {{ .telegramBotToken | quote }}
					    chat_id: {{ .telegramChatID }}
					    message: {{ ` + "`" + `'{{ template "telegram.message" . }}'` + "`" + ` }}
					    parse_mode: "` + props.Config.Telegram.ParseMode + `"
					  {{- end }}
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
			RequestsStorage: "40Mi",
			MountPath:       "/alertmanager",
			MountName:       "storage",
			StorageClass:    props.Persistence.StorageClass,
			VolumeName:      props.Persistence.VolumeName,
		}},
	})

	scope.AddApiObject(&corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name: "alertmanager",
		},
		AutomountServiceAccountToken: infrahelpers.Ptr(true),
	})

	return chart
}

// {{ define "cluster-base.monitoring.datasource.alertmanager" }}
// ---
// # apiVersion: v1
// # kind: ConfigMap
// # metadata:
// #   name: grafana-datasource-alertmanager
// #   namespace: monitoring
// #   labels:
// #     grafana_datasource: "1"
// # data:
// #   loki.yaml: |-
// #     apiVersion: 1

// #     deleteDatasources:
// #       - name: Alertmanager
// #         orgId: 1

// #     datasources:
// #       - name: Alertmanager
// #         type: alertmanager
// #         access: proxy
// #         orgId: 1
// #         uid: alertmanager
// #         url: http://vmalert-alertmanager:9093
// #         jsonData:
// #           implementation: prometheus
// {{- end }}
