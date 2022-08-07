{{ define "cluster-base.system" }}
  {{- include "cluster-base.namespace.create" (tpl .Values.system.namespace $) }}
  
  {{- if .Values.system.reloader.enabled }}
    {{- include "cluster-base.system.reloader" . }}
  {{- end }}
  
  {{- if .Values.system.helmOps.enabled }}
    {{- include "cluster-base.system.helm-ops" . }}
  {{- end }}

  {{- if .Values.system.kubernetesDashboard.enabled }}
    {{- include "cluster-base.system.kubernetes-dashboard" . }}
  {{- end }}
{{ end }}
