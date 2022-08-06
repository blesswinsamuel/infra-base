# https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus-node-exporter
{{ define "cluster-base.monitoring.node-exporter" }}
{{- with .Values.monitoring.nodeExporter }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: node-exporter
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://prometheus-community.github.io/helm-charts
  chart: prometheus-node-exporter
  version: "3.3.0"
  targetNamespace: monitoring
  valuesContent: |-
    fullnameOverride: node-exporter
    service:
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "9100"
{{- end }}
{{- end }}
