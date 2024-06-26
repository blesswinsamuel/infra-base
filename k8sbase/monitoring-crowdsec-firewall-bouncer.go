package k8sbase

import (
	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
)

type CrowdsecFirewallBouncer struct {
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Mode      string           `json:"mode"`
}

func (props *CrowdsecFirewallBouncer) Render(scope kgen.Scope) {
	// https://github.com/crowdsecurity/cs-firewall-bouncer/issues/32#issuecomment-1060890534
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name:        "crowdsec-firewall-bouncer",
		HostNetwork: true,
		Kind:        "DaemonSet",
		Containers: []k8sapp.ApplicationContainer{
			{
				Name:    "crowdsec-firewall-bouncer",
				Command: []string{"crowdsec-firewall-bouncer", "-c", "/config/crowdsec-firewall-bouncer.yaml"},
				Image:   props.ImageInfo,
				Env:     map[string]string{},
				SecurityContext: &corev1.SecurityContext{
					AllowPrivilegeEscalation: infrahelpers.Ptr(false),
					Privileged:               infrahelpers.Ptr(false),
					Capabilities: &corev1.Capabilities{
						Add: []corev1.Capability{"NET_ADMIN", "NET_RAW"},
					},
				},
			},
		},
		DNSPolicy: corev1.DNSClusterFirstWithHostNet,
		ConfigMaps: []k8sapp.ApplicationConfigMap{
			{
				Name: "crowdsec-firewall-bouncer",
				Data: map[string]string{
					// nft list ruleset
					"crowdsec-firewall-bouncer.yaml": infrahelpers.ToYamlString(map[string]any{
						// "mode":             "iptables",
						"mode":             props.Mode,
						"update_frequency": "10s",
						// "log_mode":         "file",
						// "log_dir":          "/var/log/",
						"log_level": "info",
						// "log_compression":  true,
						// "log_max_size":     100,
						// "log_max_backups":  3,
						// "log_max_age":      30,
						"api_url": "http://crowdsec." + scope.Namespace() + ".svc.cluster.local:8080",
						"api_key": "mysecretkey12345",
						// "insecure_skip_verify": false,
						"disable_ipv6": true,
						"deny_action":  "DROP",
						"deny_log":     false,
						"supported_decisions_types": []string{
							"ban-firewall",
						},
						"blacklists_ipv4": "crowdsec-blacklists",
						"blacklists_ipv6": "crowdsec6-blacklists",
						"ipset_type":      "nethash",
						"iptables_chains": []string{
							"INPUT",
						},
						"nftables": map[string]map[string]any{
							"ipv4": {
								"enabled":  true,
								"set-only": false,
								"table":    "crowdsec",
								"chain":    "crowdsec-chain",
								"priority": -10,
							},
							"ipv6": {
								"enabled":  true,
								"set-only": false,
								"table":    "crowdsec6",
								"chain":    "crowdsec6-chain",
								"priority": -10,
							},
						},
						"nftables_hooks": []string{
							"input",
							"forward",
						},
						// "prometheus": map[string]any{
						// 	"enabled":     false,
						// 	"listen_addr": "127.0.0.1",
						// 	"listen_port": 60601,
						// },
					}),
				},
				MountName: "config",
				MountPath: "/config",
				ReadOnly:  true,
			},
		},
	})
}
