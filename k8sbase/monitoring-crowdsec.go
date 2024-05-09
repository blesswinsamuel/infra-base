package k8sbase

import (
	"sort"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apiextensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Crowdsec struct {
	ImageInfo                  k8sapp.ImageInfo          `json:"image"`
	ExtraCollections           []string                  `json:"extraCollections"`
	ExtraParsers               []string                  `json:"extraParsers"`
	ExtraScenarios             []string                  `json:"extraScenarios"`
	ExtraAcquisitions          map[string]map[string]any `json:"extraAcquisitions"`
	InstanceName               string                    `json:"instanceName"`
	NodePortService            bool                      `json:"nodePortService"`
	EnableSshRemediation       bool                      `json:"enableFirewallRemediation"`
	BouncerKeys                map[string]string         `json:"bouncerKeys"`
	ForwardedHeadersTrustedIPs []string                  `json:"forwardedHeadersTrustedIPs"`
	Notifiers                  []string                  `json:"notifiers"`
	// HelmChartInfo k8sapp.ChartInfo `json:"helm"`
}

// https://github.com/crowdsecurity/helm-charts/tree/main/charts/crowdsec
func (props *Crowdsec) Render(scope kubegogen.Scope) {
	// https://github.com/crowdsecurity/example-docker-compose/blob/main/swag/docker-compose.yml
	// https://docs.crowdsec.net/u/getting_started/installation/kubernetes/
	// https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin/blob/main/examples/behind-proxy/docker-compose.cloudflare.yml
	// https://app.crowdsec.net/hub/collections
	collections := []string{
		"crowdsecurity/linux",
		"crowdsecurity/sshd",

		"crowdsecurity/traefik",
		// base-http-scenarios and http-cve are included in traefik, but added explicitly here for pruning to work correctly
		"crowdsecurity/base-http-scenarios",
		"crowdsecurity/http-cve",
	}
	parsers := []string{}
	scenarios := []string{}
	parsers = append(parsers, props.ExtraParsers...)
	collections = append(collections, props.ExtraCollections...)
	sort.Strings(collections)
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
	for _, k := range infrahelpers.MapKeys(props.ExtraAcquisitions) {
		v := props.ExtraAcquisitions[k]
		if v["filenames"] != nil {
			// to fix the error "file is a symlink, but inotify polling is enabled. Crowdsec will not be able to detect rotation. Consider setting poll_without_inotify to true in your configuration"
			v["poll_without_inotify"] = true
			// https://discourse.crowdsec.net/t/error-could-not-create-fsnotify-watcher-too-many-open-files-kubernetes/1584/8
			v["force_inotify"] = true
		}
		extraAcquisitionsCm["acquis-"+k+".yaml"] = infrahelpers.ToYamlString(v)
		extraAcquisitionsVolMounts = append(extraAcquisitionsVolMounts, corev1.VolumeMount{Name: "config", MountPath: "/etc/crowdsec/acquis.d/" + k + ".yaml", SubPath: "acquis-" + k + ".yaml", ReadOnly: true})
	}
	extraAcquisitionsCm["acquis.yaml"] = ``
	extraAcquisitionsVolMounts = append(extraAcquisitionsVolMounts, corev1.VolumeMount{Name: "config", MountPath: "/etc/crowdsec/acquis.yaml", SubPath: "acquis.yaml", ReadOnly: true})

	notifications := []string{}
	for _, notifier := range props.Notifiers {
		switch notifier {
		case "telegram":
			notifications = append(notifications, "telegram_default")
		case "slack":
			notifications = append(notifications, "slack_default")
		}
	}

	firewallBouncerProfile := infrahelpers.ToYamlString(map[string]interface{}{
		"name": "firewall_ip_remediation",
		// "debug": true,
		"filters": []string{
			`Alert.Remediation == true && Alert.GetScope() == "Ip" && Alert.GetScenario() contains "ssh"`,
		},
		"decisions": []map[string]interface{}{
			{"type": "ban-firewall", "duration": "24h"},
		},
		"notifications": notifications,
		"on_success":    "break",
	})
	traefikBouncerProfile := infrahelpers.ToYamlString(map[string]interface{}{
		"name": "traefik_ip_remediation",
		// "debug": true,
		"filters": []string{
			`Alert.Remediation == true && Alert.GetScope() == "Ip"`,
		},
		"decisions": []map[string]interface{}{
			{"type": "ban", "duration": "24h"},
		},
		"notifications": notifications,
		// "duration_expr": `Sprintf('%dh', (GetDecisionsCount(Alert.GetValue()) + 1) * 4)`,
		// # notifications:
		// #   - slack_default  # Set the webhook in /etc/crowdsec/notifications/slack.yaml before enabling this.
		// #   - splunk_default # Set the splunk url and token in /etc/crowdsec/notifications/splunk.yaml before enabling this.
		// #   - http_default   # Set the required http parameters in /etc/crowdsec/notifications/http.yaml before enabling this.
		// #   - email_default  # Set the required email parameters in /etc/crowdsec/notifications/email.yaml before enabling this.
		"on_success": "break",
	})
	profiles := []string{}
	if props.EnableSshRemediation {
		profiles = append(profiles, firewallBouncerProfile)
	}
	profiles = append(profiles, traefikBouncerProfile)

	bouncerKeys := map[string]string{}
	for key, value := range props.BouncerKeys {
		bouncerKeys["BOUNCER_KEY_"+key] = value
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
					{Name: "SLACK_API_URL", ValueFrom: &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{LocalObjectReference: corev1.LocalObjectReference{Name: "crowdsec"}, Key: "SLACK_API_URL"}}},
				},
				// Command: []string{"sh", "-c", "apk add --no-cache envsubst && envsubst < /config/notifications-telegram.yaml > /config-envsubst/notifications-telegram.yaml"},
				Command: []string{
					"sh", "-c",
					`sed "s#\${TELEGRAM_BOT_TOKEN}#$TELEGRAM_BOT_TOKEN#g" /config/notifications-telegram.yaml > /config-envsubst/notifications-telegram.yaml` + " && " +
						`sed "s#\${SLACK_API_URL}#$SLACK_API_URL#g" /config/notifications-slack.yaml > /config-envsubst/notifications-slack.yaml`,
				},
				ExtraVolumeMounts: []corev1.VolumeMount{
					// {Name: "config", MountPath: "/config"},
					{Name: "config-envsubst", MountPath: "/config-envsubst"},
				},
			},
			{
				Name:  "prune-collections",
				Image: props.ImageInfo,
				// Command: []string{"sh", "-c", `cscli collections remove --all`},
				Env: map[string]string{
					"COLLECTIONS": strings.Join(collections, " "),
				},
				Command: []string{"sh", "-c", `
cscli collections list -o raw | cut -d',' -f1 | tail -n +2 | sort -u > /tmp/existing-collections;
echo $COLLECTIONS | tr ' ' '\n' | sort -u > /tmp/required-collections;
DIFF=$(comm -23 /tmp/existing-collections /tmp/required-collections | xargs);
if [ -n "$DIFF" ]; then
  echo "Pruning collections: $DIFF";
  cscli collections remove $DIFF;
else
  echo "No collections to prune";
fi;
`},
				ExtraVolumeMounts: infrahelpers.MergeLists(extraAcquisitionsVolMounts, []corev1.VolumeMount{
					{Name: "config", MountPath: "/etc/crowdsec/profiles.yaml", SubPath: "profiles.yaml", ReadOnly: true},
				}),
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
				Env: infrahelpers.MergeMaps(bouncerKeys, map[string]string{
					"GID":                  "1000",
					"ENROLL_INSTANCE_NAME": props.InstanceName,
					"ENROLL_TAGS":          "k8s",
					"COLLECTIONS":          strings.Join(collections, " "),
					"PARSERS":              strings.Join(parsers, " "),
					"SCEANRIOS":            strings.Join(scenarios, " "),
					"AGENT_USERNAME":       props.InstanceName,
					// "DISABLE_ONLINE_API":   "true", // If it's a test, we don't want to share signals with CrowdSec so disable the Online API.
					// "PARSERS": "crowdsecurity/cri-logs",
					// "DISABLE_PARSERS": "crowdsecurity/whitelists",
				}),
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false),
					Privileged:               infrahelpers.Ptr(false),
				},
				ExtraVolumeMounts: infrahelpers.MergeLists(extraAcquisitionsVolMounts, []corev1.VolumeMount{
					{Name: "container-logs", MountPath: "/var/log"},
					{Name: "config", MountPath: "/etc/crowdsec/profiles.yaml", SubPath: "profiles.yaml", ReadOnly: true},
					{Name: "config-envsubst", MountPath: "/etc/crowdsec/notifications/http.yaml", SubPath: "notifications-telegram.yaml"},
					{Name: "config-envsubst", MountPath: "/etc/crowdsec/notifications/slack.yaml", SubPath: "notifications-slack.yaml"},
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
					// https://docs.crowdsec.net/docs/profiles/captcha_profile
					"profiles.yaml": strings.Join(profiles, "---\n"),
					"notifications-telegram.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"type": "http",             // Don't change
						"name": "telegram_default", // Must match the registered plugin in the profile

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
						// https://docs.crowdsec.net/docs/notification_plugins/template_helpers
						// {{$alert.Source.IP}} ({{$alert.Source.Cn}}) triggered *{{$alert.Scenario}}* ({{$alert.Source.AsName}}) : Maliciousness Score is
						// {{- $cti := $alert.Source.IP | CrowdsecCTI  -}}
						// {{" "}}{{mulf $cti.GetMaliciousnessScore 100 | floor}} %
						// 						"format": infrahelpers.YAMLRawMessage(`{
						//     "chat_id": "{{ env "TELEGRAM_CHAT_ID" }}",
						//     "text": "
						// {{- range . }}
						// {{- $alert := . }}
						// 🌎 {{$alert.Source.Cn}}
						// 🌐 <a href="https://app.crowdsec.net/cti/{{$alert.Source.IP}}">{{$alert.Source.IP}}</a>
						// 📜 <i>{{$alert.Scenario}}</i>
						// 📝 <b>{{$alert.Source.AsName}}</b>
						// 🏷 Decisions:
						// {{- range .Decisions }}
						// {{.Value}} will get {{.Type}} for next {{.Duration}} for triggering {{.Scenario}}.
						// {{- end }}
						// {{- end -}}
						// ",
						//     "parse_mode": "HTML"
						// }`),
						// https://emojidb.org/decision-emojis
						//   🏷 Labels: <code>{{ $alert.Labels }}</code>
						"format": infrahelpers.YAMLRawMessage(`|
  {{- $telegramChatID := env "TELEGRAM_CHAT_ID" -}}
  {
    "chat_id": "{{ $telegramChatID }}",
    "text": "
  {{- range . }}
  {{- $alert := . }}
  🕵️ <b>{{$alert.Source.AsName}}</b>
  🌐 <a href=\"https://app.crowdsec.net/cti/{{$alert.Source.IP}}\">{{$alert.Source.IP}}</a> 🌎 {{$alert.Source.Cn}}
  📜 <i>{{$alert.Scenario}}</i>
  🎯 Decisions:
  {{- range .Decisions }}
    <code>{{.Value}}</code> will get <i>{{.Type}}</i> for next <b>{{.Duration}}</b> for triggering <i>{{.Scenario}}</i>.
  {{- end }}
  {{- end -}}
  ",
    "parse_mode": "HTML"
  }`),

						"url": "https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage", // Replace XXX:YYY with your API key

						"method": "POST",
						"headers": map[string]string{
							"Content-Type": "application/json",
						},
					}),
					"notifications-slack.yaml": infrahelpers.ToYamlString(map[string]interface{}{
						"type": "slack",         // Don't change
						"name": "slack_default", // Must match the registered plugin in the profile

						// One of "trace", "debug", "info", "warn", "error", "off"
						"log_level": "info",

						// This template receives list of models.Alert objects. The message would be composed from this
						"format": infrahelpers.YAMLRawMessage(`|
  {{range . -}}
  {{$alert := . -}}
  {{range .Decisions -}}
  {{if $alert.Source.Cn -}}
  :flag-{{$alert.Source.Cn | lower}}: <https://www.whois.com/whois/{{.Value}}|{{.Value}}> will get {{.Type}} for next {{.Duration}} for triggering {{.Scenario}} on machine '{{$alert.MachineID}}'. <https://www.shodan.io/host/{{.Value}}|Shodan>{{end}}
  {{if not $alert.Source.Cn -}}
  :pirate_flag: <https://www.whois.com/whois/{{.Value}}|{{.Value}}> will get {{.Type}} for next {{.Duration}} for triggering {{.Scenario}} on machine '{{$alert.MachineID}}'.  <https://www.shodan.io/host/{{.Value}}|Shodan>{{end}}
  {{end -}}
  {{end -}}
`),

						"webhook": "${SLACK_API_URL}",
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
					"SLACK_API_URL":      "SLACK_API_URL",
				},
			},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: "crowdsec-db", RequestsStorage: "1Gi", MountPath: "/var/lib/crowdsec/data", MountName: "crowdsec-db"},
			{Name: "crowdsec-config", RequestsStorage: "100Mi", MountPath: "/etc/crowdsec", MountName: "crowdsec-config"},
		},
	})
	crowdsecBouncerPluginConfig := map[string]interface{}{
		"Enabled": true,
		// "LogLevel":           "DEBUG",
		"CrowdsecMode":       "stream",
		"CrowdsecLapiScheme": "http",
		"CrowdsecLapiHost":   "crowdsec." + scope.Namespace() + ".svc.cluster.local:8080",
		"CrowdsecLapiKey":    props.BouncerKeys["TRAEFIK"],
		// "clienttrustedips": "10.0.10.30/32",
	}
	if props.ForwardedHeadersTrustedIPs != nil {
		crowdsecBouncerPluginConfig["forwardedHeadersTrustedIPs"] = props.ForwardedHeadersTrustedIPs
	}
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "crowdsec-traefik-bouncer"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Plugin: map[string]apiextensionv1.JSON{
				"crowdsec-bouncer": {
					Raw: []byte(infrahelpers.ToJSONString(crowdsecBouncerPluginConfig)),
				},
			},
		},
	})

	if props.NodePortService {
		scope.AddApiObject(&corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "crowdsec-np"},
			Spec: corev1.ServiceSpec{
				Type:     corev1.ServiceTypeNodePort,
				Ports:    []corev1.ServicePort{{Name: "http", Port: 8080, TargetPort: intstr.FromInt32(8080), NodePort: 31491}},
				Selector: map[string]string{"app.kubernetes.io/name": "crowdsec"},
			},
		})
	}

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
