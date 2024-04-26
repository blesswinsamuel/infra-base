package k8sbase

import (
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apiextensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Crowdsec struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	// HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec
func (props *Crowdsec) Render(scope kubegogen.Scope) {
	// k8sapp.NewHelm(scope, &k8sapp.HelmProps{
	// 	ChartInfo:   props.HelmChartInfo,
	// 	ReleaseName: "crowdsec",
	// 	Namespace:   scope.Namespace(),
	// 	Values: map[string]any{
	// 		"container_runtime": "containerd",
	// 		"lapi": map[string]any{
	// 			"service": map[string]any{
	// 				"annotations": map[string]any{
	// 					"prometheus.io/scrape": "true",
	// 					"prometheus.io/port":   "6060",
	// 				},
	// 			},
	// 			"metrics": map[string]any{
	// 				"enabled": true,
	// 			},
	// 			"env": []map[string]any{
	// 				// {"name": "ENROLL_KEY", "valueFrom": map[string]any{"secretKeyRef": map[string]any{"name": "crowdsec-keys", "key": "ENROLL_KEY"}}},
	// 				// {"name": "ENROLL_INSTANCE_NAME", "value": "homelab"},
	// 				// {"name": "DISABLE_ONLINE_API", "value": "false"},
	// 			},
	// 			"ingress": map[string]any{
	// 				"enabled":     false,
	// 				"annotations": GetCertIssuerAnnotation(scope),
	// 				"host":        "crowdsec-lapi" + "." + GetDomain(scope),
	// 				"tls": []map[string]any{
	// 					{
	// 						"hosts":      []string{"crowdsec-lapi" + "." + GetDomain(scope)},
	// 						"secretName": "crowdsec-lapi-tls",
	// 					},
	// 				},
	// 			},
	// 			"dashboard": map[string]any{
	// 				"enabled": false,
	// 				"ingress": map[string]any{
	// 					"enabled":     true,
	// 					"annotations": GetCertIssuerAnnotation(scope),
	// 					"host":        "crowdsec-lapi-dashboard" + "." + GetDomain(scope),
	// 					"tls": []map[string]any{
	// 						{
	// 							"hosts":      []string{"crowdsec-lapi-dashboard" + "." + GetDomain(scope)},
	// 							"secretName": "crowdsec-lapi-dashboard-tls",
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		"agent": map[string]any{
	// 			"service": map[string]any{
	// 				"annotations": map[string]any{
	// 					"prometheus.io/scrape": "true",
	// 					"prometheus.io/port":   "6060",
	// 				},
	// 			},
	// 			"metrics": map[string]any{
	// 				"enabled": true,
	// 			},
	// 			"acquisition": []map[string]any{
	// 				{
	// 					"namespace": "ingress",
	// 					"podName":   "traefik-*",
	// 					"program":   "traefik",
	// 				},
	// 			},
	// 			// "additionalAcquisition": []map[string]any{
	// 			// 	{
	// 			// 		"filenames":     []string{"/var/log/containers/traefik-*_ingress_*.log"},
	// 			// 		"force_inotify": true,
	// 			// 		"labels": map[string]any{
	// 			// 			"type":    "containerd",
	// 			// 			"program": "traefik",
	// 			// 		},
	// 			// 	},
	// 			// },
	// 			"env": []map[string]any{
	// 				{"name": "COLLECTIONS", "value": "crowdsecurity/traefik"},
	// 				{"name": "PARSERS", "value": "crowdsecurity/cri-logs"},
	// 				{"name": "DISABLE_PARSERS", "value": "crowdsecurity/whitelists"},
	// 				// {"name": "DISABLE_ONLINE_API", "value": "false"},
	// 			},
	// 		},
	// 		"secrets": map[string]any{
	// 			"username": "crowdsec",
	// 			"password": "crowdsec@123",
	// 		},
	// 	},
	// })
	// for _, obj := range *helmResources.ApiObjects() {
	// 	if *obj.Metadata().Name() == "crowdsec-agent" && *obj.Kind() == "DaemonSet" {
	// 		obj.AddJsonPatch(
	// 			packager.JsonPatch_Test(("/spec/template/spec/containers/0/env/1/name"), "DISABLE_ONLINE_API"),
	// 			packager.JsonPatch_Replace(("/spec/template/spec/containers/0/env/1"), map[string]any{
	// 				"name":  "DISABLE_ONLINE_API",
	// 				"value": "false",
	// 			}),
	// 		)
	// 	}
	// }

	// https://github.com/crowdsecurity/example-docker-compose/blob/main/swag/docker-compose.yml
	// https://docs.crowdsec.net/u/getting_started/installation/kubernetes/
	// https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin/blob/main/examples/behind-proxy/docker-compose.cloudflare.yml
	collections := []string{"crowdsecurity/linux", "crowdsecurity/traefik", "crowdsecurity/http-cve", "crowdsecurity/whitelist-good-actors", "crowdsecurity/sshd"}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name: "crowdsec",
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:  "crowdsec",
				Image: props.ImageInfo,
				Ports: []k8sapp.ContainerPort{
					{Name: "http", Port: 6060, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
					{Name: "lapi", Port: 8080},
				},
				Env: map[string]string{
					"GID":                  "1000",
					"ENROLL_INSTANCE_NAME": "homelab",
					"ENROLL_TAGS":          "k8s traefik homelab",
					"COLLECTIONS":          strings.Join(collections, " "),
					// "DISABLE_ONLINE_API":   "true", // If it's a test, we don't want to share signals with CrowdSec so disable the Online API.
					// "PARSERS": "crowdsecurity/cri-logs",
					// "DISABLE_PARSERS": "crowdsecurity/whitelists",
					"BOUNCER_KEY_TRAEFIK": "mysecretkey12345",
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false),
					Privileged:               infrahelpers.Ptr(false),
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					{Name: "container-logs", MountPath: "/var/log"},
				},
				EnvFromSecretRef: []string{"crowdsec"},
			},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "container-logs", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/log"}}},
		},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "crowdsec-acquis-config",
				Data: map[string]string{
					"acquis.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"filenames":            []string{"/var/log/containers/traefik-*_ingress_*.log"},
						"force_inotify":        true,
						"poll_without_inotify": true,
						"labels": map[string]interface{}{
							// "type": "traefik",
							"type":    "containerd",
							"program": "traefik",
						},
					}),
				},
				MountName: "acquis-config",
				MountPath: "/etc/crowdsec/acquis.yaml",
				SubPath:   "acquis.yaml",
				ReadOnly:  true,
			},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "crowdsec",
				RemoteRefs: map[string]string{
					"ENROLL_KEY": "CROWDSEC_ENROLL_KEY",
				},
			},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: "crowdsec-db", RequestsStorage: "1Gi", MountPath: "/var/lib/crowdsec/data", MountName: "crowdsec-db"},
			{Name: "crowdsec-config", RequestsStorage: "100Mi", MountPath: "/etc/crowdsec", MountName: "crowdsec-config"},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "crowdsec-traefik-bouncer"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Plugin: map[string]apiextensionv1.JSON{
				"crowdsec-bouncer": {
					Raw: []byte(infrahelpers.ToJSONString(map[string]interface{}{
						"Enabled": true,
						// "LogLevel":           "DEBUG",
						"CrowdsecMode":       "stream",
						"CrowdsecLapiScheme": "http",
						"CrowdsecLapiHost":   "crowdsec." + scope.Namespace() + ".svc.cluster.local:8080",
						"CrowdsecLapiKey":    "mysecretkey12345",
						// "clienttrustedips": "10.0.10.30/32",
					})),
				},
			},
		},
	})

	// // https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec-traefik-bouncer
	// k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
	// 	Name: "crowdsec-traefik-bouncer",
	// 	Containers: []k8sapp.ApplicationContainer{
	// 		{
	// 			Name:  "crowdsec-traefik-bouncer",
	// 			Image: props.ImageInfo,
	// 			Ports: []k8sapp.ContainerPort{
	// 				{Name: "http", Port: 8080, ServicePort: 80, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{}},
	// 				// {Name: "lapi", Port: 8080},
	// 			},
	// 			Env: map[string]string{
	// 				"CROWDSEC_BOUNCER_API_KEY": "",
	// 				"CROWDSEC_AGENT_HOST":      "",
	// 				"GIN_MODE":                 "debug",
	// 			},
	// 			LivenessProbe:  &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/v1/ping", Port: intstr.FromString("http")}}},
	// 			ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/v1/ping", Port: intstr.FromString("http")}}},
	// 			SecurityContext: &corev1.SecurityContext{
	// 				AllowPrivilegeEscalation: infrahelpers.Ptr(false),
	// 				Privileged:               infrahelpers.Ptr(false),
	// 			},
	// 			ExtraVolumeMounts: []corev1.VolumeMount{
	// 				{Name: "container-logs", MountPath: "/var/log"},
	// 			},
	// 			EnvFromSecretRef: []string{"crowdsec"},
	// 		},
	// 	},
	// })
	// scope.AddApiObject(&traefikv1alpha1.Middleware{
	// 	ObjectMeta: metav1.ObjectMeta{Name: "crowdsec-traefik-bouncer"},
	// 	Spec: traefikv1alpha1.MiddlewareSpec{
	// 		ForwardAuth: &traefikv1alpha1.ForwardAuth{
	// 			Address: fmt.Sprintf("http://crowdsec-traefik-bouncer.%s.svc.cluster.local/api/v1/forwardAuth", scope.Namespace()),
	// 		},
	// 	},
	// })
}

// kubectl exec -it -n ingress deploy/crowdsec -- cscli decisions list
// kubectl exec -it -n ingress deploy/crowdsec -- cscli metrics
