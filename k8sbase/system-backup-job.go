package k8sbase

import (
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/externalsecretsio"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

type BackupJobProps struct {
	Enabled bool `yaml:"enabled"`
	Kopia   struct {
		Image ImageInfo `yaml:"image"`
	} `yaml:"kopia"`
	Postgres struct {
		Enabled           bool        `yaml:"enabled"`
		Image             ImageInfo   `yaml:"image"`
		Schedule          string      `yaml:"schedule"`
		LocalBackupVolume *k8s.Volume `yaml:"localBackupVolume"`
		Databases         []string    `yaml:"databases"`
	} `yaml:"postgres"`
	Filesystem struct {
		Enabled bool      `yaml:"enabled"`
		Image   ImageInfo `yaml:"image"`
		Paths   []struct {
			Schedule string `yaml:"schedule"`
			Source   string `yaml:"source"`
			Policy   struct {
			} `yaml:"policy"`
		} `yaml:"paths"`
	} `yaml:"filesystem"`
}

func NewBackupJob(scope constructs.Construct, props BackupJobProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	chart := cdk8s.NewChart(scope, jsii.String("backup-job"), &cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	})
	NewBackupJobPostgres(chart, props)
	return chart
}

func NewBackupJobPostgres(chart constructs.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	NewExternalSecret(chart, jsii.String("external-secret-pg"), &ExternalSecretProps{
		Name: jsii.String("backup-job-postgres"),
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
	NewExternalSecret(chart, jsii.String("external-secret-s3"), &ExternalSecretProps{
		Name:            jsii.String("backup-job-s3"),
		RefreshInterval: jsii.String("2m"),
		Secrets: map[string]string{
			"S3_ACCESS_KEY":  "BACKUP_S3_ACCESS_KEY",
			"S3_SECRET_KEY":  "BACKUP_S3_SECRET_KEY",
			"S3_ENDPOINT":    "BACKUP_S3_ENDPOINT",
			"S3_BUCKET":      "BACKUP_S3_BUCKET",
			"KOPIA_PASSWORD": "BACKUP_ENCRYPTION_PASSWORD",
		},
	})
	NewExternalSecret(chart, jsii.String("external-secret-heartbeat"), &ExternalSecretProps{
		Name: jsii.String("backup-job-heartbeat"),
		Secrets: map[string]string{
			"HEARTBEAT_URL": "DB_BACKUP_HEARTBEAT_URL",
		},
	})

	sharedMountPath := jsii.String("/pgdumps")
	k8s.NewKubeConfigMap(chart, jsii.String("backup-job-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("backup-job-scripts"),
		},
		Data: &map[string]*string{
			"take-postgres-dump.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				rm -f /pgdumps/*
			
				{{- range $n, $database := . }}
				FILENAME="{{ $database }}.pgdump"
				echo "Creating dump '$FILENAME'"
				pg_dump -Fc {{ $database }} > "/pgdumps/$FILENAME"
				{{- end }}
			`)), props.Postgres.Databases),
			"kopia-snapshot.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia
				kopia snapshot create /pgdumps
			`)), props.Kopia),
		},
	})

	props.Postgres.LocalBackupVolume.Name = jsii.String("shared-backup-data")
	k8s.NewKubeCronJob(chart, jsii.String("backup-job-postgres"), &k8s.KubeCronJobProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("backup-job-postgres"),
			Annotations: &map[string]*string{
				"secret.reloader.stakater.com/reload": jsii.String("backup-job-postgres,postgres-backups-heartbeat"),
			},
		},
		Spec: &k8s.CronJobSpec{
			Schedule: jsii.String(props.Postgres.Schedule),
			JobTemplate: &k8s.JobTemplateSpec{
				Spec: &k8s.JobSpec{
					Template: &k8s.PodTemplateSpec{
						Spec: &k8s.PodSpec{
							Hostname: jsii.String("backup-job-postgres"),
							Volumes: &[]*k8s.Volume{
								props.Postgres.LocalBackupVolume,
								{
									Name: jsii.String("scripts"),
									ConfigMap: &k8s.ConfigMapVolumeSource{
										Name:        jsii.String("backup-job-scripts"),
										DefaultMode: jsii.Number(0o755),
									},
								},
							},
							InitContainers: &[]*k8s.Container{
								{
									Name:  jsii.String("take-dump"),
									Image: props.Postgres.Image.ToString(),
									Command: &[]*string{
										jsii.String("/script/take-postgres-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("backup-job-postgres"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/take-postgres-dump.sh"), SubPath: jsii.String("take-postgres-dump.sh")},
									},
								},
								{
									Name:  jsii.String("kopia-snapshot"),
									Image: props.Kopia.Image.ToString(),
									Command: &[]*string{
										jsii.String("/script/kopia-snapshot.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-job-s3")}},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/kopia-snapshot.sh"), SubPath: jsii.String("kopia-snapshot.sh")},
									},
								},
							},
							Containers: &[]*k8s.Container{
								{
									Name:  jsii.String("job-done"),
									Image: jsii.String("curlimages/curl"),
									Command: jsii.PtrSlice(
										"sh",
										"-c",
										`curl -s -i --connect-timeout 5 --max-time 10 --retry 5 --retry-delay 0 --retry-max-time 40 "$HEARTBEAT_URL" && echo 'Backup complete' && sleep 1`,
									),
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-job-heartbeat")}},
									},
								},
							},
							RestartPolicy: jsii.String("OnFailure"),
						},
					},
				},
			},
		},
	})
	NewRestoreJobPostgres(chart, props)
}

