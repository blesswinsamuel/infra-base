# https://github.com/bitnami/charts/tree/master/bitnami/postgresql
{{ define "cluster-base.database.postgres" }}
{{- with .Values.databases.postgres }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: postgres
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://charts.bitnami.com/bitnami
  chart: postgresql
  version: "11.1.22"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    nameOverride: postgres
    image:
      repository: postgres
      tag: 14.2
    auth:
      {{- with .database }}
      database: {{ . }}
      {{- end }}
      {{- with .username }}
      username: {{ . }}
      {{- end }}
      existingSecret: postgres-passwords
    metrics:
      enabled: true
      image:
        repository: prometheuscommunity/postgres-exporter
        tag: v0.10.1
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: postgres-passwords
  namespace: '{{ tpl .namespace $ }}'
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: secretstore
    kind: ClusterSecretStore
  target:
    name: postgres-passwords
  data:
  - secretKey: postgres-password
    remoteRef: { key: doppler-secrets, property: POSTGRES_ADMIN_PASSWORD }
  - secretKey: password
    remoteRef: { key: doppler-secrets, property: POSTGRES_USER_PASSWORD }
  - secretKey: replication-password
    remoteRef: { key: doppler-secrets, property: POSTGRES_REPLICATION_PASSWORD }
{{- end }}
{{- end }}
