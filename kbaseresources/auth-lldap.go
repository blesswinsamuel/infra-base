package kbaseresources

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
	UseMailpit  bool             `json:"useMailpit"`
	Postgres    struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Database    string `json:"database"`
		Username    string `json:"username"`
		PasswordRef string `json:"passwordRef"`
	} `json:"postgres"`
}

func (props *LLDAP) Render(scope kgen.Scope) {
	app := &k8sapp.ApplicationProps{
		Name: "lldap",
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "lldap",
			Image: props.ImageInfo,
			Command: []string{
				"/bin/sh", "-c",
				`echo -n "$LLDAP_PRIVATE_KEY" | base64 -d  > "$LLDAP_KEY_FILE"  && /app/lldap run --config-file /app/lldap_config.docker_template.toml`,
			},
			Ports: []k8sapp.ContainerPort{
				{Name: "web", Port: 17170, Ingress: &k8sapp.ApplicationIngress{Host: fmt.Sprintf("lldap.%s", k8sapp.GetDomain(scope))}},
				{Name: "ldap", Port: 3890},
			},
			Env: map[string]string{
				// "LLDAP_VERBOSE": "true",
				"LLDAP_SMTP_OPTIONS__ENABLE_PASSWORD_RESET": "true",
				"LLDAP_SMTP_OPTIONS__FROM":                  fmt.Sprintf("LLDAP <lldap@%s>", props.EmailDomain),
				"LLDAP_SMTP_OPTIONS__REPLY_TO":              fmt.Sprintf("LLDAP no-reply <lldap-no-reply@%s>", props.EmailDomain),
				"LLDAP_LDAP_BASE_DN":                        props.BaseDN,
				"LLDAP_LDAP_USER_DN":                        "admin",
				"LLDAP_LDAP_USER_EMAIL":                     fmt.Sprintf("admin@%s", props.EmailDomain),
				"LLDAP_KEY_FILE":                            "/data/private_key",
				"LLDAP_HTTP_URL":                            "https://lldap." + k8sapp.GetDomain(scope),
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
					"LLDAP_JWT_SECRET":     "LLDAP_JWT_SECRET",
					"LLDAP_LDAP_USER_PASS": "LLDAP_LDAP_USER_PASS",
					"LLDAP_PRIVATE_KEY":    "LLDAP_PRIVATE_KEY",
				},
			},
			{
				Name: "lldap-postgres",
				Template: map[string]string{
					"LLDAP_DATABASE_URL": fmt.Sprintf("postgres://%s:%s@%s:%d/%s", props.Postgres.Username, "{{ .PGPASSWORD | urlquery }}", props.Postgres.Host, props.Postgres.Port, props.Postgres.Database),
				},
				RemoteRefs: map[string]string{
					"PGPASSWORD": props.Postgres.PasswordRef,
				},
			},
		},
		Homepage: &k8sapp.ApplicationHomepage{
			Name:        "Light LDAP",
			Description: "LDAP server",
			SiteMonitor: "http://lldap." + scope.Namespace() + ".svc.cluster.local:17170/health",
			Group:       "Infra",
			Icon:        "authentik",
		},
		NetworkPolicy: &k8sapp.ApplicationNetworkPolicy{
			Ingress: k8sapp.NetworkPolicyIngress{
				AllowFromAppRefs: map[string][]intstr.IntOrString{
					"authelia": {intstr.FromString("ldap")},
				},
			},
			Egress: k8sapp.NetworkPolicyEgress{
				AllowToAppRefs: []string{"postgres"},
			},
		},
	}
	if props.UseMailpit {
		app.Containers[0].Env["LLDAP_SMTP_OPTIONS__SERVER"] = "mailpit.smtp.svc.cluster.local."
		app.Containers[0].Env["LLDAP_SMTP_OPTIONS__PORT"] = "1025"
		app.NetworkPolicy.Egress.AllowToAppRefs = append(app.NetworkPolicy.Egress.AllowToAppRefs, "mailpit")
	} else {
		app.Containers[0].Env["LLDAP_SMTP_OPTIONS__SMTP_ENCRYPTION"] = "STARTTLS"
		app.ExternalSecrets[0].RemoteRefs["LLDAP_SMTP_OPTIONS__USER"] = "SMTP_USERNAME"
		app.ExternalSecrets[0].RemoteRefs["LLDAP_SMTP_OPTIONS__PASSWORD"] = "SMTP_PASSWORD"
		app.ExternalSecrets[0].RemoteRefs["LLDAP_SMTP_OPTIONS__SERVER"] = "SMTP_HOST"
		app.ExternalSecrets[0].RemoteRefs["LLDAP_SMTP_OPTIONS__PORT"] = "SMTP_PORT"
	}
	k8sapp.NewApplication(scope, app)
}
