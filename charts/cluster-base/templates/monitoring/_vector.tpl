# https://github.com/vectordotdev/helm-charts/tree/develop/charts/vector
{{ define "cluster-base.monitoring.vector" }}
{{- with .Values.monitoring.vector }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: vector
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://helm.vector.dev
  chart: vector
  version: "0.10.3"
  targetNamespace: monitoring
  valuesContent: |-
    role: Agent
    customConfig:
      data_dir: /vector-data-dir
      api:
        enabled: true
        address: 127.0.0.1:8686
        playground: false
      sources:
        kubernetes_logs:
          type: kubernetes_logs
        # vector_logs:
        #   type: internal_logs
        host_metrics:
          filesystem:
            devices:
              excludes: [binfmt_misc]
            filesystems:
              excludes: [binfmt_misc]
            mountPoints:
              excludes: ["*/proc/sys/fs/binfmt_misc"]
          type: host_metrics
        internal_metrics:
          type: internal_metrics
      transforms:
        parse_and_merge_log_message:
          type: "remap"
          inputs:
            - kubernetes_logs
          source: |
            parsed_message, err = parse_json(.message) # ?? parse_common_log(.message) ?? parse_logfmt(.message) # ?? parse_syslog(.message)
            if err == null {
              del(.message)
              ., err = merge(., parsed_message)
              if err != null {
                log("Failed to merge message into log: " + err, level: "error")
              }
            }
        log_transform:
          type: "remap"
          inputs:
            - parse_and_merge_log_message
          source: |
            # .@timestamp = del(.timestamp)
            del(.kubernetes.pod_labels)
            del(.kubernetes.pod_annotations)
            del(.kubernetes.namespace_labels)
            del(.kubernetes.container_id)
            del(.kubernetes.pod_uid)
            del(.kubernetes.pod_ip)
            del(.kubernetes.pod_ips)
            del(.file)
      sinks:
        prom_exporter:
          type: prometheus_exporter
          inputs: [host_metrics, internal_metrics]
          address: 0.0.0.0:9090
        loki_sink:
          type: loki
          inputs: [log_transform]
          endpoint: "http://loki:3100"
          labels:
            container_name: '{{`{{ print "{{ kubernetes.container_name }}" }}`}}'
            pod_name: '{{`{{ print "{{ kubernetes.pod_name }}" }}`}}'
            pod_namespace: '{{`{{ print "{{ kubernetes.pod_namespace }}" }}`}}'
            pod_node_name: '{{`{{ print "{{ kubernetes.pod_node_name }}" }}`}}'
            level: '{{`{{ print "{{ level }}" }}`}}'
          encoding:
            codec: json
            timestamp_format: rfc3339
          # healthcheck:
          #   enabled: true

# Debug command:
# yq e '.spec.valuesContent' vector.yaml | helm template --namespace monitoring --repo https://helm.vector.dev --version 0.10.3 vector vector --values - --debug
{{- end }}
{{- end }}
