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
  version: "0.0.4"
  targetNamespace: '{{ tpl $.Values.system.namespace $ }}'
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
    scripts:
      predeploy: |
        helm repo add blesswinsamuel https://blesswinsamuel.github.io/helm-charts
        yq e -i '.dependencies[0].repository="https://blesswinsamuel.github.io/helm-charts"' $HELM_CHART_DIR/Chart.yaml
        yq e -i '.dependencies[0].repository="https://blesswinsamuel.github.io/helm-charts"' $HELM_CHART_DIR/Chart.lock

        helm dependency update $HELM_CHART_DIR
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: helm-ops
  namespace: '{{ tpl $.Values.system.namespace $ }}'
  annotations:
    reloader.stakater.com/match: "true"
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    name: helm-ops
  data:
    - secretKey: git-private-key
      remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: GITHUB_DEPLOY_KEY }
    - secretKey: known_hosts
      remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: GITHUB_KNOWN_HOSTS }
{{- end }}
{{- end }}
