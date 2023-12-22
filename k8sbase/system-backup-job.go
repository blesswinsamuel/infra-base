package k8sbase

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/muesli/reflow/dedent"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackupJobProps struct {
	Kopia struct {
		Image k8sapp.ImageInfo `json:"image"`
	} `json:"kopia"`
	Postgres struct {
		Enabled           bool             `json:"enabled"`
		Image             k8sapp.ImageInfo `json:"image"`
		ImagePullSecrets  []string         `json:"imagePullSecrets"`
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

type cronJobProps struct {
	DisabledByDefault bool
	Name              string
	ImagePullSecrets  []string
	Schedule          string
	LocalBackupVolume corev1.Volume
	SharedFolder      string
	ScriptsConfigMap  string
	Hearbeat          cronJobHeartbeat
	Commands          []cronJobScript
}

type cronJobHeartbeat struct {
	Enabled         bool
	HeartBeatSecret string
}

type cronJobScript struct {
	Name          string
	Image         k8sapp.ImageInfo
	Command       []string
	EnvFromSecret string
	MountScript   string
	Env           map[string]string
}

func echoContainer(msg string) corev1.Container {
	return corev1.Container{
		Name:                     "job-done",
		Image:                    "busybox",
		Command:                  []string{"sh", "-c", fmt.Sprintf("echo %q && sleep 1", msg)},
		TerminationMessagePolicy: "FallbackToLogsOnError",
	}
}

func newCronJob(chart kubegogen.Construct, id string, props cronJobProps) kubegogen.ApiObject {
	initContainers := []corev1.Container{}
	props.LocalBackupVolume.Name = "shared-backup-data"
	for _, command := range props.Commands {
		container := corev1.Container{
			Name:    command.Name,
			Image:   command.Image.String(),
			Command: command.Command,
			VolumeMounts: []corev1.VolumeMount{
				{Name: "shared-backup-data", MountPath: props.SharedFolder},
				{Name: "scripts", MountPath: "/script/" + command.MountScript, SubPath: command.MountScript},
			},
			TerminationMessagePolicy: "FallbackToLogsOnError",
		}
		for k, v := range command.Env {
			container.Env = append(container.Env, corev1.EnvVar{Name: k, Value: v})
		}
		if command.EnvFromSecret != "" {
			container.EnvFrom = []corev1.EnvFromSource{
				{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: command.EnvFromSecret}}},
			}
		}
		initContainers = append(initContainers, container)
	}
	containers := []corev1.Container{}
	if props.Hearbeat.Enabled {
		containers = append(containers, corev1.Container{
			Name:  "heartbeat",
			Image: "curlimages/curl:8.1.0",
			Command: []string{
				"sh",
				"-c",
				`curl -s -i --connect-timeout 5 --max-time 10 --retry 5 --retry-delay 0 --retry-max-time 40 "$HEARTBEAT_URL" && echo 'Cron Job complete' && sleep 1`,
			},
			EnvFrom: []corev1.EnvFromSource{
				{SecretRef: &corev1.SecretEnvSource{LocalObjectReference: corev1.LocalObjectReference{Name: props.Hearbeat.HeartBeatSecret}}},
			},
		})
	} else {
		containers = append(containers, echoContainer("Cron Job complete"))
	}
	var imagePullSecrets []corev1.LocalObjectReference
	for _, secret := range props.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, corev1.LocalObjectReference{Name: secret})
	}
	jobSpec := batchv1.CronJobSpec{
		Schedule:          props.Schedule,
		ConcurrencyPolicy: batchv1.ReplaceConcurrent,
		JobTemplate: batchv1.JobTemplateSpec{
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						ImagePullSecrets: imagePullSecrets,
						Volumes: []corev1.Volume{
							props.LocalBackupVolume,
							{
								Name:         "scripts",
								VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: props.ScriptsConfigMap}, DefaultMode: infrahelpers.Ptr(int32(0o755))}},
							},
						},
						InitContainers: initContainers,
						Containers:     containers,
						RestartPolicy:  "OnFailure",
					},
				},
			},
		},
	}
	if props.DisabledByDefault {
		jobSpec.Suspend = infrahelpers.Ptr(true)
		jobSpec.Schedule = "* * 31 2 *"
		// jobSpec.JobTemplate.Spec.Template.Spec.RestartPolicy = "Never"

	}
	return k8sapp.NewK8sObject(chart, id, &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name: props.Name,
		},
		Spec: jobSpec,
	})
}

