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
	ImageInfo         k8sapp.ImageInfo          `json:"image"`
	ExtraCollections  []string                  `json:"extraCollections"`
	ExtraParsers      []string                  `json:"extraParsers"`
	ExtraScenarios    []string                  `json:"extraScenarios"`
	ExtraAcquisitions map[string]map[string]any `json:"extraAcquisitions"`
	// HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec
func (props *Crowdsec) Render(scope kubegogen.Scope) {
	// https://github.com/crowdsecurity/example-docker-compose/blob/main/swag/docker-compose.yml
	// https://docs.crowdsec.net/u/getting_started/installation/kubernetes/
	// https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin/blob/main/examples/behind-proxy/docker-compose.cloudflare.yml
	// https://app.crowdsec.net/hub/collections
	collections := []string{"crowdsecurity/traefik"}
	parsers := []string{}
	scenarios := []string{}
	parsers = append(parsers, props.ExtraParsers...)
	collections = append(collections, props.ExtraCollections...)
	scenarios = append(scenarios, props.ExtraScenarios...)

	if props.ExtraAcquisitions == nil {
		props.ExtraAcquisitions = map[string]map[string]any{}
	}
	props.ExtraAcquisitions["traefik"] = map[string]any{
		"filenames": []string{"/var/log/containers/traefik-*_ingress_traefik-*.log"},
		"labels":    map[string]interface{}{"type": "containerd", "program": "traefik"},
	}

	extraAcquisitionsCm := map[string]string{}
	extraAcquisitionsVolMounts := []corev1.VolumeMount{}
	for k, v := range props.ExtraAcquisitions {
		// to fix the error "file is a symlink, but inotify polling is enabled. Crowdsec will not be able to detect rotation. Consider setting poll_without_inotify to true in your configuration"
		v["poll_without_inotify"] = true
		// https://discourse.crowdsec.net/t/error-could-not-create-fsnotify-watcher-too-many-open-files-kubernetes/1584/8
		v["force_inotify"] = true
		extraAcquisitionsCm["acquis-"+k+".yaml"] = infrahelpers.ToYamlString(v)
		extraAcquisitionsVolMounts = append(extraAcquisitionsVolMounts, corev1.VolumeMount{Name: "config", MountPath: "/etc/crowdsec/acquis.d/" + k + ".yaml", SubPath: "acquis-" + k + ".yaml", ReadOnly: true})
	}
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name: "crowdsec",
		InitContainers: []k8sapp.ApplicationContainer{
			{
				Name: "envsubst",
				Image: k8sapp.ImageInfo{
					Repository: "alpine",
					Tag:        "latest",
				},
				ExtraEnvs: []corev1.EnvVar{
					{Name: "TELEGRAM_BOT_TOKEN", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "crowdsec"}, Key: "TELEGRAM_BOT_TOKEN"}}},
				},
				// Command: []string{"sh", "-c", "apk add --no-cache envsubst && envsubst < /config/notifications-telegram.yaml > /config-envsubst/notifications-telegram.yaml"},
				Command: []string{"sh", "-c", `sed "s/\${TELEGRAM_BOT_TOKEN}/$TELEGRAM_BOT_TOKEN/g" /config/notifications-telegram.yaml > /config-envsubst/notifications-telegram.yaml`},
				ExtraVolumeMounts: []corev1.VolumeMount{
					// {Name: "config", MountPath: "/config"},
					{Name: "config-envsubst", MountPath: "/config-envsubst"},
				},
			},
		},
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
					"PARSERS":              strings.Join(parsers, " "),
					"SCEANRIOS":            strings.Join(scenarios, " "),
					// "DISABLE_ONLINE_API":   "true", // If it's a test, we don't want to share signals with CrowdSec so disable the Online API.
					// "PARSERS": "crowdsecurity/cri-logs",
					// "DISABLE_PARSERS": "crowdsecurity/whitelists",
					"BOUNCER_KEY_TRAEFIK": "mysecretkey12345",
				},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false),
					Privileged:               infrahelpers.Ptr(false),
				},
				ExtraVolumeMounts: infrahelpers.MergeLists(extraAcquisitionsVolMounts, []corev1.VolumeMount{
					{Name: "container-logs", MountPath: "/var/log"},
					{Name: "config", MountPath: "/etc/crowdsec/profiles.yaml", SubPath: "profiles.yaml", ReadOnly: true},
					{Name: "config-envsubst", MountPath: "/etc/crowdsec/notifications/http.yaml", SubPath: "notifications-telegram.yaml"},
				}),
				EnvFromSecretRef: []string{"crowdsec"},
			},
		},
		ExtraVolumes: []corev1.Volume{
			{Name: "container-logs", VolumeSource: corev1.VolumeSource{HostPath: &corev1.HostPathVolumeSource{Path: "/var/log"}}},
			// {Name: "config", VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{LocalObjectReference: corev1.LocalObjectReference{Name: "crowdsec-config"}}}},
			{Name: "config-envsubst", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}},
		},
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "crowdsec-config",
				Data: infrahelpers.MergeMaps(extraAcquisitionsCm, map[string]string{
					"profiles.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"name": "default_ip_remediation",
						// "debug": true,
						"filters": []string{
							`Alert.Remediation == true && Alert.GetScope() == "Ip"`,
						},
						"decisions": []map[string]interface{}{
							{"type": "ban", "duration": "4h"},
						},
						"notifications": []string{"telegram"},
						// "duration_expr": `Sprintf('%dh', (GetDecisionsCount(Alert.GetValue()) + 1) * 4)`,
						// # notifications:
						// #   - slack_default  # Set the webhook in /etc/crowdsec/notifications/slack.yaml before enabling this.
						// #   - splunk_default # Set the splunk url and token in /etc/crowdsec/notifications/splunk.yaml before enabling this.
						// #   - http_default   # Set the required http parameters in /etc/crowdsec/notifications/http.yaml before enabling this.
						// #   - email_default  # Set the required email parameters in /etc/crowdsec/notifications/email.yaml before enabling this.
						"on_success": "break",
					}),
					"notifications-telegram.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"type": "http",     // Don't change
						"name": "telegram", // Must match the registered plugin in the profile

						// One of "trace", "debug", "info", "warn", "error", "off"
						"log_level": "info",

						// group_wait:         // Time to wait collecting alerts before relaying a message to this plugin, eg "30s"
						// group_threshold:    // Amount of alerts that triggers a message before <group_wait> has expired, eg "10"
						// max_retry:          // Number of attempts to relay messages to plugins in case of error
						// timeout:            // Time to wait for response from the plugin before considering the attempt a failure, eg "10s"

						//-------------------------
						// plugin-specific options

						// The following template receives a list of models.Alert objects
						// The output goes in the http request body

						// Replace XXXXXXXXX with your Telegram chat ID
						"format": infrahelpers.YAMLRawMessage(`|
  {
    "chat_id": "{{ env "TELEGRAM_CHAT_ID" }}", 
    "text": "
      {{ range . -}}  
      {{$alert := . -}}  
      {{range .Decisions -}}
      {{.Value}} will get {{.Type}} for next {{.Duration}} for triggering {{.Scenario}}.
      {{end -}}
      {{end -}}
    ",
    "reply_markup": {
      "inline_keyboard": [
           {{ $arrLength := len . -}}
           {{ range $i, $value := . -}}
           {{ $V := $value.Source.Value -}}
           [
               {
                   "text": "See {{ $V }} on shodan.io",
                   "url": "https://www.shodan.io/host/{{ $V -}}"
               },
               {
                   "text": "See {{ $V }} on crowdsec.net",
                   "url": "https://app.crowdsec.net/cti/{{ $V -}}"
               }
           ]{{if lt $i ( sub $arrLength 1) }},{{end }}
       {{end -}}
      ]
  }`),

						"url": "https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage", // Replace XXX:YYY with your API key

						"method": "POST",
						"headers": map[string]string{
							"Content-Type": "application/json",
						},
					}),
				}),
				MountToContainers: []string{"envsubst"},
				MountName:         "config",
				MountPath:         "/config",
				ReadOnly:          true,
			},
		},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "crowdsec",
				RemoteRefs: map[string]string{
					"ENROLL_KEY":         "CROWDSEC_ENROLL_KEY",
					"TELEGRAM_CHAT_ID":   "TELEGRAM_CHAT_ID",
					"TELEGRAM_BOT_TOKEN": "TELEGRAM_BOT_TOKEN",
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
