{{ define "cluster-base.auth" }}
  {{- include "cluster-base.namespace.create" .Values.traefikForwardAuth.namespace }}
  {{- include "cluster-base.auth.traefik-forward-auth" . }}
{{ end }}
