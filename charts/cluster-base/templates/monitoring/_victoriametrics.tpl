# https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-single
{{ define "cluster-base.monitoring.victoriametrics" }}
{{- with .Values.monitoring.victoriametrics }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: victoriametrics
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://victoriametrics.github.io/helm-charts
  chart: victoria-metrics-single
  version: "0.8.51"
  targetNamespace: monitoring
  valuesContent: |-
    server:
      retentionPeriod: {{ .retentionPeriod }}
      statefulSet:
        service:
          annotations:
            prometheus.io/port: "8428"
            prometheus.io/scrape: "true"
      extraArgs:
        vmalert.proxyURL: http://vmalert-victoria-metrics-alert-server:8880

      {{- with .resources }}
      resources:
        {{- . | toYaml | nindent 8 }}
      {{- end }}

      {{- with .persistentVolume }}
      persistentVolume:
        {{- . | toYaml | nindent 8 }}
      {{- end }}

      ingress:
        enabled: {{ .ingress.enabled }}
        annotations:
          {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
          {{ include "cluster-base.ingress.annotation.router-middlewares" $ }}
        hosts:
        - name: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
          path: /
          port: http
        tls:
        - secretName: victoriametrics-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
        pathType: Prefix
{{- end }}
{{- end }}
