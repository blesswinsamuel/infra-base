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
  version: "0.0.8"
  targetNamespace: '{{ tpl $.Values.system.namespace $ }}'
  valuesContent: |-
    deployment:
      annotations:
        secret.reloader.stakater.com/reload: "helm-ops"
    gitRepo:
      url: "{{ .gitRepo.url }}"
      branch: "{{ .gitRepo.branch }}"
    helmChartPath: "{{ .helmChartPath }}"
    helmReleaseName: "{{ .helmReleaseName }}"
    helmReleaseNamespace: "{{ .helmReleaseNamespace }}"
    helmExtraValuesFiles: {{ .helmExtraValuesFiles | toJson }}
    clusterName: "test"
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
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  data:
    - secretKey: git-private-key
      remoteRef: { key: GITHUB_DEPLOY_KEY }
    - secretKey: known_hosts
      remoteRef: { key: GITHUB_KNOWN_HOSTS }
{{- end }}
{{- end }}
