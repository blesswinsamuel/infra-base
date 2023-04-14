# https://github.com/grafana/helm-charts/blob/main/charts/grafana
{{ define "cluster-base.monitoring.grafana" }}
{{- with .Values.monitoring.grafana }}
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: grafana
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://grafana.github.io/helm-charts
  chart: grafana
  version: "6.52.9"
  targetNamespace: monitoring
  valuesContent: |-
    env:
      GF_SERVER_ENABLE_GZIP: true

      {{- if .anonymousAuthEnabled }}
      GF_AUTH_ANONYMOUS_HIDE_VERSION: true
      GF_AUTH_ANONYMOUS_ENABLED: true
      GF_AUTH_ANONYMOUS_ORG_NAME: Main Org.
      GF_AUTH_ANONYMOUS_ORG_ROLE: Admin
      GF_AUTH_DISABLE_LOGIN_FORM: true
      {{- end }}

      {{- if .authProxyEnabled }}
      GF_AUTH_PROXY_ENABLED: true
      GF_AUTH_PROXY_HEADER_NAME: Remote-User
      GF_AUTH_PROXY_HEADER_PROPERTY: username
      GF_AUTH_PROXY_AUTO_SIGN_UP: true
      GF_AUTH_PROXY_HEADERS: "Groups:Remote-Group"
      GF_AUTH_PROXY_ENABLE_LOGIN_TOKEN: false
      {{- end }}
    podAnnotations:
      prometheus.io/port: "3000"
      prometheus.io/scrape: "true"
    ingress:
      enabled: true
      annotations:
        {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
      hosts:
        - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
      tls:
        - secretName: grafana-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
    sidecar:
      datasources:
        enabled: true
        label: {{ .datasourceLabel }}
        labelValue: {{ .datasourceLabelValue }}
        resource: configmap
      dashboards:
        enabled: true
        label: {{ .dashboardLabel }}
        labelValue: {{ .dashboardLabelValue }}
        resource: configmap
        folderAnnotation: grafana_folder
        provider:
          foldersFromFilesStructure: true
    rbac:
      namespaced: {{ .namespaced }}
      pspEnabled: false
  # - name: "My Provider"
  #   orgId: 1
  #   folder: "My folder"
  #   folderUid: ""
  #   type: file
  #   disableDeletion: false
  #   editable: true
  #   updateIntervalSeconds: 10
  #   allowUiUpdates: true
  #   options:
  #     path: /etc/grafana/dashboards
  #     foldersFromFilesStructure: true
{{- end }}
{{- end }}
