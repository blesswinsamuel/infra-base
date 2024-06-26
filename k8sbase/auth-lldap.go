package k8sbase

import (
	"fmt"

	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/kgen"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type LLDAP struct {
	ImageInfo   k8sapp.ImageInfo `json:"image"`
	BaseDN      string           `json:"base_dn"`
	EmailDomain string           `json:"email_domain"`
}

func (props *LLDAP) Render(scope kgen.Scope) {
	k8sapp.NewApplication(scope, &k8sapp.ApplicationProps{
		Name: "lldap",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "lldap",
			Image: props.ImageInfo,
			Command: []string{
				"/bin/sh", "-c",
				`echo -n "$LLDAP_PRIVATE_KEY" | base64 -d  > "$LLDAP_KEY_FILE"  && /app/lldap run --config-file /app/lldap_config.docker_template.toml`,
			},
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 17170, Ingress: &k8sapp.ApplicationIngress{Host: fmt.Sprintf("lldap.%s", GetDomain(scope))}},
				{Name: "ldap", Port: 3890},
			},
			Env: map[string]string{
				// "LLDAP_VERBOSE": "true",
				"LLDAP_SMTP_OPTIONS__ENABLE_PASSWORD_RESET": "true",
				"LLDAP_SMTP_OPTIONS__SMTP_ENCRYPTION":       "STARTTLS",
				"LLDAP_SMTP_OPTIONS__FROM":                  fmt.Sprintf("LLDAP <lldap@%s>", props.EmailDomain),
				"LLDAP_SMTP_OPTIONS__REPLY_TO":              fmt.Sprintf("LLDAP no-reply <lldap-no-reply@%s>", props.EmailDomain),
				"LLDAP_LDAP_BASE_DN":                        props.BaseDN,
				"LLDAP_LDAP_USER_DN":                        "admin",
				"LLDAP_LDAP_USER_EMAIL":                     fmt.Sprintf("admin@%s", props.EmailDomain),
				"LLDAP_KEY_FILE":                            "/data/private_key",
				"LLDAP_HTTP_URL":                            "https://lldap." + GetDomain(scope),
			},
			EnvFromSecretRef: []string{
				"lldap", "lldap-postgres",
			},
			LivenessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/health", Port: intstr.FromInt(17170)},
			}},
			ReadinessProbe: &corev1.Probe{ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/health", Port: intstr.FromInt(17170)},
			}},
		}},
		ExternalSecrets: []k8sapp.ApplicationExternalSecret{
			{
				Name: "lldap",
				RemoteRefs: map[string]string{
					"LLDAP_JWT_SECRET":             "LLDAP_JWT_SECRET",
					"LLDAP_LDAP_USER_PASS":         "LLDAP_LDAP_USER_PASS",
					"LLDAP_SMTP_OPTIONS__USER":     "SMTP_USERNAME",
					"LLDAP_SMTP_OPTIONS__PASSWORD": "SMTP_PASSWORD",
					"LLDAP_SMTP_OPTIONS__SERVER":   "SMTP_HOST",
					"LLDAP_SMTP_OPTIONS__PORT":     "SMTP_PORT",
					"LLDAP_PRIVATE_KEY":            "LLDAP_PRIVATE_KEY",
				},
			},
			{
				Name: "lldap-postgres",
				Template: map[string]string{
					"LLDAP_DATABASE_URL": "postgres://{{ .PGUSER }}:{{ .PGPASSWORD | urlquery }}@postgres.database.svc.cluster.local:5432/lldap",
				},
				RemoteRefs: map[string]string{
					"PGPASSWORD": "POSTGRES_USER_PASSWORD",
					"PGUSER":     "POSTGRES_USERNAME",
				},
			},
		},
	})
}
