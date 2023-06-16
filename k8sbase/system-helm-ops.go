package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/packager"
	corev1 "k8s.io/api/core/v1"
)

type HelmOpsProps struct {
	Enabled   bool             `json:"enabled"`
	ImageInfo k8sapp.ImageInfo `json:"image"`
}

func NewHelmOps(scope packager.Construct, props HelmOpsProps) packager.Construct {
	if !props.Enabled {
		return nil
	}
	return k8sapp.NewApplicationChart(scope, "helm-ops", &k8sapp.ApplicationProps{
		Name: "helm-ops",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "helm-ops",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 51515, Ingress: &k8sapp.ApplicationIngress{Host: "helm-ops." + GetDomain(scope)}},
			},
			Args: []string{
				"server",
				"start",
				"--disable-csrf-token-checks",
				"--insecure",
				"--address=0.0.0.0:51515",
				// "--server-username=" + props.User,
				// "--server-password=helm-ops-secret-password",
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
			EnvFromSecretRef: []string{"helm-ops-password"},
			Env: map[string]string{
				"USER": props.User,
			},
			ExtraVolumeMounts: []corev1.VolumeMount{
				{Name: "helm-ops-config", MountPath: "/app/config/repository.config", SubPath: "repository.config", ReadOnly: true},
			},
		}},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "helm-ops",
				RemoteRefs: map[string]string{
					"git-private-key": "GITHUB_DEPLOY_KEY",
					"known_hosts":     "GITHUB_KNOWN_HOSTS",
				},
			},
		},
	})
}

// gitRepo:
// url: "{{ .gitRepo.url }}"
// branch: "{{ .gitRepo.branch }}"
// helmChartPath: "{{ .helmChartPath }}"
// helmReleaseName: "{{ .helmReleaseName }}"
// helmReleaseNamespace: "{{ .helmReleaseNamespace }}"
// helmExtraValuesFiles: {{ .helmExtraValuesFiles | toJson }}
// clusterName: "test"
// scripts:
// predeploy: |
//   helm repo add blesswinsamuel https://blesswinsamuel.github.io/helm-charts
//   yq e -i '.dependencies[0].repository="https://blesswinsamuel.github.io/helm-charts"' $HELM_CHART_DIR/Chart.yaml
//   yq e -i '.dependencies[0].repository="https://blesswinsamuel.github.io/helm-charts"' $HELM_CHART_DIR/Chart.lock

//   helm dependency update $HELM_CHART_DIR
