package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Grafana struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Sidecar   struct {
		ImageInfo k8sapp.ImageInfo `json:"image"`
	} `json:"sidecar"`
	AnonymousAuthEnabled bool `json:"anonymousAuthEnabled"`
	AuthProxyEnabled     bool `json:"authProxyEnabled"`
	DisableSanitizeHTML  bool `json:"disableSanitizeHTML"`
	Ingress              struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
}

// https://github.com/grafana/helm-charts/tree/main/charts/grafana
func (props *Grafana) Chart(scope kubegogen.Construct) kubegogen.Construct {
	app := k8sapp.NewApplicationChart(scope, "grafana", &k8sapp.ApplicationProps{
		Name: "grafana",
		// AutomountServiceAccountToken: true,
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "grafana",
				Image: props.ImageInfo,
				Env: infrahelpers.MergeMaps(
					map[string]string{
						// "GF_SECURITY_DISABLE_INITIAL_ADMIN_CREATION": "true",
						// "GF_AUTH_BASIC_ENABLED":          "true",

						"GF_PATHS_DATA":         "/var/lib/grafana/",
						"GF_PATHS_LOGS":         "/var/log/grafana",
						"GF_PATHS_PLUGINS":      "/var/lib/grafana/plugins",
						"GF_PATHS_PROVISIONING": "/etc/grafana/provisioning",
						"GF_SERVER_DOMAIN":      props.Ingress.SubDomain + "." + GetDomain(scope),

						"GF_SERVER_ENABLE_GZIP":                 "true",
						"GF_PANELS_DISABLE_SANITIZE_HTML":       infrahelpers.Ternary(props.DisableSanitizeHTML, "true", "false"),
						"GF_ANALYTICS_REPORTING_ENABLED":        "false",
						"GF_ANALYTICS_CHECK_FOR_UPDATES":        "false",
						"GF_ANALYTICS_CHECK_FOR_PLUGIN_UPDATES": "false",
					},
					infrahelpers.Ternary(props.AnonymousAuthEnabled, map[string]string{
						"GF_AUTH_ANONYMOUS_HIDE_VERSION": "true",
						"GF_AUTH_ANONYMOUS_ENABLED":      "true",
						"GF_AUTH_ANONYMOUS_ORG_NAME":     "Main Org.",
						"GF_AUTH_ANONYMOUS_ORG_ROLE":     "Admin",
						"GF_AUTH_DISABLE_LOGIN_FORM":     "true",
					}, nil),
					infrahelpers.Ternary(props.AuthProxyEnabled, map[string]string{
						"GF_AUTH_PROXY_ENABLED":            "true",
						"GF_AUTH_PROXY_HEADER_NAME":        "Remote-User",
						"GF_AUTH_PROXY_HEADER_PROPERTY":    "username",
						"GF_AUTH_PROXY_AUTO_SIGN_UP":       "true",
						"GF_AUTH_PROXY_HEADERS":            "Groups:Remote-Group",
						"GF_AUTH_PROXY_ENABLE_LOGIN_TOKEN": "false",

						"GF_USERS_ALLOW_SIGN_UP":        "false",
						"GF_USERS_AUTO_ASSIGN_ORG":      "true",
						"GF_USERS_AUTO_ASSIGN_ORG_ROLE": "Admin",
					}, nil),
				),
				ExtraEnvs: []corev1.EnvVar{
					{Name: "GF_SECURITY_ADMIN_USER", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "username", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
					{Name: "GF_SECURITY_ADMIN_PASSWORD", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "password", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false), Capabilities: &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}}, SeccompProfile: &corev1.SeccompProfile{Type: corev1.SeccompProfileTypeRuntimeDefault},
				},
				Ports: []k8sapp.ContainerPort{{Name: "http", Port: 3000, ServicePort: 80, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}}},
				LivenessProbe: &corev1.Probe{
					InitialDelaySeconds: int32(60),
					ProbeHandler:        corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/api/health"}},
					TimeoutSeconds:      int32(30),
				},
				ReadinessProbe: &corev1.Probe{
					ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Port: intstr.FromString("http"), Path: "/api/health"}},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "storage", MountPath: "/var/lib/grafana"},
					{Name: "sc-dashboard-volume", MountPath: "/tmp/dashboards"},
					{Name: "sc-datasources-volume", MountPath: "/etc/grafana/provisioning/datasources"},
				},
			},
			// https://github.com/kiwigrid/k8s-sidecar#configuration-environment-variables
			{
				Name:  "grafana-sc-dashboard",
				Image: props.Sidecar.ImageInfo,
				Env: map[string]string{
					"METHOD":            "WATCH",
					"LABEL":             "grafana_dashboard",
					"LABEL_VALUE":       "1",
					"FOLDER":            "/tmp/dashboards",
					"FOLDER_ANNOTATION": "grafana_folder",
					"NAMESPACE":         "ALL",
					"RESOURCE":          "configmap",
					"REQ_URL":           "http://localhost:3000/api/admin/provisioning/dashboards/reload",
					"REQ_METHOD":        "POST",
					"LOG_LEVEL":         "DEBUG",
				},
				ExtraEnvs: []corev1.EnvVar{
					{Name: "REQ_USERNAME", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "username", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
					{Name: "REQ_PASSWORD", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "password", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false), Capabilities: &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}}, SeccompProfile: &corev1.SeccompProfile{Type: corev1.SeccompProfileTypeRuntimeDefault},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "sc-dashboard-volume", MountPath: "/tmp/dashboards"},
				},
			},
			{
				Name:  "grafana-sc-datasource",
				Image: props.Sidecar.ImageInfo,
				Env: map[string]string{
					"METHOD":      "WATCH",
					"LABEL":       "grafana_datasource",
					"LABEL_VALUE": "1",
					"FOLDER":      "/etc/grafana/provisioning/datasources",
					"NAMESPACE":   "ALL",
					"RESOURCE":    "configmap",
					"REQ_URL":     "http://localhost:3000/api/admin/provisioning/datasources/reload",
					"REQ_METHOD":  "POST",
					"LOG_LEVEL":   "DEBUG",
				},
				ExtraEnvs: []corev1.EnvVar{
					{Name: "REQ_USERNAME", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "username", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
					{Name: "REQ_PASSWORD", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "password", LocalObjectReference: corev1.LocalObjectReference{Name: "grafana-admin-credentials"}}}},
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false), Capabilities: &corev1.Capabilities{Drop: []corev1.Capability{"ALL"}}, SeccompProfile: &corev1.SeccompProfile{Type: corev1.SeccompProfileTypeRuntimeDefault},
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "sc-datasources-volume", MountPath: "/etc/grafana/provisioning/datasources"},
				},
			},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "sc-datasources-volume", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
			{Name: "sc-dashboard-volume", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
			{Name: "storage", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
		ServiceAccountName:   "grafana",
		CreateServiceAccount: true,
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "grafana-admin-credentials",
				Template: map[string]string{
					"username": "admin",
					"password": "{{ .GRAFANA_ADMIN_PASSWORD }}",
				},
				RemoteRefs: map[string]string{
					"GRAFANA_ADMIN_PASSWORD": "GRAFANA_ADMIN_PASSWORD",
				},
			},
		},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "grafana-config-dashboards",
				Data: map[string]string{
					"provider.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"apiVersion": 1,
						"providers": []map[string]interface{}{
							{
								"name":                  "sidecarProvider",
								"orgId":                 1,
								"type":                  "file",
								"disableDeletion":       false,
								"allowUiUpdates":        false,
								"updateIntervalSeconds": 30,
								"options": map[string]interface{}{
									"path":                      "/tmp/dashboards",
									"foldersFromFilesStructure": true,
								},
							},
						},
					}),
				},
				MountToContainers: []string{"grafana"},
				MountName:         "sc-dashboard-provider",
				MountPath:         "/etc/grafana/provisioning/dashboards/sc-dashboardproviders.yaml",
				SubPath:           "provider.yaml",
			},
			{
				Name: "grafana",
				Data: map[string]string{
					"grafana.ini": `
[log]
mode = console
`,
				},
				MountToContainers: []string{"grafana"},
				MountName:         "config",
				MountPath:         "/etc/grafana/grafana.ini",
				SubPath:           "grafana.ini",
				ReadOnly:          true,
			},
		},
	})
	k8sapp.NewK8sObject(app, "clusterrole", &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{Name: "grafana"},
		Rules: []rbacv1.PolicyRule{
			{APIGroups: []string{""}, Resources: []string{"configmaps", "secrets"}, Verbs: []string{"get", "list", "watch"}},
		},
	})
	k8sapp.NewK8sObject(app, "clusterrolebinding", &rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{Name: "grafana"},
		RoleRef:    rbacv1.RoleRef{Kind: "ClusterRole", Name: "grafana", APIGroup: "rbac.authorization.k8s.io"},
		Subjects:   []rbacv1.Subject{{Kind: "ServiceAccount", Name: "grafana", Namespace: app.Namespace()}},
	})
	return app
}
