{{ define "cluster-base.namespace.create" }}
---
apiVersion: v1
kind: Namespace
metadata:
  name: "{{ . }}"
  labels:
    name: "{{ . }}"
{{ end }}

{{ define "cluster-base.namespaces" }}
  {{- if .Values.global.helmChartsNamespaceCreate }}
    {{- include "cluster-base.namespace.create" (include "cluster-base.namespace.helm-chart" $) }}
  {{- end }}

  {{- range .Values.createNamespaces }}
    {{- include "cluster-base.namespace.create" (tpl . $) }}
  {{- end }}
{{ end }}
