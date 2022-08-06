{{ define "cluster-base.ingress" }}
  {{- if .Values.certIssuer.enabled }}
    {{- include "cluster-base.ingress.certissuer" . }}
  {{- end }}
  {{- include "cluster-base.ingress.traefik" . }}
{{ end }}
