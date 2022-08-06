{{ define "cluster-base.auth" }}
  {{- if .Values.traefikForwardAuth.enabled }}
    {{- include "cluster-base.namespace.create" .Values.traefikForwardAuth.namespace }}
    {{- include "cluster-base.auth.traefik-forward-auth" . }}
  {{- end }}
{{ end }}
