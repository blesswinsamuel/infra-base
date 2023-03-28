# https://github.com/stakater/Reloader/blob/master/deployments/kubernetes/chart/reloader
{{- define "cluster-base.system.reloader" -}}
{{- with .Values.system.reloader }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: reloader
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://stakater.github.io/stakater-charts
  chart: reloader
  version: "v1.0.16"
  targetNamespace: '{{ tpl $.Values.system.namespace $ }}'
  valuesContent: |-
    # reloader:
    #   custom_annotations:
    #     configmap: "my.company.com/configmap"
    #     secret: "my.company.com/secret"
    service:
      port: 9090
      annotations:
        prometheus.io/port: "9090"
        prometheus.io/scrape: "true"
{{- end }}
{{- end }}
