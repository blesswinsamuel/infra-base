# https://github.com/kubernetes/dashboard/tree/master/aio/deploy/helm-chart/kubernetes-dashboard
{{- define "cluster-base.system.kubernetes-dashboard" -}}
{{- with .Values.system.kubernetesDashboard }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: kubernetes-dashboard
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://kubernetes.github.io/dashboard/
  chart: kubernetes-dashboard
  version: "5.7.0"
  targetNamespace: '{{ tpl $.Values.system.namespace $ }}'
  valuesContent: |-
    protocolHttp: true
    extraArgs:
      - --enable-skip-login
      - --enable-insecure-login
    service:
      externalPort: 9090
    rbac:
      create: false
      # clusterReadOnlyRole: true
    serviceAccount:
      create: false
      name: kubernetes-dashboard
    metricsScraper:
      enabled: true
    ingress:
      enabled: {{ .ingress.enabled }}
      annotations:
        {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
        {{ include "cluster-base.ingress.annotation.router-middlewares" $ }}
        # traefik.ingress.kubernetes.io/service.serversscheme: https
      hosts:
        - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
      tls:
        - secretName: kubernetes-dashboard-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kubernetes-dashboard
  namespace: '{{ tpl $.Values.system.namespace $ }}'
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: kubernetes-dashboard
    namespace: '{{ tpl $.Values.system.namespace $ }}'
# ---
# apiVersion: v1
# kind: Secret
# metadata:
#   name: kubernetes-dashboard
#   namespace: '{{ tpl $.Values.system.namespace $ }}'
#   annotations:
#     kubernetes.io/service-account.name: kubernetes-dashboard
# type: kubernetes.io/service-account-token
{{- end }}
{{- end }}
