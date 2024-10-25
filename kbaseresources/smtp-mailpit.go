package kbaseresources

import (
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func init() {
	RegisterModule("mailpit", &Mailpit{})
}

type Mailpit struct {
	ImageInfo            k8sapp.ImageInfo `json:"image"`
	PersistentVolumeName string           `json:"persistentVolumeName"`
}

func (props *Mailpit) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "mailpit",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 8025, Ingress: &k8sapp.ApplicationIngress{Host: scope.ID() + "." + k8sapp.GetDomain(scope)}},
				{Name: "smtp", Port: 1025},
			},
			Env: map[string]string{
				"MP_MAX_MESSAGES":             "5000",
				"MP_DATABASE":                 "/data/mailpit.db",
				"MP_SMTP_AUTH_ACCEPT_ANY":     "1",
				"MP_SMTP_AUTH_ALLOW_INSECURE": "1",
				// "MP_SMTP_RELAY_CONFIG":        "/config/smtp-relay-config.yaml",
				"MP_SMTP_RELAY_STARTTLS": "true",
				"MP_SMTP_RELAY_AUTH":     "plain",
				"MP_SMTP_RELAY_ALL":      "true",
			},
			EnvFromSecretRef: []string{
				scope.ID() + "-smtp",
			},
			LivenessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/livez", Port: intstr.FromString("web")},
			}},
			ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/readyz", Port: intstr.FromString("web")},
			}},
		}},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: scope.ID() + "-smtp",
				RemoteRefs: map[string]string{
					"MP_SMTP_RELAY_HOST":     "SMTP_HOST",
					"MP_SMTP_RELAY_PORT":     "SMTP_PORT",
					"MP_SMTP_RELAY_USERNAME": "SMTP_USERNAME",
					"MP_SMTP_RELAY_PASSWORD": "SMTP_PASSWORD",
				},
			},
		},
		PersistentVolumes: []k8sapp.ApplicationPersistentVolume{
			{Name: scope.ID(), VolumeName: props.PersistentVolumeName, RequestsStorage: "2Gi", MountName: "data", MountPath: "/data"},
		},
		Homepage: &k8sapp.ApplicationHomepage{
			Name:        "Mailpit",
			Description: "Email relay service",
			SiteMonitor: "http://" + scope.ID() + "." + scope.Namespace() + ".svc.cluster.local:8025/readyz",
			Group:       "Infra",
			Icon:        "mdi-email-fast",
		},
		NetworkPolicy: &k8sapp.ApplicationNetworkPolicy{
			Ingress: k8sapp.NetworkPolicyIngress{
				AllowFromAllNamespaces: []intstr.IntOrString{intstr.FromString("smtp")},
			},
			Egress: k8sapp.NetworkPolicyEgress{
				AllowToAllInternet: []int{2525, 25, 587, 465}, // for smtp relay
			},
		},
	})
}
