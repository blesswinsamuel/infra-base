# https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-state-metrics
{{ define "cluster-base.monitoring.kube-state-metrics" }}
{{- with .Values.monitoring.kubeStateMetrics }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: kube-state-metrics
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://prometheus-community.github.io/helm-charts
  chart: kube-state-metrics
  version: "5.4.2"
  targetNamespace: monitoring
  valuesContent: |-
    fullnameOverride: kube-state-metrics
    service:
      annotations:
        prometheus.io/port: "8080"
{{- end }}
{{- end }}
