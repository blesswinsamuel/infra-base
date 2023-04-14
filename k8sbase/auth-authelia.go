package k8sbase

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type AutheliaProps struct {
	Enabled   bool      `yaml:"enabled"`
	ChartInfo ChartInfo `yaml:"helm"`
	Ingress   struct {
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
	Rules []map[string]interface{} `yaml:"rules"`
	OIDC  struct {
		IssuerCertificateChain string                   `yaml:"issuer_certificate_chain"`
		Clients                []map[string]interface{} `yaml:"clients"`
	} `yaml:"oidc"`
	LDAP struct {
		BaseDN string `yaml:"baseDN"`
	} `yaml:"ldap"`
	SMTP struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Username    string `yaml:"username"`
		EmailDomain string `yaml:"emailDomain"`
	} `yaml:"smtp"`
}

func NewAuthelia(scope constructs.Construct, props AutheliaProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: GetNamespace(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("authelia"), &cprops)
	NewHelmCached(chart, jsii.String("helm"), &HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: jsii.String("authelia"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"domain": GetDomain(scope),
			"pod": map[string]interface{}{
				"annotations": map[string]interface{}{
					"secret.reloader.stakater.com/reload": "authelia",
				},
				"kind": "Deployment",
				"env": []map[string]interface{}{
					{"name": "TZ", "value": "UTC"},
				},
			},
			"service": map[string]interface{}{
				"annotations": map[string]interface{}{
					"prometheus.io/scrape": "true",
					"prometheus.io/port":   "9959",
					"prometheus.io/path":   "/metrics",
				},
			},
			"ingress": map[string]interface{}{
				"enabled":   true,
				"subdomain": props.Ingress.SubDomain,
				"traefikCRD": map[string]interface{}{
					"enabled":             true,
					"disableIngressRoute": true,
				},
				"annotations": MergeAnnotations(
					GetCertIssuerAnnotation(scope),
					map[string]string{
						"traefik.ingress.kubernetes.io/router.middlewares": "auth-chain-authelia@kubernetescrd",
					},
				),
			},
			"secret": map[string]interface{}{
				"existingSecret": "authelia",
			},
			"configMap": map[string]interface{}{
				"telemetry": map[string]interface{}{
					"metrics": map[string]interface{}{
						"enabled": true,
					},
				},
				"regulation": map[string]interface{}{
					"max_retries": 3,
					"find_time":   "2m",
					"ban_time":    "5m",
				},
				"default_redirection_url": "https://dashy." + GetDomain(scope),
				"access_control": map[string]interface{}{
					"default_policy": "deny",
					"rules":          props.Rules,
				},
				"session": map[string]interface{}{
					"redis": map[string]interface{}{
						"host": "redis-master.database.svc.cluster.local",
					},
				},
				"storage": map[string]interface{}{
					"postgres": map[string]interface{}{
						"host":     "postgres.database.svc.cluster.local",
						"username": "homelab",
					},
				},
				"notifier": map[string]interface{}{
					"smtp": map[string]interface{}{
						"host":                  props.SMTP.Host,
						"port":                  props.SMTP.Port,
						"username":              props.SMTP.Username,
						"sender":                fmt.Sprintf("Authelia <authelia@%s>", props.SMTP.EmailDomain),
						"identifier":            props.SMTP.EmailDomain,
						"subject":               "[authelia] {title}",
						"startup_check_address": fmt.Sprintf("test@%s", props.SMTP.EmailDomain),
						"enabledSecret":         true,
					},
				},
				"authentication_backend": map[string]interface{}{
					"password_reset": map[string]interface{}{
						"disable": false,
					},
					// # How often authelia should check if there is an user update in LDAP
					// # refresh_interval: 1m
					"refresh_interval": "always",
					// https://github.com/nitnelave/lldap/blob/main/example_configs/authelia_config.yml
					"ldap": map[string]interface{}{
						"enabled":             true,
						"implementation":      "custom",
						"url":                 "ldap://lldap:3890",
						"timeout":             "5s",
						"start_tls":           false,
						"base_dn":             props.LDAP.BaseDN,
						"username_attribute":  "uid",
						"additional_users_dn": "ou=people",
						// # users_filter: "(&({username_attribute}={input})(objectClass=person))"
						"users_filter":           "(&({username_attribute}={input})(!({username_attribute}=admin))(objectClass=person))",
						"additional_groups_dn":   "ou=groups",
						"groups_filter":          "(member={dn})",
						"group_name_attribute":   "cn",
						"mail_attribute":         "mail",
						"display_name_attribute": "displayName",
						"user":                   "uid=admin,ou=people," + props.LDAP.BaseDN,
					},
				},
				"identity_providers": map[string]interface{}{
					"oidc": map[string]interface{}{
						"enabled": true,
						"cors": map[string]interface{}{
							"endpoints": []string{
								"token",
								"userinfo",
								// below ones may not be needed
								"authorization",
								"revocation",
								"introspection",
							},
						},
						"issuer_certificate_chain": props.OIDC.IssuerCertificateChain,
						"clients":                  props.OIDC.Clients,
					},
				},
			},
		},
	})

	NewExternalSecret(chart, jsii.String("external-secrets"), &ExternalSecretProps{
		Name:            jsii.String("authelia"),
		Namespace:       GetNamespace(scope),
		RefreshInterval: jsii.String("10m"),
		Secrets: map[string]string{
			"SMTP_PASSWORD":          "SMTP_PASSWORD",
			"JWT_TOKEN":              "AUTHELIA_JWT_TOKEN",
			"SESSION_ENCRYPTION_KEY": "AUTHELIA_SESSION_ENCRYPTION_KEY",
			"STORAGE_ENCRYPTION_KEY": "AUTHELIA_STORAGE_ENCRYPTION_KEY",
			"STORAGE_PASSWORD":       "POSTGRES_USER_PASSWORD",
			"OIDC_PRIVATE_KEY":       "AUTHELIA_OIDC_PRIVATE_KEY",
			"OIDC_HMAC_SECRET":       "AUTHELIA_OIDC_HMAC_SECRET",
			"LDAP_PASSWORD":          "LLDAP_LDAP_USER_PASS",
		},
	})
	return chart
}
