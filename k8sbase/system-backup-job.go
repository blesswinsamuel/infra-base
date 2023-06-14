package k8sbase

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	"github.com/muesli/reflow/dedent"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackupJobProps struct {
	Enabled bool `json:"enabled"`
	Kopia   struct {
		Image k8sapp.ImageInfo `json:"image"`
	} `json:"kopia"`
	Postgres struct {
		Enabled           bool             `json:"enabled"`
		Image             k8sapp.ImageInfo `json:"image"`
		Schedule          string           `json:"schedule"`
		Host              string           `json:"host"`
		LocalBackupVolume corev1.Volume    `json:"localBackupVolume"`
		Databases         []string         `json:"databases"`
	} `json:"postgres"`
	Filesystem struct {
		Enabled bool `json:"enabled"`
		Jobs    []struct {
			Name         string        `json:"name"`
			Schedule     string        `json:"schedule"`
			SourceVolume corev1.Volume `json:"sourceVolume"`
			Paths        []string      `json:"paths"`
			Policy       struct {
			} `json:"policy"`
		} `json:"jobs"`
	} `json:"filesystem"`
}

func NewBackupJob(scope packager.Construct, props BackupJobProps) packager.Construct {
	if !props.Enabled {
		return nil
	}
	chart := scope.Chart("backup-job", packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	})
	k8sapp.NewExternalSecret(chart, "external-secret-pg", &k8sapp.ExternalSecretProps{
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
	k8sapp.NewExternalSecret(chart, "external-secret-s3", &k8sapp.ExternalSecretProps{
		Name: "backup-restore-job-s3",
		RemoteRefs: map[string]string{
			"S3_ACCESS_KEY":  "BACKUP_S3_ACCESS_KEY",
			"S3_SECRET_KEY":  "BACKUP_S3_SECRET_KEY",
			"S3_ENDPOINT":    "BACKUP_S3_ENDPOINT",
			"S3_BUCKET":      "BACKUP_S3_BUCKET",
			"KOPIA_PASSWORD": "BACKUP_ENCRYPTION_PASSWORD",
		},
	})

	k8sapp.NewConfigMap(chart, "backup-job-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "backup-job-scripts",
		Data: map[string]string{
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

func NewBackupJobPostgres(chart packager.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	k8sapp.NewExternalSecret(chart, "external-secret-heartbeat", &k8sapp.ExternalSecretProps{
		Name: "backup-job-heartbeat",
		RemoteRefs: map[string]string{
			"HEARTBEAT_URL": "DB_BACKUP_HEARTBEAT_URL",
		},
	})

	sharedMountPath := "/pgdumps"

	props.Postgres.LocalBackupVolume.Name = "shared-backup-data"
	k8sapp.NewK8sObject(chart, "backup-job-postgres", &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: "backup-job-postgres",
			Annotations: map[string]string{
				"secret.reloader.stakater.com/reload": "backup-restore-job-postgres,postgres-backups-heartbeat",
			},
		},
		Spec: batchv1.CronJobSpec{
			Schedule: props.Postgres.Schedule,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Hostname: "backup-job-postgres",
							Volumes: []corev1.Volume{
								props.Postgres.LocalBackupVolume,
								{
									Name:         "scripts",
									VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-job-scripts"}, DefaultMode: infrahelpers.Ptr(int32(0o755))}},
								},
							},
							InitContainers: []corev1.Container{
								{
									Name:  "take-dump",
									Image: props.Postgres.Image.String(),
									Command: []string{
										"/script/take-postgres-dump.sh",
									},
									EnvFrom: []corev1.EnvFromSource{
										{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-restore-job-postgres"}}},
									},
									Env: []corev1.EnvVar{
										{Name: "FOLDER", Value: sharedMountPath},
									},
									VolumeMounts: []corev1.VolumeMount{
										{Name: "shared-backup-data", MountPath: sharedMountPath},
										{Name: "scripts", MountPath: "/script/take-postgres-dump.sh", SubPath: "take-postgres-dump.sh"},
									},
								},
								{
									Name:  "kopia-snapshot",
									Image: props.Kopia.Image.String(),
									Command: []string{
										"/script/kopia-postgres-snapshot.sh",
									},
									EnvFrom: []corev1.EnvFromSource{
										{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-restore-job-s3"}}},
									},
									Env: []corev1.EnvVar{
										{Name: "FOLDER", Value: sharedMountPath},
									},
									VolumeMounts: []corev1.VolumeMount{
										{Name: "shared-backup-data", MountPath: sharedMountPath},
										{Name: "scripts", MountPath: "/script/kopia-postgres-snapshot.sh", SubPath: "kopia-postgres-snapshot.sh"},
									},
								},
							},
							Containers: []corev1.Container{
								{
									Name:  "job-done",
									Image: "curlimages/curl:8.1.0",
									Command: []string{
										"sh",
										"-c",
										`curl -s -i --connect-timeout 5 --max-time 10 --retry 5 --retry-delay 0 --retry-max-time 40 "$HEARTBEAT_URL" && echo 'Backup complete' && sleep 1`,
									},
									EnvFrom: []corev1.EnvFromSource{
										{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-job-heartbeat"}}},
									},
								},
							},
							RestartPolicy: "OnFailure",
						},
					},
				},
			},
		},
	})
}

func NewRestoreJobPostgres(chart packager.Construct, props BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	sharedMountPath := "/pgdumps"
	k8sapp.NewConfigMap(chart, "restore-job-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "restore-job-scripts",
		Data: map[string]string{
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

	props.Postgres.LocalBackupVolume.Name = "shared-backup-data"
	k8sapp.NewK8sObject(chart, "restore-job-postgres", &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: "restore-job-postgres",
			Annotations: map[string]string{
				"secret.reloader.stakater.com/reload": "backup-restore-job-postgres",
			},
		},
		Spec: batchv1.CronJobSpec{
			Suspend:  infrahelpers.Ptr(true),
			Schedule: "* * 31 2 *",
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Hostname: "backup-job-postgres",
							Volumes: []corev1.Volume{
								props.Postgres.LocalBackupVolume,
								{
									Name: "scripts",
									VolumeSource: corev1.VolumeSource{
										ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "restore-job-scripts"}, DefaultMode: infrahelpers.Ptr(int32(0o755))},
									},
								},
							},
							InitContainers: []corev1.Container{
								{
									Name:  "kopia-restore",
									Image: props.Kopia.Image.String(),
									Command: []string{
										"/script/kopia-restore.sh",
									},
									EnvFrom: []corev1.EnvFromSource{
										{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-restore-job-s3"}}},
									},
									VolumeMounts: []corev1.VolumeMount{
										{Name: "shared-backup-data", MountPath: sharedMountPath},
										{Name: "scripts", MountPath: "/script/kopia-restore.sh", SubPath: "kopia-restore.sh"},
									},
									TerminationMessagePolicy: "FallbackToLogsOnError",
								},
								{
									Name:  "restore-dump",
									Image: props.Postgres.Image.String(),
									Command: []string{
										"/script/restore-postgres-dump.sh",
									},
									EnvFrom: []corev1.EnvFromSource{
										{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-restore-job-postgres"}}},
									},
									VolumeMounts: []corev1.VolumeMount{
										{Name: "shared-backup-data", MountPath: sharedMountPath},
										{Name: "scripts", MountPath: "/script/restore-postgres-dump.sh", SubPath: "restore-postgres-dump.sh"},
									},
									TerminationMessagePolicy: "FallbackToLogsOnError",
								},
							},
							Containers:    []corev1.Container{echoContainer("Restore complete")},
							RestartPolicy: "OnFailure",
						},
					},
				},
			},
		},
	})
}

