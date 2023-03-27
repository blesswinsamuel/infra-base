# https://github.com/blesswinsamuel/helm-charts/blob/main/charts/alerting-rules/values.yaml
{{- define "cluster-base.monitoring.default-rules" }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: alerting-rules
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://blesswinsamuel.github.io/helm-charts
  chart: alerting-rules
  version: "0.0.2"
  targetNamespace: monitoring
{{- end }}
