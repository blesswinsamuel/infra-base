# https://github.com/keel-hq/keel/tree/master/chart/keel
{{- define "cluster-base.system.keel" -}}
{{- with .Values.system.keel }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: keel
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://charts.keel.sh
  chart: keel
  version: "0.9.10"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    image:
      repository: blesswinsamuel/docker-keel
      tag: edge
      pullPolicy: Always
    helmProvider:
      enabled: false
    # https://github.com/keel-hq/keel/issues/491#issuecomment-1022326010
    slack:
      enabled: true
      botName: keel_app
      token: xoxb-1234567890123-1234567890123-123456789012abcdefghijkl
      channels: keel
      approvalsChannel: keel
{{- end }}
{{- end }}
