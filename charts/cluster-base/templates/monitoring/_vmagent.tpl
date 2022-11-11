{{- define "vmagent.config" }}
# Copied from https://github.com/VictoriaMetrics/helm-charts/blob/master/charts/victoria-metrics-agent/values.yaml
# 1. remove labelmap
# 2. Change job=kubernetes-apiservers to job=apiserver
# 3. Change job=kubernetes-nodes-cadvisor to job=kubelet and metrics_path=/metrics/cadvisor
# 4. Change job=kubernetes-nodes to job=kubelet and metrics_path=/metrics
# 5. Set honor_labels: true (https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config) to use labels from scraped metrics instead of adding exported_ prefix
# 6. Set global scrape interval to 1m
global:
  scrape_interval: 1m

scrape_configs:
- job_name: vmagent
  static_configs:
    - targets: ["localhost:8429"]

  ## COPY from Prometheus helm chart https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml

  # Scrape config for API servers.
- job_name: "kubernetes-apiservers"
  kubernetes_sd_configs:
    - role: endpoints
  scheme: https
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    insecure_skip_verify: true
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
  honor_labels: true
  relabel_configs:
    - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
      action: keep
      regex: default;kubernetes;https
    - target_label: job
      replacement: apiserver
- job_name: "kubernetes-nodes"
  scheme: https
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    insecure_skip_verify: true
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
  kubernetes_sd_configs:
    - role: node
  honor_labels: true
  relabel_configs:
    # - action: labelmap
    #   regex: __meta_kubernetes_node_label_(.+)
    - target_label: __address__
      replacement: kubernetes.default.svc:443
    - source_labels: [__meta_kubernetes_node_name]
      regex: (.+)
      target_label: __metrics_path__
      replacement: /api/v1/nodes/$1/proxy/metrics
    - target_label: job
      replacement: kubelet
    - target_label: metrics_path
      replacement: /metrics
- job_name: "kubernetes-nodes-cadvisor"
  scheme: https
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    insecure_skip_verify: true
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
  kubernetes_sd_configs:
    - role: node
  honor_labels: true
  relabel_configs:
    # - action: labelmap
    #   regex: __meta_kubernetes_node_label_(.+)
    - target_label: __address__
      replacement: kubernetes.default.svc:443
    - source_labels: [__meta_kubernetes_node_name]
      regex: (.+)
      target_label: __metrics_path__
      replacement: /api/v1/nodes/$1/proxy/metrics/cadvisor
    - target_label: job
      replacement: kubelet
    - target_label: metrics_path
      replacement: /metrics/cadvisor

# Scrape config for service endpoints.
- job_name: "kubernetes-service-endpoints"
  kubernetes_sd_configs:
    - role: endpoints
  honor_labels: true
  relabel_configs:
    - action: drop
      source_labels: [__meta_kubernetes_pod_container_init]
      regex: true
    - action: keep_if_equal
      source_labels: [__meta_kubernetes_service_annotation_prometheus_io_port, __meta_kubernetes_pod_container_port_number]
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_scrape]
      action: keep
      regex: true
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_scheme]
      action: replace
      target_label: __scheme__
      regex: (https?)
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_path]
      action: replace
      target_label: __metrics_path__
      regex: (.+)
    - source_labels:
        [
          __address__,
          __meta_kubernetes_service_annotation_prometheus_io_port,
        ]
      action: replace
      target_label: __address__
      regex: ([^:]+)(?::\d+)?;(\d+)
      replacement: $1:$2
    # - action: labelmap
    #   regex: __meta_kubernetes_service_label_(.+)
    - source_labels: [__meta_kubernetes_pod_name]
      target_label: pod
    - source_labels: [__meta_kubernetes_pod_container_name]
      target_label: container
    - source_labels: [__meta_kubernetes_namespace]
      target_label: namespace
    - source_labels: [__meta_kubernetes_service_name]
      target_label: service
    - source_labels: [__meta_kubernetes_service_name]
      target_label: job
      replacement: ${1}
    - source_labels: [__meta_kubernetes_pod_node_name]
      action: replace
      target_label: node
# Scrape config for slow service endpoints; same as above, but with a larger
# timeout and a larger interval
- job_name: "kubernetes-service-endpoints-slow"
  scrape_interval: 5m
  scrape_timeout: 30s
  kubernetes_sd_configs:
    - role: endpoints
  honor_labels: true
  relabel_configs:
    - action: drop
      source_labels: [__meta_kubernetes_pod_container_init]
      regex: true
    - action: keep_if_equal
      source_labels: [__meta_kubernetes_service_annotation_prometheus_io_port, __meta_kubernetes_pod_container_port_number]
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_scrape_slow]
      action: keep
      regex: true
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_scheme]
      action: replace
      target_label: __scheme__
      regex: (https?)
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_path]
      action: replace
      target_label: __metrics_path__
      regex: (.+)
    - source_labels:
        [
          __address__,
          __meta_kubernetes_service_annotation_prometheus_io_port,
        ]
      action: replace
      target_label: __address__
      regex: ([^:]+)(?::\d+)?;(\d+)
      replacement: $1:$2
    # - action: labelmap
    #   regex: __meta_kubernetes_service_label_(.+)
    - source_labels: [__meta_kubernetes_pod_name]
      target_label: pod
    - source_labels: [__meta_kubernetes_pod_container_name]
      target_label: container
    - source_labels: [__meta_kubernetes_namespace]
      target_label: namespace
    - source_labels: [__meta_kubernetes_service_name]
      target_label: service
    - source_labels: [__meta_kubernetes_service_name]
      target_label: job
      replacement: ${1}
    - source_labels: [__meta_kubernetes_pod_node_name]
      action: replace
      target_label: node
