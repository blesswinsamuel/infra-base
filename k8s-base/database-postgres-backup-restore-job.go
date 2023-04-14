package resourcesbase

import (
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8s-base/imports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8s-base/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

type PostgresBackupRestoreProps struct {
	Enabled       bool      `yaml:"enabled"`
	Schedule      string    `yaml:"schedule"`
	Bucket        string    `yaml:"bucket"`
	Path          string    `yaml:"path"`
	PostgresImage ImageInfo `yaml:"postgresImage"`
	MinioImage    ImageInfo `yaml:"minioImage"`
	Databases     []string  `yaml:"databases"`
}

func NewPostgresBackupRestoreJob(scope constructs.Construct, props PostgresBackupRestoreProps) cdk8s.Chart {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("postgres-backup-restore"), &cprops)

	NewExternalSecret(chart, jsii.String("external-secret-pg"), &ExternalSecretProps{
		Name:            jsii.String("postgres-backup-restore-secrets"),
		RefreshInterval: jsii.String("2m"),
		Template: &externalsecretsio.ExternalSecretV1Beta1SpecTargetTemplate{
			Data: &map[string]*string{
				"PGHOST":     jsii.String("postgres.database.svc.cluster.local"),
				"PGPORT":     jsii.String("5432"),
				"PGUSER":     jsii.String("{{ .PGUSER }}"),
				"PGPASSWORD": jsii.String("{{ .PGPASSWORD }}"),
			},
		},
		Secrets: map[string]string{
			"PGUSER":     "POSTGRES_USERNAME",
			"PGPASSWORD": "POSTGRES_USER_PASSWORD",
		},
	})

	NewExternalSecret(chart, jsii.String("external-secret-minio"), &ExternalSecretProps{
		Name:            jsii.String("minio-secrets"),
		RefreshInterval: jsii.String("2m"),
		Secrets: map[string]string{
			"MINIO_ENDPOINT":   "S3_ENDPOINT",
			"MINIO_ACCESS_KEY": "S3_ACCESS_KEY",
			"MINIO_SECRET_KEY": "S3_SECRET_KEY",
		},
	})

	backupScripts(chart, props)
	restoreScripts(chart, props)

	return chart
}

func backupScripts(chart cdk8s.Chart, props PostgresBackupRestoreProps) {
	k8s.NewKubeConfigMap(chart, jsii.String("postgres-backup-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("postgres-backup-scripts"),
		},
		Data: &map[string]*string{
			"take-dump.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				rm -f /shared/*  # not sure if shared emptyDir is persisted across job runs
			
				{{- range $n, $database := .Databases }}
				FILENAME="{{ $database }}-$(date +"%FT%TZ").pgdump"
				echo "Creating dump '$FILENAME'"
				pg_dump -Fc {{ $database }} > "/shared/$FILENAME"
				{{- end }}
			`)), props),
			"upload-to-s3.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				mc alias set b2 $MINIO_ENDPOINT $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
				{{- $backup := . }}
				{{- range $n, $database := .Databases }}
				mc cp /shared/{{ $database }}-*.pgdump b2/{{ $backup.Bucket }}/{{ $backup.Path }}/{{ $database }}/
				{{- end }}
			`)), props),
		},
	})

	k8s.NewKubeCronJob(chart, jsii.String("postgres-backup-cronjob"), &k8s.KubeCronJobProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("postgres-backup"),
			Annotations: &map[string]*string{
				"secret.reloader.stakater.com/reload": jsii.String("minio-secrets"),
			},
		},
		Spec: &k8s.CronJobSpec{
			Schedule: jsii.String(props.Schedule),
			JobTemplate: &k8s.JobTemplateSpec{
				Spec: &k8s.JobSpec{
					Template: &k8s.PodTemplateSpec{
						Spec: &k8s.PodSpec{
							Volumes: &[]*k8s.Volume{
								{
									Name:     jsii.String("shared-data"),
									EmptyDir: &k8s.EmptyDirVolumeSource{},
								},
								{
									Name: jsii.String("scripts"),
									ConfigMap: &k8s.ConfigMapVolumeSource{
										Name:        jsii.String("postgres-backup-scripts"),
										DefaultMode: jsii.Number(0o755),
									},
								},
							},
							InitContainers: &[]*k8s.Container{
								{
									Name:  jsii.String("take-dump"),
									Image: props.PostgresImage.ToString(),
									Command: &[]*string{
										jsii.String("/script/take-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("postgres-backup-restore-secrets"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-data"), MountPath: jsii.String("/shared")},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/take-dump.sh"), SubPath: jsii.String("take-dump.sh")},
									},
								},
								{
									Name:  jsii.String("upload-to-s3"),
									Image: props.MinioImage.ToString(),
									Command: &[]*string{
										jsii.String("/script/upload-to-s3.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("minio-secrets"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-data"), MountPath: jsii.String("/shared")},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/upload-to-s3.sh"), SubPath: jsii.String("upload-to-s3.sh")},
									},
								},
							},
							Containers: &[]*k8s.Container{
								{
									Name:    jsii.String("job-done"),
									Image:   jsii.String("busybox"),
									Command: jsii.PtrSlice("sh", "-c", "echo 'Backup complete' && sleep 1"),
								},
							},
							RestartPolicy: jsii.String("OnFailure"),
						},
					},
				},
			},
		},
	})
}

