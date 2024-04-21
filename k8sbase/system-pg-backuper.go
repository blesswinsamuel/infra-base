package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
)

type PgBackuper struct {
	ImageInfo              k8sapp.ImageInfo `json:"image"`
	Schedule               string           `json:"schedule"`
	Host                   string           `json:"host"`
	LocalBackupVolume      corev1.Volume    `json:"localBackupVolume"`
	Databases              []string         `json:"databases"`
	PersistentVolumeClaims []struct {
		Name         string `json:"name"`
		StorageClass string `json:"storageClass"`
		VolumeName   string `json:"volumeName"`
	} `json:"persistentVolumeClaims"`
}

func (props *PgBackuper) Chart(scope kubegogen.Scope) kubegogen.Scope {
	config := map[string]any{
		"backupDestinations": []map[string]any{
			{"filesystem": map[string]any{"pathTemplate": "{{`/data/{{.Database}}/{{.Database}}.pgdump`}}"}},
			{
				"s3": map[string]any{
					"bucket":       "{{.S3_BUCKET}}",
					"endpoint":     "{{.S3_ENDPOINT}}",
					"accessKey":    "{{.S3_ACCESS_KEY}}",
					"secretKey":    "{{.S3_SECRET_KEY}}",
					"pathTemplate": "{{`backups/postgres/{{.Database}}/{{.Database}}-{{.Timestamp}}.pgdump`}}",
				},
				"encryptionKey": "{{.ENCRYPTION_PASSWORD}}",
			},
		},
		"databases": map[string]any{},
		"notify": []map[string]any{
			{"url": "{{ .HEARTBEAT_URL }}"},
		},
	}
	for _, db := range props.Databases {
		config["databases"].(map[string]any)[db] = map[string]any{
			"host":     props.Host,
			"username": "{{.DB_USERNAME}}",
			"password": "{{.DB_PASSWORD}}",
			"port":     5432,
			"schedule": props.Schedule,
		}
	}
	app := k8sapp.NewApplicationChart(scope, "pg-backuper", &k8sapp.ApplicationProps{
		Name: "pg-backuper",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "pg-backuper",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 2112, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{Path: "/metrics"}},
			},
			Env: map[string]string{
				"CONFIG_FILE": "/config.yaml",
			},
			ExtraVolumeMounts: []corev1.VolumeMount{
				{Name: "data", MountPath: "/data"},
			},
		}},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{{
			Name: "pg-backuper-config",
			RemoteRefs: map[string]string{
				"DB_PASSWORD":         "POSTGRES_USER_PASSWORD",
				"DB_USERNAME":         "POSTGRES_USERNAME",
				"HEARTBEAT_URL":       "DB_BACKUP_HEARTBEAT_URL",
				"S3_SECRET_KEY":       "BACKUP_S3_SECRET_KEY",
				"S3_ACCESS_KEY":       "BACKUP_S3_ACCESS_KEY",
				"S3_BUCKET":           "BACKUP_S3_BUCKET",
				"S3_ENDPOINT":         "BACKUP_S3_ENDPOINT",
				"ENCRYPTION_PASSWORD": "BACKUP_ENCRYPTION_PASSWORD",
			},
			Template: map[string]string{
				"config.yaml": infrahelpers.ToYamlString(config),
			},
			MountName: "config",
			MountPath: "/config.yaml",
			SubPath:   "config.yaml",
			ReadOnly:  true,
		}},
		ExtraVolumes:      []corev1.Volume{{Name: "data", VolumeSource: props.LocalBackupVolume.VolumeSource}},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			// {Name: "pg-backuper-local", StorageClass: "-", VolumeName: "applications-pg-backuper-local", RequestsStorage: "1Gi", MountName: "storage-local", MountPath: "/data"},
			// {Name: "pg-backuper-music", StorageClass: "-", VolumeName: "media-music-applications-pg-backuper", RequestsStorage: "1Gi", MountName: "music", MountPath: "/music"},
		},
	})
	if props.PersistentVolumeClaims != nil {
		for _, pvc := range props.PersistentVolumeClaims {
			if pvc.StorageClass == "__none__" {
				pvc.StorageClass = "-"
			}
			k8sapp.NewPersistentVolumeClaim(app, &k8sapp.PersistentVolumeClaim{
				Name:            pvc.Name,
				RequestsStorage: "1Gi",
				StorageClass:    pvc.StorageClass,
				VolumeName:      pvc.VolumeName,
				AccessModes:     []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
			})
		}
	}
	return app
}