func echoContainer(msg string) corev1.Container {
	return corev1.Container{
		Name:                     "job-done",
		Image:                    "busybox",
		Command:                  []string{"sh", "-c", fmt.Sprintf("echo %q && sleep 1", msg)},
		TerminationMessagePolicy: "FallbackToLogsOnError",
	}
}

func NewBackupJobFilesystem(chart packager.Construct, props BackupJobProps) {
	if !props.Filesystem.Enabled {
		return
	}

	for _, job := range props.Filesystem.Jobs {
		sharedMountPath := "/filesystem"

		job.SourceVolume.Name = "source-data"
		folders := []string{}
		for _, path := range job.Paths {
			folders = append(folders, sharedMountPath+"/"+strings.TrimPrefix(path, "/"))
		}
		k8sapp.NewK8sObject(chart, "backup-job-filesystem-"+job.Name, &batchv1.CronJob{
			ObjectMeta: metav1.ObjectMeta{
				Name: "backup-job-filesystem-" + job.Name,
			},
			Spec: batchv1.CronJobSpec{
				Schedule: job.Schedule,
				JobTemplate: batchv1.JobTemplateSpec{
					Spec: batchv1.JobSpec{
						Template: corev1.PodTemplateSpec{
							Spec: corev1.PodSpec{
								Hostname: "backup-job-filesystem-" + job.Name,
								Volumes: []corev1.Volume{
									job.SourceVolume,
									{
										Name: "scripts",
										VolumeSource: corev1.VolumeSource{
											ConfigMap: &corev1.ConfigMapVolumeSource{
												LocalObjectReference: corev1.LocalObjectReference{
													Name: "backup-job-scripts",
												},
												DefaultMode: infrahelpers.Ptr(int32(0o755)),
											},
										},
									},
								},
								InitContainers: []corev1.Container{
									{
										Name:  "kopia-snapshot",
										Image: props.Kopia.Image.String(),
										Command: []string{
											"/script/kopia-filesystem-snapshot.sh",
										},
										EnvFrom: []corev1.EnvFromSource{
											{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: "backup-restore-job-s3"}}},
										},
										Env: []corev1.EnvVar{
											{Name: "FOLDERS", Value: strings.Join(folders, " ")},
										},
										VolumeMounts: []corev1.VolumeMount{
											{Name: "source-data", MountPath: sharedMountPath},
											{Name: "scripts", MountPath: "/script/kopia-filesystem-snapshot.sh", SubPath: "kopia-filesystem-snapshot.sh"},
										},
									},
								},
								Containers:    []corev1.Container{echoContainer("Backup complete")},
								RestartPolicy: "OnFailure",
							},
						},
					},
				},
			},
		})
	}
}
