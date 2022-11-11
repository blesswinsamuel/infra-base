# https://github.com/grafana/loki/tree/main/production/helm/loki
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
  version: "3.3.4"
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
