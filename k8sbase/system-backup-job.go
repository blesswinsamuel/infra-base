package k8sbase

import (
	"fmt"
	"strings"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/muesli/reflow/dedent"
)

type BackupJobProps struct {
	Enabled bool `yaml:"enabled"`
	Kopia   struct {
		Image k8sapp.ImageInfo `yaml:"image"`
	} `yaml:"kopia"`
	Postgres struct {
		Enabled           bool             `yaml:"enabled"`
		Image             k8sapp.ImageInfo `yaml:"image"`
		Schedule          string           `yaml:"schedule"`
		Host              string           `yaml:"host"`
		LocalBackupVolume *k8s.Volume      `yaml:"localBackupVolume"`
		Databases         []string         `yaml:"databases"`
	} `yaml:"postgres"`
	Filesystem struct {
		Enabled bool `yaml:"enabled"`
		Jobs    []struct {
			Name         string      `yaml:"name"`
			Schedule     string      `yaml:"schedule"`
			SourceVolume *k8s.Volume `yaml:"sourceVolume"`
			Paths        []string    `yaml:"paths"`
			Policy       struct {
			} `yaml:"policy"`
		} `yaml:"jobs"`
	} `yaml:"filesystem"`
}

func NewBackupJob(scope constructs.Construct, props BackupJobProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	chart := cdk8s.NewChart(scope, jsii.String("backup-job"), &cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	})
	k8sapp.NewExternalSecret(chart, jsii.String("external-secret-pg"), &k8sapp.ExternalSecretProps{
		Name: "backup-restore-job-postgres",
		Template: map[string]string{
			"PGHOST":     props.Postgres.Host,
			"PGPORT":     "5432",
			"PGUSER":     "{{ .PGUSER }}",
			"PGPASSWORD": "{{ .PGPASSWORD }}",
		},
		RemoteRefs: map[string]string{
			"PGUSER":     "POSTGRES_USERNAME",
			"PGPASSWORD": "POSTGRES_USER_PASSWORD",
		},
	})
	k8sapp.NewExternalSecret(chart, jsii.String("external-secret-s3"), &k8sapp.ExternalSecretProps{
		Name: "backup-restore-job-s3",
		RemoteRefs: map[string]string{
			"S3_ACCESS_KEY":  "BACKUP_S3_ACCESS_KEY",
			"S3_SECRET_KEY":  "BACKUP_S3_SECRET_KEY",
			"S3_ENDPOINT":    "BACKUP_S3_ENDPOINT",
			"S3_BUCKET":      "BACKUP_S3_BUCKET",
			"KOPIA_PASSWORD": "BACKUP_ENCRYPTION_PASSWORD",
		},
	})

	k8s.NewKubeConfigMap(chart, jsii.String("backup-job-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("backup-job-scripts"),
		},
		Data: &map[string]*string{
			"take-postgres-dump.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				rm -f $FOLDER/*
			
				{{- range $n, $database := . }}
				FILENAME="{{ $database }}.pgdump"
				echo "Creating dump '$FILENAME'"
				pg_dump -Fc {{ $database }} > "$FOLDER/$FILENAME"
				{{- end }}
			`)), props.Postgres.Databases),
			"kopia-postgres-snapshot.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia
				kopia snapshot create $FOLDER
			`)), props.Kopia),
			"kopia-filesystem-snapshot.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia
				for FOLDER in ${FOLDERS}
				do
					kopia snapshot create $FOLDER
				done
			`)), props.Kopia),
		},
	})

	NewBackupJobPostgres(chart, props)
	NewRestoreJobPostgres(chart, props)
	NewBackupJobFilesystem(chart, props)
	return chart
}

func NewBackupJobPostgres(chart constructs.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	NewExternalSecret(chart, jsii.String("external-secret-heartbeat"), &ExternalSecretProps{
		Name: jsii.String("backup-job-heartbeat"),
		Secrets: map[string]string{
			"HEARTBEAT_URL": "DB_BACKUP_HEARTBEAT_URL",
		},
	})

	sharedMountPath := jsii.String("/pgdumps")

	props.Postgres.LocalBackupVolume.Name = jsii.String("shared-backup-data")
	k8s.NewKubeCronJob(chart, jsii.String("backup-job-postgres"), &k8s.KubeCronJobProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("backup-job-postgres"),
			Annotations: &map[string]*string{
				"secret.reloader.stakater.com/reload": jsii.String("backup-restore-job-postgres,postgres-backups-heartbeat"),
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
									Image: jsii.String(props.Postgres.Image.String()),
									Command: &[]*string{
										jsii.String("/script/take-postgres-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-restore-job-postgres")}},
									},
									Env: &[]*k8s.EnvVar{
										{Name: jsii.String("FOLDER"), Value: sharedMountPath},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/take-postgres-dump.sh"), SubPath: jsii.String("take-postgres-dump.sh")},
									},
								},
								{
									Name:  jsii.String("kopia-snapshot"),
									Image: jsii.String(props.Kopia.Image.String()),
									Command: &[]*string{
										jsii.String("/script/kopia-postgres-snapshot.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-restore-job-s3")}},
									},
									Env: &[]*k8s.EnvVar{
										{Name: jsii.String("FOLDER"), Value: sharedMountPath},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/kopia-postgres-snapshot.sh"), SubPath: jsii.String("kopia-postgres-snapshot.sh")},
									},
								},
							},
							Containers: &[]*k8s.Container{
								{
									Name:  jsii.String("job-done"),
									Image: jsii.String("curlimages/curl:8.1.0"),
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
}

func NewRestoreJobPostgres(chart constructs.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	sharedMountPath := jsii.String("/pgdumps")
	k8s.NewKubeConfigMap(chart, jsii.String("restore-job-scripts-cm"), &k8s.KubeConfigMapProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("restore-job-scripts"),
		},
		Data: &map[string]*string{
			"kopia-restore.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				# rm -f /pgdumps/*

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia
				kopia snapshot restore /pgdumps --snapshot-time latest
				ls -lah /pgdumps
			`)), props.Kopia),
			"restore-postgres-dump.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
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
				"secret.reloader.stakater.com/reload": jsii.String("backup-restore-job-postgres"),
			},
		},
		Spec: &k8s.CronJobSpec{
			Suspend:  jsii.Bool(true),
			Schedule: jsii.String("* * 31 2 *"),
			JobTemplate: &k8s.JobTemplateSpec{
				Spec: &k8s.JobSpec{
					Template: &k8s.PodTemplateSpec{
						Spec: &k8s.PodSpec{
							Hostname: jsii.String("backup-job-postgres"), // should be same as backup job hostname
							Volumes: &[]*k8s.Volume{
								props.Postgres.LocalBackupVolume,
								{
									Name: jsii.String("scripts"),
									ConfigMap: &k8s.ConfigMapVolumeSource{
										Name:        jsii.String("restore-job-scripts"),
										DefaultMode: jsii.Number(0o755),
									},
								},
							},
							InitContainers: &[]*k8s.Container{
								{
									Name:  jsii.String("kopia-restore"),
									Image: jsii.String(props.Kopia.Image.String()),
									Command: &[]*string{
										jsii.String("/script/kopia-restore.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-restore-job-s3")}},
									},
									VolumeMounts: &[]*k8s.VolumeMount{
										{Name: jsii.String("shared-backup-data"), MountPath: sharedMountPath},
										{Name: jsii.String("scripts"), MountPath: jsii.String("/script/kopia-restore.sh"), SubPath: jsii.String("kopia-restore.sh")},
									},
									TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
								},
								{
									Name:  jsii.String("restore-dump"),
									Image: jsii.String(props.Postgres.Image.String()),
									Command: &[]*string{
										jsii.String("/script/restore-postgres-dump.sh"),
									},
									EnvFrom: &[]*k8s.EnvFromSource{
										{
											SecretRef: &k8s.SecretEnvSource{
												Name: jsii.String("backup-restore-job-postgres"),
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
							Containers:    &[]*k8s.Container{echoContainer("Restore complete")},
							RestartPolicy: jsii.String("OnFailure"),
						},
					},
				},
			},
		},
	})
}

