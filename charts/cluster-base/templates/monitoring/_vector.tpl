# https://github.com/vectordotdev/helm-charts/tree/develop/charts/vector
# https://helm.vector.dev/index.yaml
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
  version: "0.16.3"
  targetNamespace: monitoring
  valuesContent: |-
    role: Agent
    # Prometheus scrape is disabled because it's creating duplicate metrics. Also, there is a peer_addr which blows up the cardinality
    # service:
    #   annotations:
    #     prometheus.io/port: "9090"
    #     prometheus.io/scrape: "true"
    {{- if .ingress.enabled }}
    ingress:
      enabled: {{ .ingress.enabled }}
      annotations:
        {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
        {{ include "cluster-base.ingress.annotation.router-middlewares" $ }}
      hosts:
        - host: {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
          paths:
            - path: /
              pathType: ImplementationSpecific
              port:
                name: api
      tls:
        - secretName: vector-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ tpl $.Values.global.domain $ }}
    {{- end }}
    customConfig:
      data_dir: /vector-data-dir
      api:
        enabled: true
        address: 0.0.0.0:8686
        playground: false
      sources:
        {{- if .syslogServer.enabled }}
        syslog_server:
          type: syslog
          address: 0.0.0.0:514
          max_length: 102400
          mode: tcp
          path: /syslog-socket
        {{- end }}
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
        {{- if .syslogServer.enabled }}
        syslog_transform:
          type: "remap"
          inputs:
            - syslog_server
          source: |
            .kubernetes = {}
            .kubernetes.container_name = .appname
            .kubernetes.pod_name = .appname
            .kubernetes.pod_node_name = .host
            .kubernetes.pod_namespace = "syslog"
            .level = .severity
        {{- end }}
        kubernetes_parse_and_merge_log_message:
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
        kubernetes_log_transform:
          type: "remap"
          inputs:
            - kubernetes_parse_and_merge_log_message
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
          inputs:
            - kubernetes_log_transform
          {{- if .syslogServer.enabled }}
            - syslog_transform
          {{- end }}
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
        # debug_sink:
        #   type: console
        #   inputs:
        #     - syslog_server
        #   target: stdout
        #   encoding:
        #     codec: json
          # healthcheck:
          #   enabled: true
{{- if .syslogServer.enabled }}
---
apiVersion: v1
kind: Service
metadata:
  name: vector-syslog-server
  namespace: monitoring
spec:
  type: NodePort
  ports:
  - name: syslog-server
    port: 514
    protocol: TCP
    targetPort: 514
    nodePort: 30514
  selector:
    app.kubernetes.io/component: Agent
    app.kubernetes.io/instance: vector
    app.kubernetes.io/name: vector
{{- end }}
# Debug command:
# yq e '.spec.valuesContent' vector.yaml | helm template --namespace monitoring --repo https://helm.vector.dev --version 0.10.3 vector vector --values - --debug
{{- end }}
{{- end }}
