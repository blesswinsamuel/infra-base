package k8sbase

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	corev1 "k8s.io/api/core/v1"
)

type KopiaProps struct {
	Enabled   bool             `yaml:"enabled"`
	ImageInfo k8sapp.ImageInfo `yaml:"image"`
	Hostname  string           `yaml:"hostname"`
	User      string           `yaml:"user"`
}

// https://kopia.io/docs/installation/#docker-images
func NewKopia(scope constructs.Construct, props KopiaProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	return k8sapp.NewApplicationChart(scope, "kopia", &k8sapp.ApplicationProps{
		Name:     "kopia",
		Hostname: props.Hostname,
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "kopia",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 51515, Ingress: &k8sapp.ApplicationIngress{Host: "kopia." + GetDomain(scope)}},
			},
			Args: []string{
				"server",
				"start",
				"--disable-csrf-token-checks",
				"--insecure",
				"--address=0.0.0.0:51515",
				// "--server-username=" + props.User,
				// "--server-password=kopia-secret-password",
				"--without-password",
				"--log-level=debug",
				"--file-log-level=error",
				"--json-log-console",
				// "--override-username=" + props.User,
				// "--refresh-interval=60s",
				"--no-check-for-updates",
				"--no-grpc",
				"--no-legacy-api",
			},
			EnvFromSecretRef: []string{"kopia-password"},
			Env: map[string]string{
				"USER": props.User,
			},
			ExtraVolumeMounts: []corev1.VolumeMount{
				{Name: "kopia-config", MountPath: "/app/config/repository.config", SubPath: "repository.config", ReadOnly: true},
			},
		}},
		ExtraVolumes: []corev1.Volume{
			{Name: "kopia-config", VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: "kopia-config"}}},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "kopia-password",
				RemoteRefs: map[string]string{
					"KOPIA_PASSWORD": "BACKUP_ENCRYPTION_PASSWORD",
				},
			},
			{
				Name: "kopia-config",
				Template: map[string]string{
					"repository.config": infrahelpers.ToJSONString(map[string]any{
						"storage": map[string]any{
							"type": "s3",
							"config": map[string]any{
								"bucket":          "{{ .S3_BUCKET }}",
								"endpoint":        "{{ .S3_ENDPOINT }}",
								"accessKeyID":     "{{ .S3_ACCESS_KEY }}",
								"secretAccessKey": "{{ .S3_SECRET_KEY }}",
								"sessionToken":    "",
							},
						},
						"caching": map[string]any{
							"cacheDirectory":       "/app/cache",
							"maxCacheSize":         5242880000,
							"maxMetadataCacheSize": 5242880000,
							"maxListCacheDuration": 30,
						},
						// "hostname":                "bless-mac-wired",
						"username":                props.User,
						"description":             "Repository in S3: {{ .S3_ENDPOINT }} {{ .S3_BUCKET }}",
						"enableActions":           false,
						"formatBlobCacheDuration": 900000000000,
					}),
				},
				RemoteRefs: map[string]string{
					"S3_ACCESS_KEY": "BACKUP_S3_ACCESS_KEY",
					"S3_SECRET_KEY": "BACKUP_S3_SECRET_KEY",
					"S3_ENDPOINT":   "BACKUP_S3_ENDPOINT",
					"S3_BUCKET":     "BACKUP_S3_BUCKET",
				},
			},
		},
	})
}

// # Mount local folders needed by kopia
// - /path/to/config/dir:/app/config
// - /path/to/cache/dir:/app/cache
// - /path/to/logs/dir:/app/logs
// # Mount local folders to snapshot
// - /path/to/data/dir:/data:ro
// # Mount repository location
// - /path/to/repository/dir:/repository
// # Mount path for browsing mounted snaphots
// - /path/to/tmp/dir:/tmp:shared
