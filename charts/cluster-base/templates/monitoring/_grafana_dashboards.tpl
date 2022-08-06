# https://github.com/blesswinsamuel/helm-charts/blob/main/charts/grafana-dashboards/values.yaml
{{- define "cluster-base.monitoring.grafana-dashboards" }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: grafana-dashboards
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://blesswinsamuel.github.io/helm-charts
  chart: grafana-dashboards
  version: "0.0.4"
  targetNamespace: monitoring
{{- end }}
