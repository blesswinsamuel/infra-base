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
  version: "16.8.9"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    image:
      repository: redis
      tag: 6
    architecture: standalone
    auth:
      enabled: false
{{- end }}
{{- end }}
