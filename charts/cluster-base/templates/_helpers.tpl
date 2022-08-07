{{- define "cluster-base.var-dump" -}}
{{- . | mustToPrettyJson | printf "\nThe JSON output of the dumped var is: \n%s" | fail -}}
{{- end -}}

{{- define "cluster-base.ingress.annotation.cert-issuer" -}}
{{- include "cluster-base.values.setup" . -}}
cert-manager.io/cluster-issuer: {{ .Values.global.clusterCertIssuerName }}
{{- end -}}

{{- define "cluster-base.ingress.annotation.router-middlewares" -}}
traefik.ingress.kubernetes.io/router.middlewares: {{ include "cluster-base.ingress.router-middleware.https-redirect" $ }},{{ include "cluster-base.ingress.router-middleware.auth" $ }}
{{- end -}}

{{/* HTTPS redirect middleware name */}}
{{- define "cluster-base.ingress.router-middleware.https-redirect" -}}
kube-system-traefik-redirect-https@kubernetescrd
{{- end -}}

{{/* Auth middleware name */}}
{{- define "cluster-base.ingress.router-middleware.auth" -}}
{{- include "cluster-base.values.setup" . -}}
{{- if eq $.Values.global.internetAuthType "basic-auth" -}}
kube-system-traefik-basic-auth@kubernetescrd
{{- else if eq $.Values.global.internetAuthType "traefik-forward-auth" -}}
auth-traefik-forward-auth@kubernetescrd
{{- else -}}
{{- .Values.global.internetAuthType | printf "Invalid global.internetAuthType: %s" | fail -}}
{{- end -}}
{{- end -}}

{{/* Merge the local chart values and the common chart defaults */}}
{{- define "cluster-base.values.setup" -}}
  {{- if hasKey .Values "cluster-base" -}}
    {{- $defaultValues := deepCopy (get .Values "cluster-base") -}}
    {{- $userValues := deepCopy (omit .Values "cluster-base") -}}
    {{- $mergedValues := mustMergeOverwrite $defaultValues $userValues -}}
    {{- $_ := set . "Values" (deepCopy $mergedValues) -}}
  {{- end -}}
{{- end -}}

{{/* HelmChart namespace */}}
{{- define "cluster-base.namespace.helm-chart" -}}
{{- include "cluster-base.values.setup" . -}}
{{- tpl $.Values.global.helmChartsNamespace $ -}}
{{- end -}}
