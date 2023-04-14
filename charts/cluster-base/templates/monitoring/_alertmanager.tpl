# https://github.com/prometheus-community/helm-charts/blob/main/charts/alertmanager
# https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert
{{ define "cluster-base.monitoring.alertmanager" }}
{{- with .Values.monitoring.alertmanager }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: alertmanager
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://prometheus-community.github.io/helm-charts
  chart: alertmanager
  version: "0.29.0"
  targetNamespace: monitoring
  valuesContent: |-
    service:
      annotations:
        prometheus.io/port: "9093"
        prometheus.io/scrape: "true"
    statefulSet:
      annotations:
        secret.reloader.stakater.com/reload: "alertmanager-config"
    ingress:
      enabled: {{ .ingress.enabled }}
      annotations:
        {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
      hosts:
      - host: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
        paths:
          - path: /
            pathType: Prefix
      tls:
      - secretName: alertmanager-tls
        hosts:
          - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
    # extraArgs:
    #   web.external-url: "https://{{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}"
    config: null
    extraSecretMounts:
    - name: config
      secretName: alertmanager-config
      mountPath: /etc/alertmanager
      readOnly: true
    - name: templates
      secretName: alertmanager-templates
      mountPath: /etc/alertmanager/templates
      readOnly: true
---
apiVersion: v1
kind: Secret
metadata:
  name: alertmanager-templates
  namespace: monitoring
stringData:
  templates.tmpl: |
    {{`{{- define "slack.title" -}}`}}
    {{ .config.slack.title | nindent 4 }}
    {{`{{- end -}}`}}
    {{`{{- define "slack.text" -}}`}}
    {{ .config.slack.text | nindent 4 }}
    {{`{{- end -}}`}}
    {{`{{- define "telegram.message" -}}`}}
    {{ .config.telegram.message | nindent 4 }}
    {{`{{- end -}}`}}
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: alertmanager-config
  namespace: monitoring
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    template:
      data:
        alertmanager.yml: |
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
            {{`{{- range (.alertmanagerWatchdogUrls | fromJson) }}
            - url: {{ . }}
            {{- end }}`}}
          - name: notify-main
            slack_configs:
            - channel: {{ .config.slack.channel | quote }}
              send_resolved: true
              api_url: {{`{{ .slackApiUrl | quote }}`}}
              icon_url: 'https://avatars3.githubusercontent.com/u/3380462'
              title: {{ printf `{{%s'{{ template "slack.title" . }}'%s}}` "`" "`" }}
              text: {{ printf `{{%s'{{ template "slack.text" . }}'%s}}` "`" "`" }}
            telegram_configs:
            - api_url: https://api.telegram.org
              bot_token: {{`{{ .telegramBotToken | quote }}`}}
              chat_id: {{`{{ .telegramChatID }}`}}
              message: {{ printf `{{%s'{{ template "telegram.message" . }}'%s}}` "`" "`" }}
              parse_mode: {{ .config.telegram.parseMode | quote }}
  data:
  - secretKey: telegramBotToken
    remoteRef:
      key: TELEGRAM_BOT_TOKEN
  - secretKey: telegramChatID
    remoteRef:
      key: TELEGRAM_CHAT_ID
  - secretKey: slackApiUrl
    remoteRef:
      key: SLACK_API_URL
  - secretKey: alertmanagerWatchdogUrls
    remoteRef:
      key: ALERTMANAGER_WATCHDOG_URLS
{{- end }}
{{- end }}
