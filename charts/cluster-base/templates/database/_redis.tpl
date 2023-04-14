# https://github.com/bitnami/charts/tree/master/bitnami/redis
{{ define "cluster-base.database.redis" }}
{{- with .Values.databases.redis }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: redis
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://charts.bitnami.com/bitnami
  chart: redis
  version: "17.9.4"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    architecture: standalone
    auth:
      enabled: false
    metrics:
      enabled: true
{{- end }}
{{- end }}
