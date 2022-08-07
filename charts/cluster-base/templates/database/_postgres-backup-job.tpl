{{ define "cluster-base.database.postgres.backup-job" }}
{{- with .Values.databases.postgres }}
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: postgres-backup-restore-secrets
  namespace: '{{ tpl .namespace $ }}'
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    name: postgres-backup-restore-secrets
    template:
      data:
        PGHOST: "postgres.{{ tpl .namespace $ }}.svc.cluster.local"
        PGPORT: "5432"
        PGUSER: "{{`{{.PGUSER}}`}}"
        PGPASSWORD: "{{`{{ .PGPASSWORD }}`}}"
  data:
  - secretKey: PGPASSWORD
    remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: POSTGRES_USER_PASSWORD }
  - secretKey: PGUSER
    remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: POSTGRES_USERNAME }
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: minio-secrets
  namespace: '{{ tpl .namespace $ }}'
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    name: minio-secrets
  data:
  - secretKey: MINIO_ENDPOINT
    remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: S3_ENDPOINT }
  - secretKey: MINIO_ACCESS_KEY
    remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: S3_ACCESS_KEY }
  - secretKey: MINIO_SECRET_KEY
    remoteRef: { key: '{{ tpl $.Values.global.externalSecretRemoteRefKey $ }}', property: S3_SECRET_KEY }
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-backup-scripts
  namespace: '{{ tpl .namespace $ }}'
data:
  take-dump.sh: |
    #!/bin/bash
    set -e
    set -o pipefail

    rm -f /shared/*  # not sure if shared emptyDir is persisted across job runs

    {{- range $n, $database := .backup.databases }}
    FILENAME="{{ $database }}-$(date +"%FT%TZ").pgdump"
    echo "Creating dump '$FILENAME'"
    pg_dump -Fc {{ $database }} > "/shared/$FILENAME"
    {{- end }}

  upload-to-s3.sh: |
    #!/bin/bash
    set -e
    set -o pipefail

    mc alias set b2 $MINIO_ENDPOINT $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
    {{- $backup := .backup }}
    {{- range $n, $database := .backup.databases }}
    mc cp /shared/{{ $database }}-*.pgdump b2/{{ $backup.bucket }}/{{ tpl $backup.path $ }}/{{ $database }}/
    {{- end }}
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: postgres-backup
  namespace: '{{ tpl .namespace $ }}'
spec:
  # Every 12 hours
  # https://crontab.guru/#0_*/12_*_*_*
  schedule: {{ .backup.schedule }}
  jobTemplate:
    spec:
      template:
        spec:
          volumes:
            - name: shared-data
              emptyDir: {}
            - name: scripts
              configMap:
                name: postgres-backup-scripts
                defaultMode: 0755
          initContainers:
            - name: take-dump
              image: postgres:14.2
              command:
                - /script/take-dump.sh
              envFrom:
              - secretRef:
                  name: postgres-backup-restore-secrets
              volumeMounts:
                - name: shared-data
                  mountPath: /shared
                - name: scripts
                  mountPath: /script/take-dump.sh
                  subPath: take-dump.sh
            - name: upload-to-s3
              image: minio/mc:RELEASE.2022-04-26T18-00-22Z
              command:
                - /script/upload-to-s3.sh
              envFrom:
              - secretRef:
                  name: minio-secrets
              volumeMounts:
                - name: shared-data
                  mountPath: /shared
                - name: scripts
                  mountPath: /script/upload-to-s3.sh
                  subPath: upload-to-s3.sh
          containers:
            - name: job-done
              image: busybox
              command: ["sh", "-c", 'echo "Backup complete" && sleep 1']
          restartPolicy: OnFailure
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-restore-scripts
  namespace: '{{ tpl .namespace $ }}'
data:
  download-from-s3.sh: |
    #!/bin/bash
    set -e
    set -o pipefail

    rm -f /shared/*  # not sure if shared emptyDir is persisted across job runs

    sleep 1  # sometimes, networking takes a bit to warm up

    mc alias set b2 $MINIO_ENDPOINT $MINIO_ACCESS_KEY $MINIO_SECRET_KEY

    # max_retry=5
    # counter=0
    # until mc stat b2/{{ .backup.bucket }}
    # do
    #   [[ counter -eq $max_retry ]] && echo "Failed!" && exit 1
    #   echo "Trying again in 1s. Try #$counter"
    #   sleep 1
    #   ((counter++))
    #   echo "Trying again. Try #$counter"
    # done

    {{- $backup := .backup }}
    {{- range $n, $database := .backup.databases }}
    echo "Getting the latest backup file for database '{{ $database }}'"
    LATEST_FILE=$(mc ls --no-color b2/{{ $backup.bucket }}/{{ tpl $backup.path $ }}/{{ $database }}/ | awk '{ print $6 }' | grep "^{{ $database }}-.*.pgdump$" | sort -r | head -1)
    if [ "$LATEST_FILE" = "" ]; then
      echo "LATEST_FILE is empty"
      exit 1
    fi
    echo "Downloading latest backup file for database '{{ $database }}': '$LATEST_FILE'"
    mc cp b2/{{ $backup.bucket }}/{{ tpl $backup.path $ }}/{{ $database }}/$LATEST_FILE /shared/{{ $database }}.pgdump
    {{- end }}

  restore-dump.sh: |
    #!/bin/bash
    set -e
    set -o pipefail

    {{- range $n, $database := .backup.databases }}
    echo "Restoring dump for database '{{ $database }}'"
    echo "Terminating active connections..."
    createdb {{ $database }} || echo 'Database already exists'
    psql -d {{ $database }} -c 'SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE datname = current_database() AND pid <> pg_backend_pid();'
    echo "Dropping the database..."
    dropdb {{ $database }} || true
    echo "Creating the database..."
    createdb {{ $database }}
    echo "Restoring the database..."
    pg_restore -Fc -d {{ $database }} --no-owner < "/shared/{{ $database }}.pgdump"
    {{- end }}
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: postgres-restore
  namespace: '{{ tpl .namespace $ }}'
spec:
  suspend: true
  schedule: '* * 31 2 *'
  jobTemplate:
    spec:
      # metadata:
      #   name: postgres-restore
      template:
        spec:
          volumes:
            - name: shared-data
              emptyDir: {}
            - name: scripts
              configMap:
                name: postgres-restore-scripts
                defaultMode: 0755
          initContainers:
            - name: download-from-s3
              image: minio/mc:RELEASE.2022-04-26T18-00-22Z
              command:
                - /script/download-from-s3.sh
              envFrom:
              - secretRef:
                  name: minio-secrets
              volumeMounts:
                - name: shared-data
                  mountPath: /shared
                - name: scripts
                  mountPath: /script/download-from-s3.sh
                  subPath: download-from-s3.sh
            - name: restore-dump
              image: postgres:14.2
              command:
                - /script/restore-dump.sh
              envFrom:
              - secretRef:
                  name: postgres-backup-restore-secrets
              volumeMounts:
                - name: shared-data
                  mountPath: /shared
                - name: scripts
                  mountPath: /script/restore-dump.sh
                  subPath: restore-dump.sh
          containers:
            - name: job-done
              image: busybox
              command: ["sh", "-c", 'echo "Restore complete" && sleep 1']
          restartPolicy: OnFailure
{{- end }}
{{- end }}

{{- /*
kubectl create job --from=cronjob/postgres-backup postgres-backup-manual -n database
kubectl delete job postgres-backup-manual -n database

kubectl create job --from=cronjob/postgres-restore postgres-restore-manual -n database
kubectl delete job postgres-restore-manual -n database
*/ -}}
