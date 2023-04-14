package k8sbase

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sbase/imports/k8s"
)

type LLDAPProps struct {
	Enabled     bool      `yaml:"enabled"`
	ImageInfo   ImageInfo `yaml:"image"`
	BaseDN      string    `yaml:"base_dn"`
	EmailDomain string    `yaml:"email_domain"`
}

func NewLLDAP(scope constructs.Construct, props LLDAPProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	return NewApplication(scope, jsii.String("lldap"), &ApplicationProps{
		Name: "lldap",
		Container: ApplicationContainer{
			Name:      "lldap",
			ImageInfo: props.ImageInfo,
			Command: []string{
				"/bin/sh", "-c",
				`echo -n "$LLDAP_PRIVATE_KEY" | base64 -d  > "$LLDAP_KEY_FILE"  && /app/lldap run --config-file /app/lldap_config.docker_template.toml`,
			},
			Ports: []ContainerPort{
				{Name: "web", Port: 17170},
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
			LivenessProbe: &k8s.Probe{
				HttpGet: &k8s.HttpGetAction{Path: jsii.String("/health"), Port: k8s.IntOrString_FromNumber(jsii.Number(17170))},
			},
			ReadinessProbe: &k8s.Probe{
				HttpGet: &k8s.HttpGetAction{Path: jsii.String("/health"), Port: k8s.IntOrString_FromNumber(jsii.Number(17170))},
			},
		},
		Ingress: []ApplicationIngress{
			{Host: "lldap." + GetDomain(scope), PortName: "web"},
		},
		ExternalSecrets: []ApplicationExternalSecret{
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