func echoContainer(msg string) *k8s.Container {
	return &k8s.Container{
		Name:                     jsii.String("job-done"),
		Image:                    jsii.String("busybox"),
		Command:                  jsii.PtrSlice("sh", "-c", fmt.Sprintf("echo %q && sleep 1", msg)),
		TerminationMessagePolicy: jsii.String("FallbackToLogsOnError"),
	}
}

func NewBackupJobFilesystem(chart constructs.Construct, props BackupJobProps) {
	if !props.Filesystem.Enabled {
		return
	}

	for _, job := range props.Filesystem.Jobs {
		sharedMountPath := jsii.String("/filesystem")

		job.SourceVolume.Name = jsii.String("source-data")
		folders := []string{}
		for _, path := range job.Paths {
			folders = append(folders, *sharedMountPath+"/"+strings.TrimPrefix(path, "/"))
		}
		k8s.NewKubeCronJob(chart, jsii.String("backup-job-filesystem-"+job.Name), &k8s.KubeCronJobProps{
			Metadata: &k8s.ObjectMeta{
				Name: jsii.String("backup-job-filesystem-" + job.Name),
			},
			Spec: &k8s.CronJobSpec{
				Schedule: jsii.String(job.Schedule),
				JobTemplate: &k8s.JobTemplateSpec{
					Spec: &k8s.JobSpec{
						Template: &k8s.PodTemplateSpec{
							Spec: &k8s.PodSpec{
								Hostname: jsii.String("backup-job-filesystem-" + job.Name),
								Volumes: &[]*k8s.Volume{
									job.SourceVolume,
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
										Name:  jsii.String("kopia-snapshot"),
										Image: jsii.String(props.Kopia.Image.String()),
										Command: &[]*string{
											jsii.String("/script/kopia-filesystem-snapshot.sh"),
										},
										EnvFrom: &[]*k8s.EnvFromSource{
											{SecretRef: &k8s.SecretEnvSource{Name: jsii.String("backup-restore-job-s3")}},
										},
										Env: &[]*k8s.EnvVar{
											{Name: jsii.String("FOLDERS"), Value: jsii.String(strings.Join(folders, " "))},
										},
										VolumeMounts: &[]*k8s.VolumeMount{
											{Name: jsii.String("source-data"), MountPath: sharedMountPath},
											{Name: jsii.String("scripts"), MountPath: jsii.String("/script/kopia-filesystem-snapshot.sh"), SubPath: jsii.String("kopia-filesystem-snapshot.sh")},
										},
									},
								},
								Containers:    &[]*k8s.Container{echoContainer("Backup complete")},
								RestartPolicy: jsii.String("OnFailure"),
							},
						},
					},
				},
			},
		})
	}
}