func restoreScripts(chart cdk8s.Chart, props PostgresBackupRestoreProps) {
	k8s.NewKubeConfigMap(chart, jsii.String("postgres-restore-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("postgres-restore-scripts"),
		},
		Data: &map[string]*string{
			"download-from-s3.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				rm -f /shared/*  # not sure if shared emptyDir is persisted across job runs
			
				sleep 1  # sometimes, networking takes a bit to warm up
			
				mc alias set b2 $MINIO_ENDPOINT $MINIO_ACCESS_KEY $MINIO_SECRET_KEY
			
				# max_retry=5
				# counter=0
				# until mc stat b2/{{ .Bucket }}
				# do
				#   [[ counter -eq $max_retry ]] && echo "Failed!" && exit 1
				#   echo "Trying again in 1s. Try #$counter"
				#   sleep 1
				#   ((counter++))
				#   echo "Trying again. Try #$counter"
				# done
			
				{{- $backup := . }}
				{{- range $n, $database := .Databases }}
				echo "Getting the latest backup file for database '{{ $database }}'"
				LATEST_FILE=$(mc ls --no-color b2/{{ $backup.Bucket }}/{{ $backup.Path }}/{{ $database }}/ | awk '{ print $6 }' | grep "^{{ $database }}-.*.pgdump$" | sort -r | head -1)
				if [ "$LATEST_FILE" = "" ]; then
				  echo "LATEST_FILE is empty"
				  exit 1
				fi
				echo "Downloading latest backup file for database '{{ $database }}': '$LATEST_FILE'"
				mc cp b2/{{ $backup.Bucket }}/{{ $backup.Path }}/{{ $database }}/$LATEST_FILE /shared/{{ $database }}.pgdump
				{{- end }}
			`)), props),
			"restore-dump.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				{{- range $n, $database := .Databases }}
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
				echo "Database restored."
				{{- end }}
			`)), props),
		},
	})

	k8s.NewKubeCronJob(chart, jsii.String("postgres-restore-cronjob"), &k8s.KubeCronJobProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("postgres-restore"),
			Annotations: &map[string]*string{
				"secret.reloader.stakater.com/reload": jsii.String("minio-secrets"),
			},
		},
		Spec: &k8s.CronJobSpec{
			Suspend:  jsii.Bool(true),
			Schedule: jsii.String("* * 31 2 *"),
			JobTemplate: &k8s.JobTemplateSpec{
				Spec: &k8s.JobSpec{
					Template: &k8s.PodTemplateSpec{
						Spec: &k8s.PodSpec{
							Volumes: &[]*k8s.Volume{
								{
									Name:     jsii.String("shared-data"),
									EmptyDir: &k8s.EmptyDirVolumeSource{},
								},
								{
									Name: jsii.String("scripts"),
									ConfigMap: &k8s.ConfigMapVolumeSource{
										Name:        jsii.String("postgres-restore-scripts"),
										DefaultMode: jsii.Number(0o755),
									},
								},
							},
							InitContainers: &[]*k8s.Container{
								{
									Name:  jsii.String("download-from-s3"),
									Image: props.MinioImage.ToString(),
									Command: &[]*string{
										jsii.String("/script/download-from-s3.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("minio-secrets"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-data"), MountPath: jsii.String("/shared")},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/download-from-s3.sh"), SubPath: jsii.String("download-from-s3.sh")},
									},
								},
								{
									Name:  jsii.String("restore-dump"),
									Image: props.PostgresImage.ToString(),
									Command: &[]*string{
										jsii.String("/script/restore-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("postgres-backup-restore-secrets"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-data"), MountPath: jsii.String("/shared")},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/restore-dump.sh"), SubPath: jsii.String("restore-dump.sh")},
									},
								},
							},
							Containers: &[]*k8s.Container{
								{
									Name:    jsii.String("job-done"),
									Image:   jsii.String("busybox"),
									Command: jsii.PtrSlice("sh", "-c", "echo 'Restore complete' && sleep 1"),
								},
							},
							RestartPolicy: jsii.String("OnFailure"),
						},
					},
				},
			},
		},
	})
}
