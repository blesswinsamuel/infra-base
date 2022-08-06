{{ define "cluster-base.database" }}
  {{- if .Values.databases.enabled }}
    {{- include "cluster-base.namespace.create" "database" }}
    
    {{- if .Values.databases.postgres.enabled }}
      {{- include "cluster-base.database.postgres" . }}
      {{- if .Values.databases.postgres.backup.enabled }}
        {{- include "cluster-base.database.postgres.backup-job" . }}
      {{- end }}
    {{- end }}

    {{- if .Values.databases.redis.enabled }}
      {{- include "cluster-base.database.redis" . }}
    {{- end }}
  {{- end }}
{{ end }}
