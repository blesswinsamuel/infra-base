package k8sbase

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/blesswinsamuel/infra-base/k8sapp"
	"github.com/blesswinsamuel/infra-base/k8sbase/helpers"
	"github.com/blesswinsamuel/infra-base/k8sbase/infraglobal"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type AutheliaProps struct {
	Enabled   bool              `yaml:"enabled"`
	ChartInfo helpers.ChartInfo `yaml:"helm"`
	Ingress   struct {
		SubDomain string `yaml:"subDomain"`
	} `yaml:"ingress"`
	AccessControl map[string]interface{} `yaml:"accessControl"`
	OIDC          struct {
		Enabled                bool                     `yaml:"enabled"`
		IssuerCertificateChain string                   `yaml:"issuer_certificate_chain"`
		Clients                []map[string]interface{} `yaml:"clients"`
	} `yaml:"oidc"`
	AuthMode string `yaml:"authMode"` // ldap or file
	LDAP     struct {
		BaseDN string `yaml:"baseDN"`
	} `yaml:"ldap"`
	SMTP struct {
		Host        string  `yaml:"host"`
		Port        int     `yaml:"port"`
		Username    string  `yaml:"username"`
		EmailDomain string  `yaml:"emailDomain"`
		Sender      *string `yaml:"sender"`
		Subject     *string `yaml:"subject"`
	} `yaml:"smtp"`
	Database struct {
		Postgres struct {
			Host     *string `yaml:"host"`
			Port     *int    `yaml:"port"`
			Database *string `yaml:"database"`
			Username *string `yaml:"username"`
			Schema   *string `yaml:"schema"`
		} `yaml:"postgres"`
		Redis struct {
			Host *string `yaml:"host"`
		} `yaml:"redis"`
	} `yaml:"database"`
	RedirectionSubDomain string `yaml:"redirectionSubDomain"`
}

// https://github.com/authelia/chartrepo/tree/master/charts/authelia

func NewAuthelia(scope constructs.Construct, props AutheliaProps) constructs.Construct {
	if !props.Enabled {
		return nil
	}
	cprops := cdk8s.ChartProps{
		Namespace: k8sapp.GetNamespaceContextPtr(scope),
	}
	chart := cdk8s.NewChart(scope, jsii.String("authelia"), &cprops)
	pod := map[string]interface{}{
		"annotations": map[string]interface{}{
			"secret.reloader.stakater.com/reload": "authelia",
		},
		"kind": "Deployment",
		"env": []map[string]interface{}{
			{"name": "TZ", "value": "UTC"},
		},
	}
	if props.AuthMode == "file" {
		NewExternalSecret(chart, jsii.String("users-db"), &ExternalSecretProps{
			Name:            jsii.String("authelia-users-db"),
			Namespace:       k8sapp.GetNamespaceContextPtr(scope),
			RefreshInterval: jsii.String("10m"),
			Secrets: map[string]string{
				"users_database.yml": "AUTHELIA_USERS_DATABASE_YML",
			},
		})
		pod["extraVolumeMounts"] = []map[string]interface{}{
			{
				"name":      "authelia-users-db",
				"mountPath": "/config/users_database.yml",
				"subPath":   "users_database.yml",
			},
		}
		pod["extraVolumes"] = []map[string]interface{}{
			{
				"name": "authelia-users-db",
				"secret": map[string]interface{}{
					"secretName": "authelia-users-db",
				},
			},
		}
	}
	helpers.NewHelmCached(chart, jsii.String("helm"), &helpers.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: jsii.String("authelia"),
		Namespace:   chart.Namespace(),
		Values: &map[string]interface{}{
			"domain": GetDomain(scope),
			"pod":    pod,
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
				"annotations": helpers.MergeAnnotations(
					infraglobal.GetCertIssuerAnnotation(scope),
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
				"default_redirection_url": "https://" +
					helpers.Ternary(props.RedirectionSubDomain != "", props.RedirectionSubDomain+".", "") +
					GetDomain(scope),
				"access_control": props.AccessControl,
				"session": map[string]interface{}{
					"redis": map[string]interface{}{
						"host": props.Database.Redis.Host,
					},
				},
				"storage": map[string]interface{}{
					"postgres": map[string]interface{}{
						"host":     props.Database.Postgres.Host,
						"port":     props.Database.Postgres.Port,
						"database": props.Database.Postgres.Database,
						"schema":   props.Database.Postgres.Schema,
						"username": props.Database.Postgres.Username,
					},
				},
				"notifier": map[string]interface{}{
					"smtp": map[string]interface{}{
						"host":     props.SMTP.Host,
						"port":     props.SMTP.Port,
						"username": props.SMTP.Username,
						"sender": helpers.Fallback(
							props.SMTP.Sender,
							fmt.Sprintf("Authelia <authelia@%s>", props.SMTP.EmailDomain),
						),
						"identifier":            props.SMTP.EmailDomain,
						"subject":               helpers.Fallback(props.SMTP.Subject, "[authelia] {title}"),
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
						"enabled":             props.AuthMode == "ldap",
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
					"file": map[string]interface{}{
						"enabled": props.AuthMode == "file",
					},
				},
				"identity_providers": map[string]interface{}{
					"oidc": map[string]interface{}{
						"enabled": props.OIDC.Enabled,
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

	secrets := map[string]string{
		"SMTP_PASSWORD":          "SMTP_PASSWORD",
		"JWT_TOKEN":              "AUTHELIA_JWT_TOKEN",
		"SESSION_ENCRYPTION_KEY": "AUTHELIA_SESSION_ENCRYPTION_KEY",
		"STORAGE_ENCRYPTION_KEY": "AUTHELIA_STORAGE_ENCRYPTION_KEY",
		"STORAGE_PASSWORD":       "POSTGRES_USER_PASSWORD",
	}
	if props.AuthMode == "ldap" {
		secrets["LDAP_PASSWORD"] = "LLDAP_LDAP_USER_PASS"
	}
	if props.OIDC.Enabled {
		secrets["OIDC_PRIVATE_KEY"] = "AUTHELIA_OIDC_PRIVATE_KEY"
		secrets["OIDC_HMAC_SECRET"] = "AUTHELIA_OIDC_HMAC_SECRET"
	}
	NewExternalSecret(chart, jsii.String("external-secrets"), &ExternalSecretProps{
		Name:            jsii.String("authelia"),
		Namespace:       k8sapp.GetNamespaceContextPtr(scope),
		RefreshInterval: jsii.String("10m"),
		Secrets:         secrets,
	})
	return chart
}
