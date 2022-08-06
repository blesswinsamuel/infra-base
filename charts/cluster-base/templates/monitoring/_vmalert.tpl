# https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-alert
{{ define "cluster-base.monitoring.vmalert" }}
{{- with .Values.monitoring.vmalert }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: vmalert
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://victoriametrics.github.io/helm-charts
  chart: victoria-metrics-alert
  version: "0.4.35"
  targetNamespace: monitoring
  valuesContent: |-
    {{- with $.Values.monitoring.alertmanager }}
    alertmanager:
      tag: v0.24.0
      enabled: {{ .enabled }}
      service:
        annotations:
          prometheus.io/port: "9093"
          prometheus.io/scrape: "true"
      ingress:
        enabled: {{ .ingress.enabled }}
        annotations:
          {{ include "commons.annotation.cert-issuer" $ }}
          {{ include "commons.annotation.router-middlewares" $ }}
        hosts:
        - name: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
          path: /
          port: http
        tls:
        - secretName: alertmanager-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
        pathType: Prefix
      baseURL: https://{{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
      config:
        global:
          resolve_timeout: 1m
          slack_api_url: {{ .config.slackApiUrl | quote }}

        route:
          receiver: devnull
          # receiver: 'slack-notifications'     
          routes:
          # - matchers:
          #   - "alertname = Watchdog"
          #   receiver: Watchdog
          - matchers:
            - "alertname = InfoInhibitor"
            receiver: 'devnull'
          - matchers:
            - "severity = critical|warning"
            receiver: 'slack-notifications'

        receivers:
        - name: devnull
        - name: 'slack-notifications'
          slack_configs:
          - channel: {{ .config.channel | quote }}
            send_resolved: true
            icon_url: 'https://avatars3.githubusercontent.com/u/3380462'
            title: {{ .config.title | quote }}
            text: {{ .config.text | quote }}
    {{- end }}
    server:
      service:
        annotations:
          prometheus.io/port: "8880"
          prometheus.io/scrape: "true"
      ingress:
        enabled: {{ .ingress.enabled }}
        annotations:
          {{ include "commons.annotation.cert-issuer" $ }}
          {{ include "commons.annotation.router-middlewares" $ }}
        hosts:
        - name: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
          path: /
          port: http
        tls:
        - secretName: vmalert-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
        pathType: Prefix
      configMap: alerting-rules
      datasource:
        url: http://victoriametrics-victoria-metrics-single-server:8428
      remote:
        write:
          url: http://victoriametrics-victoria-metrics-single-server:8428
        read:
          url: http://victoriametrics-victoria-metrics-single-server:8428
      extraArgs:
        rule: '/config/*.yaml'
        external.url: https://grafana.{{ tpl $.Values.global.domain $ }}
        # external.alert.source: explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":""},{"mode":"Metrics"},{"ui":[true,true,true,"none"]}]
        # external.alert.source: {{ `explore?orgId=1&left=["now-1h","now","VictoriaMetrics",{"expr":"{{$expr|quotesEscape|pathEscape}}"}]` }}
        # https://github.com/VictoriaMetrics/VictoriaMetrics/blob/8edb390e215cbffe9bb267eea8337dbf1df1c76f/deployment/docker/docker-compose.yml#L75
        external.alert.source: {{ `explore?orgId=1&left={"datasource":"VictoriaMetrics","queries":[{"expr":"{{$expr|quotesEscape|crlfEscape|queryEscape}}","refId":"A"}],"range":{"from":"now-1h","to":"now"}}` }}
      # - "-external.label=env=${ENV_NAME}"
      # - "-evaluationInterval=30s"
      # - "-rule=/config/alert_rules.yml"
      {{- with .resources }}
      resources:
        {{- . | toYaml | nindent 6 }}
      {{- end }}
{{- end }}
{{- end }}
{{- /*
Debug command:
helm template . -n staging --set grafana.datasources.enabled=false --set grafana.dashboards.enabled=false --set monitoring.enabled=true --set monitoring.vmalert.enabled=true --set monitoring.alertmanager.enabled=true \
  | yq e 'select(.metadata.name == "vmalert" and .kind == "HelmChart")' \
  | yq e '.spec.valuesContent' \
  | helm template --namespace monitoring --repo https://victoriametrics.github.io/helm-charts --version 0.4.31 vmalert victoria-metrics-alert --values - --debug
*/ -}}