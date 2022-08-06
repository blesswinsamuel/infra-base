{{- define "cluster-base.all" }}

  {{- /* Merge the local chart values and the common chart defaults */ -}}
  {{- include "cluster-base.values.setup" . }}

  {{- /* Create secret resources */ -}}
  {{- include "cluster-base.secrets" . }}

  {{- /* Ingress */ -}}
  {{- include "cluster-base.ingress" . }}

  {{- /* Create namespaces */ -}}
  {{- include "cluster-base.namespaces" . }}

  {{- /* System */ -}}
  {{- include "cluster-base.system" . }}

  {{- /* Monitoring */ -}}
  {{- include "cluster-base.monitoring" . }}

  {{- /* Database */ -}}
  {{- include "cluster-base.database" . }}

  {{- /* Auth */ -}}
  {{- include "cluster-base.auth" . }}

{{- end }}
