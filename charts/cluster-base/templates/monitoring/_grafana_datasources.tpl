{{ define "cluster-base.monitoring.datasource.victoriametrics" }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: grafana-datasource-victoriametrics
  namespace: monitoring
  labels:
    grafana_datasource: "1"
data:
  victoriametrics.yaml: |-
    apiVersion: 1

    deleteDatasources:
      - name: VictoriaMetrics
        orgId: 1

    datasources:
      - name: VictoriaMetrics
        type: prometheus
        access: proxy
        orgId: 1
        uid: victoriametrics
        url: http://victoriametrics-victoria-metrics-single-server:8428
        isDefault: true
        version: 1
        editable: false
        # jsonData:
        #   alertmanagerUid: alertmanager
{{- end }}

{{ define "cluster-base.monitoring.datasource.loki" }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: grafana-datasource-loki
  namespace: monitoring
  labels:
    grafana_datasource: "1"
data:
  loki.yaml: |-
    apiVersion: 1

    deleteDatasources:
      - name: Loki
        orgId: 1

    datasources:
      - name: Loki
        type: loki
        access: proxy
        orgId: 1
        uid: loki
        url: http://loki:3100
        jsonData:
          maxLines: 1000
          # alertmanagerUid: alertmanager
{{- end }}

{{ define "cluster-base.monitoring.datasource.alertmanager" }}
---
# apiVersion: v1
# kind: ConfigMap
# metadata:
#   name: grafana-datasource-alertmanager
#   namespace: monitoring
#   labels:
#     grafana_datasource: "1"
# data:
#   loki.yaml: |-
#     apiVersion: 1

#     deleteDatasources:
#       - name: Alertmanager
#         orgId: 1

#     datasources:
#       - name: Alertmanager
#         type: alertmanager
#         access: proxy
#         orgId: 1
#         uid: alertmanager
#         url: http://vmalert-alertmanager:9093
#         jsonData:
#           implementation: prometheus
{{- end }}
