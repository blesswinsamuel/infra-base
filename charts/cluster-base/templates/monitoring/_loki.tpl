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
  version: "5.0.0"
  targetNamespace: monitoring
  valuesContent: |-
    singleBinary:
      replicas: 1
    monitoring:
      dashboards:
        enabled: false
      serviceMonitor:
        enabled: false
        metricsInstance:
          enabled: false
      alerts:
        enabled: false
      rules:
        enabled: false
        alerting: false
      selfMonitoring:
        enabled: false
        grafanaAgent:
          installOperator: false
      lokiCanary:
        enabled: false
    test:
      enabled: false
    gateway:
      enabled: false
        
    memberlist:
      service:
        # https://github.com/grafana/loki/issues/7907#issuecomment-1445336799
        publishNotReadyAddresses: true

    loki:
      auth_enabled: false
      commonConfig:
        replication_factor: 1
      compactor:
        retention_enabled: true
      rulerConfig:
        alertmanager_url: http://vmalert-alertmanager:9093
      {{- if eq .storage "local" }}
      storage:
        type: 'filesystem'
      {{- else if eq .storage "s3" }}
      storage:
        type: 's3'
        bucketNames:
          chunks: loki-chunks
          ruler: loki-ruler
          admin: loki-admin  # never used
        s3:
          endpoint: {{ .s3.endpoint }}
          secretAccessKey: {{ .s3.secret_access_key }}
          accessKeyId: {{ .s3.access_key_id }}
          s3ForcePathStyle: true
          # insecure: true
      {{- else }}
      {{- fail "Invalid storage type" }}
      {{- end }}
{{- end }}
{{- end }}