func (props *BackupJobProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	chart := scope.Chart("backup-job", kubegogen.ChartProps{
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

	NewBackupPostgresJob(chart, props)
	NewRestorePostgresJob(chart, props)
	NewBackupFilesystemJob(chart, props)
	NewRestoreFilesystemJob(chart, props)
	return chart
}

func NewBackupPostgresJob(chart kubegogen.Construct, props *BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	k8sapp.NewConfigMap(chart, "backup-job-postgres-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "backup-job-postgres-scripts",
		Data: map[string]string{
			"take-postgres-dump.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				DATABASE=$1
				FOLDER=$2
			
				mkdir -p $FOLDER
				rm -f $FOLDER/*
			
				FILENAME="$DATABASE.pgdump"

				echo "Creating dump '$FILENAME'"
				pg_dump -Fc $DATABASE > "$FOLDER/$FILENAME"
			`)), props.Postgres.Databases),
			"kopia-postgres-snapshot.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				FOLDER=$1

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia --override-hostname kopia
				kopia snapshot create $FOLDER
			`)), props.Kopia),
		},
	})

	k8sapp.NewExternalSecret(chart, "external-secret-heartbeat", &k8sapp.ExternalSecretProps{
		Name: "backup-job-heartbeat",
		RemoteRefs: map[string]string{
			"HEARTBEAT_URL": "DB_BACKUP_HEARTBEAT_URL",
		},
	})

	for _, databaseName := range props.Postgres.Databases {
		sharedMountPath := "/pgdumps/" + databaseName
		newCronJob(chart, "backup-job-postgres-"+databaseName, cronJobProps{
			Name:              "backup-job-postgres-" + databaseName,
			Schedule:          props.Postgres.Schedule,
			ScriptsConfigMap:  "backup-job-postgres-scripts",
			LocalBackupVolume: props.Postgres.LocalBackupVolume,
			SharedFolder:      "/pgdumps",
			ImagePullSecrets:  props.Postgres.ImagePullSecrets,
			Commands: []cronJobScript{
				{Name: "take-dump", Image: props.Postgres.Image, Command: []string{"/script/take-postgres-dump.sh", databaseName, sharedMountPath}, EnvFromSecret: "backup-restore-job-postgres", MountScript: "take-postgres-dump.sh"},
				{Name: "kopia-snapshot", Image: props.Kopia.Image, Command: []string{"/script/kopia-postgres-snapshot.sh", sharedMountPath}, EnvFromSecret: "backup-restore-job-s3", MountScript: "kopia-postgres-snapshot.sh"},
			},
			Hearbeat: cronJobHeartbeat{Enabled: true, HeartBeatSecret: "backup-job-heartbeat"},
		})
	}
}

func NewRestorePostgresJob(chart kubegogen.Construct, props *BackupJobProps) {
	if !props.Postgres.Enabled {
		return
	}
	k8sapp.NewConfigMap(chart, "restore-job-postgres-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "restore-job-postgres-scripts",
		Data: map[string]string{
			"kopia-restore.sh": strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				FOLDER=$1

				mkdir -p $FOLDER
				# rm -f $FOLDER/*

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia --override-hostname kopia
				kopia snapshot restore $FOLDER --snapshot-time latest
				ls -lah $FOLDER
			`)),
			"restore-postgres-dump.sh": strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail
			
				DATABASE=$1
				FOLDER=$2
				echo "Restoring dump for database '$DATABASE'"
				echo "Terminating active connections..."
				createdb $DATABASE || echo 'Database already exists'
				psql -d $DATABASE -c 'SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE datname = current_database() AND pid <> pg_backend_pid();'
				echo "Dropping the database..."
				dropdb $DATABASE || true
				echo "Creating the database..."
				createdb $DATABASE
				echo "Restoring the database..."
				pg_restore -Fc -d $DATABASE --no-owner < "$FOLDER/$DATABASE.pgdump"
				echo "Database restored."
			`)),
		},
	})

	for _, databaseName := range props.Postgres.Databases {
		sharedMountPath := "/pgdumps/" + databaseName
		newCronJob(chart, "restore-job-postgres-"+databaseName, cronJobProps{
			DisabledByDefault: true,
			Name:              "restore-job-postgres-" + databaseName,
			LocalBackupVolume: props.Postgres.LocalBackupVolume,
			ImagePullSecrets:  props.Postgres.ImagePullSecrets,
			ScriptsConfigMap:  "restore-job-postgres-scripts",
			SharedFolder:      "/pgdumps",
			Commands: []cronJobScript{
				{Name: "kopia-restore", Image: props.Kopia.Image, Command: []string{"/script/kopia-restore.sh", sharedMountPath}, EnvFromSecret: "backup-restore-job-s3", MountScript: "kopia-restore.sh"},
				{Name: "restore-dump", Image: props.Postgres.Image, Command: []string{"/script/restore-postgres-dump.sh", databaseName, sharedMountPath}, EnvFromSecret: "backup-restore-job-postgres", MountScript: "restore-postgres-dump.sh"},
			},
		})
	}
}

func NewBackupFilesystemJob(chart kubegogen.Construct, props *BackupJobProps) {
	if !props.Filesystem.Enabled {
		return
	}
	k8sapp.NewConfigMap(chart, "backup-job-filesystem-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "backup-job-filesystem-scripts",
		Data: map[string]string{
			"kopia-filesystem-snapshot.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia --override-hostname kopia
				for FOLDER in ${FOLDERS}
				do
				    kopia snapshot create $FOLDER
				done
			`)), props.Kopia),
		},
	})

	for _, job := range props.Filesystem.Jobs {
		sharedMountPath := "/filesystem"

		folders := []string{}
		for _, path := range job.Paths {
			folders = append(folders, sharedMountPath+"/"+strings.TrimPrefix(path, "/"))
		}
		newCronJob(chart, "backup-job-filesystem-"+job.Name, cronJobProps{
			Name:              "backup-job-filesystem-" + job.Name,
			Schedule:          job.Schedule,
			LocalBackupVolume: job.SourceVolume,
			ScriptsConfigMap:  "backup-job-filesystem-scripts",
			SharedFolder:      sharedMountPath,
			Commands: []cronJobScript{
				{
					Name:          "kopia-snapshot",
					Image:         props.Kopia.Image,
					Command:       []string{"/script/kopia-filesystem-snapshot.sh"},
					EnvFromSecret: "backup-restore-job-s3",
					MountScript:   "kopia-filesystem-snapshot.sh",
					Env:           map[string]string{"FOLDERS": strings.Join(folders, " ")},
				},
			},
		})
	}
}

func NewRestoreFilesystemJob(chart kubegogen.Construct, props *BackupJobProps) {
	if !props.Filesystem.Enabled {
		return
	}
	k8sapp.NewConfigMap(chart, "restore-job-filesystem-scripts-cm", &k8sapp.ConfigmapProps{
		Name: "restore-job-filesystem-scripts",
		Data: map[string]string{
			"kopia-filesystem-restore.sh": infrahelpers.GoTemplate(strings.TrimSpace(dedent.String(`
				#!/bin/bash
				set -e
				set -o pipefail

				kopia repository connect s3 --bucket=$S3_BUCKET --access-key=$S3_ACCESS_KEY --secret-access-key=$S3_SECRET_KEY --endpoint=$S3_ENDPOINT --override-username kopia --override-hostname kopia
				for FOLDER in ${FOLDERS}
				do
				    kopia snapshot restore $FOLDER --snapshot-time latest
				done
			`)), props.Kopia),
		},
	})

	for _, job := range props.Filesystem.Jobs {
		sharedMountPath := "/filesystem"

		folders := []string{}
		for _, path := range job.Paths {
			folders = append(folders, sharedMountPath+"/"+strings.TrimPrefix(path, "/"))
		}
		newCronJob(chart, "restore-job-filesystem-"+job.Name, cronJobProps{
			DisabledByDefault: true,
			Name:              "restore-job-filesystem-" + job.Name,
			LocalBackupVolume: job.SourceVolume,
			ScriptsConfigMap:  "restore-job-filesystem-scripts",
			SharedFolder:      sharedMountPath,
			Commands: []cronJobScript{
				{
					Name:          "kopia-snapshot",
					Image:         props.Kopia.Image,
					Command:       []string{"/script/kopia-filesystem-restore.sh"},
					EnvFromSecret: "backup-restore-job-s3",
					MountScript:   "kopia-filesystem-restore.sh",
					Env:           map[string]string{"FOLDERS": strings.Join(folders, " ")},
				},
			},
		})
	}
}
