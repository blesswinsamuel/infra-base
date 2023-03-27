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
  version: "12.2.6"
  targetNamespace: '{{ tpl .namespace $ }}'
  valuesContent: |-
    nameOverride: postgres
    image:
      # not straightforward to upgrade to postgres 15
      # 2023-03-18 16:30:09.021 GMT [1] FATAL:  database files are incompatible with server
      # 2023-03-18 16:30:09.021 GMT [1] DETAIL:  The data directory was initialized by PostgreSQL version 14, which is not compatible with this version 15.2.
      repository: bitnami/postgresql
      tag: 14.7.0
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
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    name: postgres-passwords
  data:
  - secretKey: postgres-password
    remoteRef: { key: POSTGRES_ADMIN_PASSWORD }
  - secretKey: password
    remoteRef: { key: POSTGRES_USER_PASSWORD }
  - secretKey: replication-password
    remoteRef: { key: POSTGRES_REPLICATION_PASSWORD }
{{- end }}
{{- end }}
