{{ define "cluster-base.monitoring" }}
  {{- if .Values.monitoring.enabled }}
    {{- include "cluster-base.namespace.create" "monitoring" }}

    {{- if .Values.monitoring.defaultRules.enabled }}
      {{- include "cluster-base.monitoring.default-rules" . }}
    {{- end }}

    {{- if .Values.monitoring.grafana.dashboards.monitoring.enabled }}
      {{- include "cluster-base.monitoring.grafana-dashboards" . }}
    {{- end }}

    {{- if .Values.monitoring.victoriametrics.enabled }}
      {{- include "cluster-base.monitoring.datasource.victoriametrics" . }}
      {{- include "cluster-base.monitoring.victoriametrics" . }}
    {{- end }}
    {{- if .Values.monitoring.loki.enabled }}
      {{- include "cluster-base.monitoring.datasource.loki" . }}
      {{- include "cluster-base.monitoring.loki" . }}
    {{- end }}
    {{- if .Values.monitoring.alertmanager.enabled }}
      {{- include "cluster-base.monitoring.datasource.alertmanager" . }}
    {{- end }}

    {{- if .Values.monitoring.grafana.enabled }}
      {{- include "cluster-base.monitoring.grafana" . }}
    {{- end }}

    {{- if .Values.monitoring.kubeStateMetrics.enabled }}
      {{- include "cluster-base.monitoring.kube-state-metrics" . }}
    {{- end }}
    {{- if .Values.monitoring.nodeExporter.enabled }}
      {{- include "cluster-base.monitoring.node-exporter" . }}
    {{- end }}

    {{- if .Values.monitoring.vector.enabled }}
      {{- include "cluster-base.monitoring.vector" . }}
    {{- end }}
    {{- if .Values.monitoring.vmagent.enabled }}
      {{- include "cluster-base.monitoring.vmagent" . }}
    {{- end }}
    {{- if .Values.monitoring.vmalert.enabled }}
      {{- include "cluster-base.monitoring.vmalert" . }}
    {{- end }}
  {{- end }}
{{ end }}
