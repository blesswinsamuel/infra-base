package k8sbase

import (
	"fmt"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/packager"

	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type AutheliaProps struct {
	ChartInfo k8sapp.ChartInfo `json:"helm"`
	Ingress   struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	AccessControl map[string]interface{} `json:"accessControl"`
	OIDC          struct {
		Enabled                bool                     `json:"enabled"`
		IssuerCertificateChain string                   `json:"issuer_certificate_chain"`
		Clients                []map[string]interface{} `json:"clients"`
	} `json:"oidc"`
	AuthMode string `json:"authMode"` // ldap or file
	LDAP     struct {
		BaseDN               string `json:"baseDN"`
		URL                  string `json:"url"`
		UsersFilter          string `json:"usersFilter"`
		GroupsFilter         string `json:"groupsFilter"`
		MailAttribute        string `json:"mailAttribute"`
		DisplayNameAttribute string `json:"displayNameAttribute"`
		User                 string `json:"user"`
		PasswordSecretKey    string `json:"passwordSecretKey"`
	} `json:"ldap"`
	SMTP struct {
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Username    string `json:"username"`
		EmailDomain string `json:"emailDomain"`
		Sender      string `json:"sender"`
		Subject     string `json:"subject"`
	} `json:"smtp"`
	Database struct {
		Postgres struct {
			Host     *string `json:"host"`
			Port     *int    `json:"port"`
			Database *string `json:"database"`
			Username *string `json:"username"`
			Schema   *string `json:"schema"`
		} `json:"postgres"`
		Redis struct {
			Host *string `json:"host"`
		} `json:"redis"`
	} `json:"database"`
	RedirectionSubDomain string `json:"redirectionSubDomain"`
}

// https://github.com/authelia/chartrepo/tree/master/charts/authelia

func (props *AutheliaProps) Chart(scope packager.Construct) packager.Construct {
	cprops := packager.ChartProps{
		Namespace: k8sapp.GetNamespaceContext(scope),
	}
	chart := scope.Chart("authelia", cprops)
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
		k8sapp.NewExternalSecret(chart, "users-db", &k8sapp.ExternalSecretProps{
			Name: "authelia-users-db",
			RemoteRefs: map[string]string{
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
	k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
		ChartInfo:   props.ChartInfo,
		ReleaseName: "authelia",
		Namespace:   chart.Namespace(),
		Values: map[string]interface{}{
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
				"annotations": infrahelpers.MergeAnnotations(
					GetCertIssuerAnnotation(scope),
					map[string]string{
						"traefik.ingress.kubernetes.io/router.middlewares": k8sapp.GetNamespaceContext(scope) + "-chain-authelia@kubernetescrd",
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
					infrahelpers.Ternary(props.RedirectionSubDomain != "", props.RedirectionSubDomain+".", "") +
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
						"sender": infrahelpers.UseOrDefault(
							props.SMTP.Sender,
							fmt.Sprintf("Authelia <authelia@%s>", props.SMTP.EmailDomain),
						),
						"identifier":            props.SMTP.EmailDomain,
						"subject":               infrahelpers.UseOrDefault(props.SMTP.Subject, "[authelia] {title}"),
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
						"url":                 props.LDAP.URL,
						"timeout":             "5s",
						"start_tls":           false,
						"base_dn":             props.LDAP.BaseDN,
						"username_attribute":  "uid",
						"additional_users_dn": "ou=people",
						// # users_filter: "(&({username_attribute}={input})(objectClass=person))"
						"users_filter":           props.LDAP.UsersFilter,
						"additional_groups_dn":   "ou=groups",
						"groups_filter":          props.LDAP.GroupsFilter,
						"group_name_attribute":   "cn",
						"mail_attribute":         props.LDAP.MailAttribute,
						"display_name_attribute": props.LDAP.DisplayNameAttribute,
						"user":                   props.LDAP.User + "," + props.LDAP.BaseDN,
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
		secrets["LDAP_PASSWORD"] = props.LDAP.PasswordSecretKey
	}
	if props.OIDC.Enabled {
		secrets["OIDC_PRIVATE_KEY"] = "AUTHELIA_OIDC_PRIVATE_KEY"
		secrets["OIDC_HMAC_SECRET"] = "AUTHELIA_OIDC_HMAC_SECRET"
	}
	k8sapp.NewExternalSecret(chart, "external-secrets", &k8sapp.ExternalSecretProps{
		Name:       "authelia",
		RemoteRefs: secrets,
	})
	return chart
}