# Example scrape config for probing services via the Blackbox Exporter.
- job_name: "kubernetes-services"
  metrics_path: /probe
  params:
    module: [http_2xx]
  kubernetes_sd_configs:
    - role: service
  honor_labels: true
  relabel_configs:
    - source_labels:
        [__meta_kubernetes_service_annotation_prometheus_io_probe]
      action: keep
      regex: true
    - source_labels: [__address__]
      target_label: __param_target
    - target_label: __address__
      replacement: blackbox
    - source_labels: [__param_target]
      target_label: instance
    # - action: labelmap
    #   regex: __meta_kubernetes_service_label_(.+)
    - source_labels: [__meta_kubernetes_namespace]
      target_label: namespace
    - source_labels: [__meta_kubernetes_service_name]
      target_label: service
# Example scrape config for pods
- job_name: "kubernetes-pods"
  kubernetes_sd_configs:
    - role: pod
  honor_labels: true
  relabel_configs:
    - action: drop
      source_labels: [__meta_kubernetes_pod_container_init]
      regex: true
    - action: keep_if_equal
      source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_port, __meta_kubernetes_pod_container_port_number]
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
      action: keep
      regex: true
    - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
      action: replace
      target_label: __metrics_path__
      regex: (.+)
    - source_labels:
        [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
      action: replace
      regex: ([^:]+)(?::\d+)?;(\d+)
      replacement: $1:$2
      target_label: __address__
    # - action: labelmap
    #   regex: __meta_kubernetes_pod_label_(.+)
    - source_labels: [__meta_kubernetes_pod_name]
      target_label: pod
    - source_labels: [__meta_kubernetes_pod_container_name]
      target_label: container
    - source_labels: [__meta_kubernetes_namespace]
      target_label: namespace
    - source_labels: [__meta_kubernetes_service_name]
      target_label: service
    - source_labels: [__meta_kubernetes_service_name]
      target_label: job
      replacement: ${1}
    - source_labels: [__meta_kubernetes_pod_node_name]
      action: replace
      target_label: node
  ## End of COPY
{{- end }}

{{ define "cluster-base.monitoring.vmagent" }}
{{- with .Values.monitoring.vmagent }}
# https://github.com/VictoriaMetrics/helm-charts/tree/master/charts/victoria-metrics-agent
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: vmagent
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://victoriametrics.github.io/helm-charts
  chart: victoria-metrics-agent
  version: "0.8.21"
  targetNamespace: monitoring
  valuesContent: |-
    remoteWriteUrls:
      - http://victoriametrics-victoria-metrics-single-server:8428/api/v1/write
    {{- with .extraScrapeConfigs }}
    extraScrapeConfigs:
      {{- . | toYaml | nindent 6 }}
    {{- end }}
    {{- with .resources }}
    resources:
      {{- . | toYaml | nindent 6 }}
    {{- end }}
    # extraArgs:
    #  http.pathPrefix: /vmagent
    service:
      enabled: true
    ingress:
      enabled: true
      annotations:
        {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
        {{ include "cluster-base.ingress.annotation.router-middlewares" $ }}
      hosts:
      - name: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
        path: /
        port: http
      tls:
      - secretName: vmagent-tls
        hosts:
          - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
      pathType: Prefix
    config:
      {{- include "vmagent.config" $ | nindent 6 }}
#     extraArgs:
#       remoteWrite.relabelConfig: /relabel-config/relabel-config.yaml
#     extraVolumeMounts:
#     - mountPath: /relabel-config/relabel-config.yaml
#       subPath: relabel-config.yaml
#       name: relabel-config
#       readOnly: true
#     extraVolumes:
#     - name: relabel-config
#       configMap:
#         name: vmagent-relabel-config
# ---
# apiVersion: v1
# kind: ConfigMap
# metadata:
#   name: vmagent-relabel-config
#   namespace: monitoring
# data:
#   relabel-config.yaml: |
#     - action: labeldrop
#       regex: "(node_role_kubernetes_io_|node_kubernetes_io_|beta_kubernetes_io_|kubernetes_io_|app_kubernetes_io_|helm_sh_).+"
#     - action: labeldrop
#       regex: "(chart|heritage|release|pod_template_hash|objectset_rio_cattle_io_hash)"
#     - source_labels: job
#       action: replace
#       regex: kubernetes-nodes
#       replacement: apiserver
{{- end }}
{{- end }}
