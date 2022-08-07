{{ define "cluster-base.system" }}
  {{- include "cluster-base.namespace.create" "system" }}
  
  {{- if .Values.system.reloader.enabled }}
    {{- include "cluster-base.system.reloader" . }}
  {{- end }}
  
  {{- if .Values.system.helmOps.enabled }}
    {{- include "cluster-base.system.helm-ops" . }}
  {{- end }}
{{ end }}
