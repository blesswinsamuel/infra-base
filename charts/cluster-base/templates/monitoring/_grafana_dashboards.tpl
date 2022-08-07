# https://github.com/blesswinsamuel/helm-charts/blob/main/charts/grafana-dashboards/values.yaml
{{- define "cluster-base.monitoring.grafana-dashboards.create" -}}
{{- range $path, $bytes := .Files }}
{{- $name := base $path }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: grafana-dashboard-{{ $.NamePrefix }}-{{ $name | replace "." "-" }}
  namespace: {{ $.Namespace }}
  labels:
    grafana_dashboard: "1"
  annotations:
    grafana_folder: "{{ $.Folder }}"
data:
  {{ $name }}: |-
{{ $.Files.Get $path | indent 4 }}
{{- end }}
{{- end }}

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
