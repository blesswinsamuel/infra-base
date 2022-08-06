# https://github.com/grafana/helm-charts/blob/main/charts/loki/values.yaml
{{ define "cluster-base.monitoring.loki" }}
{{- with .Values.monitoring.loki }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: loki
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://grafana.github.io/helm-charts
  chart: loki
  version: "2.12.0"
  targetNamespace: monitoring
  valuesContent: |-
    config:
      compactor:
        retention_enabled: true
      ruler:
        # storage:
        #   type: local  # todo: s3
        #   local:
        #     directory: /rules
        # rule_path: /tmp/scratch
        alertmanager_url: http://vmalert-alertmanager:9093
        # ring:
        #   kvstore:
        #     store: inmemory
        # enable_api: true
{{- end }}
{{- end }}
