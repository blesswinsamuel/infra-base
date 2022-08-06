{{- define "cluster-base.system.helm-ops" -}}
{{- with .Values.system.helmOps }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: helm-ops
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://blesswinsamuel.github.io/helm-charts
  chart: helm-ops
  version: "0.0.2"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    deployment:
      annotations:
        reloader.stakater.com/search: "true"
    gitRepo:
      url: "{{ .gitRepo.url }}"
      branch: "{{ .gitRepo.branch }}"
    helmChartPath: "{{ .helmChartPath }}"
    helmReleaseName: "{{ .helmReleaseName }}"
    helmReleaseNamespace: "{{ .helmReleaseNamespace }}"
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: helm-ops
  namespace: "{{ tpl .namespace $ }}"
  annotations:
    reloader.stakater.com/match: "true"
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: secretstore
    kind: ClusterSecretStore
  target:
    name: helm-ops
  data:
    - secretKey: git-private-key
      remoteRef: { key: doppler-secrets, property: BLESS_STACK_DEPLOY_KEY }
    - secretKey: known_hosts
      remoteRef: { key: doppler-secrets, property: GITHUB_KNOWN_HOSTS }
{{- end }}
{{- end }}