func NewRestoreJobPostgres(chart constructs.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	sharedMountPath := jsii.String("/pgdumps")
	k8s.NewKubeConfigMap(chart, jsii.String("backup-job-restore-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("backup-job-restore-scripts"),
		},
		Data: &map[string]*string{
			"kopia-restore.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				# rm -f /pgdumps/*

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia
				kopia snapshot restore /pgdumps --snapshot-time latest
				ls -lah /pgdumps
			`)), props.Kopia),
			"restore-postgres-dump.sh": GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				{{- range $n, $database := . }}
				echo "Restoring dump for database '{{ $database }}'"
				echo "Terminating active connections..."
				createdb {{ $database }} || echo 'Database already exists'
				psql -d {{ $database }} -c 'SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE datname = current_database() AND pid <> pg_backend_pid();'
				echo "Dropping the database..."
				dropdb {{ $database }} || true
				echo "Creating the database..."
				createdb {{ $database }}
				echo "Restoring the database..."
				pg_restore -Fc -d {{ $database }} --no-owner < "/pgdumps/{{ $database }}.pgdump"
				echo "Database restored."
				{{- end }}
			`)), props.Postgres.Databases),
		},
	})

	props.Postgres.LocalBackupVolume.Name = jsii.String("shared-backup-data")
	k8s.NewKubeCronJob(chart, jsii.String("restore-job-postgres"), &k8s.KubeCronJobProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("restore-job-postgres"),
			Annotations: &map[string]*string{
				"secret.reloader.stakater.com/reload": jsii.String("backup-job-postgres"),
			},
		},
		Spec: &k8s.CronJobSpec{
			Suspend:  jsii.Bool(true),
			Schedule: jsii.String("* * 31 2 *"),
			JobTemplate: &k8s.JobTemplateSpec{
				Spec: &k8s.JobSpec{
					Template: &k8s.PodTemplateSpec{
						Spec: &k8s.PodSpec{
							Hostname: jsii.String("backup-job-postgres"),
							Volumes: &[]*k8s.Volume{
								props.Postgres.LocalBackupVolume,
								{
									Name: jsii.String("scripts"),
									ConfigMap: &k8s.ConfigMapVolumeSource{
										Name:        jsii.String("backup-job-restore-scripts"),
										DefaultMode: jsii.Number(0o755),
									},
								},
							},
							InitContainers: &[]*k8s.Container{
								{
									Name:  jsii.String("kopia-restore"),
									Image: props.Kopia.Image.ToString(),
									Command: &[]*string{
										jsii.String("/script/kopia-restore.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-job-s3")}},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/kopia-restore.sh"), SubPath: jsii.String("kopia-restore.sh")},
									},
									TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
								},
								{
									Name:  jsii.String("restore-dump"),
									Image: props.Postgres.Image.ToString(),
									Command: &[]*string{
										jsii.String("/script/restore-postgres-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("backup-job-postgres"),
											},
										},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/restore-postgres-dump.sh"), SubPath: jsii.String("restore-postgres-dump.sh")},
									},
									TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
								},
							},
							Containers: &[]*k8s.Container{
								{
									Name:                     jsii.String("job-done"),
									Image:                    jsii.String("busybox"),
									Command:                  jsii.PtrSlice("sh", "-c", "echo 'Restore complete' && sleep 1"),
									TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
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
