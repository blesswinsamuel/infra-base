{{- define "cluster-base.var_dump" -}}
{{- . | mustToPrettyJson | printf "\nThe JSON output of the dumped var is: \n%s" | fail -}}
{{- end -}}

{{- define "commons.annotation.cert-issuer" -}}
{{- if eq .Values.global.certIssuerKind "ClusterIssuer" -}}
cert-manager.io/cluster-issuer: {{ .Values.global.certIssuer }}
{{- else if eq .Values.global.certIssuerKind "Issuer" -}}
cert-manager.io/issuer: {{ .Values.global.certIssuer }}
{{- else -}}
{{- .Values.global.certIssuerKind | printf "Invalid global.certIssuerKind: %s" | fail -}}
{{- end -}}
{{- end -}}

{{- define "commons.annotation.router-middlewares" -}}
{{- if eq $.Values.global.internetAuthType "basic-auth" -}}
traefik.ingress.kubernetes.io/router.middlewares: kube-system-traefik-redirect-https@kubernetescrd,kube-system-traefik-basic-auth@kubernetescrd
{{- else if eq $.Values.global.internetAuthType "traefik-forward-auth" -}}
traefik.ingress.kubernetes.io/router.middlewares: kube-system-traefik-redirect-https@kubernetescrd,auth-traefik-forward-auth@kubernetescrd
{{- else -}}
{{- .Values.global.internetAuthType | printf "Invalid global.internetAuthType: %s" | fail -}}
{{- end -}}
{{- end -}}

{{- define "commons.router-middleware.https-redirect" -}}
kube-system-traefik-redirect-https@kubernetescrd
{{- end -}}

{{- define "commons.router-middleware.auth" -}}
{{- if eq $.Values.global.internetAuthType "basic-auth" -}}
kube-system-traefik-basic-auth@kubernetescrd
{{- else if eq $.Values.global.internetAuthType "traefik-forward-auth" -}}
auth-traefik-forward-auth@kubernetescrd
{{- else -}}
{{- .Values.global.internetAuthType | printf "Invalid global.internetAuthType: %s" | fail -}}
{{- end -}}
{{- end -}}

{{- define "commons.monitoring.grafana.dashboard-cm.tpl" -}}
{{- range $path, $bytes := .Files }}
{{- $name := base $path }}
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
---
{{- end }}
{{- end }}
