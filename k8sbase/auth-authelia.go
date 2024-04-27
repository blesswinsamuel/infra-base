package k8sbase

import (
	"fmt"
	"strings"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"golang.org/x/exp/slices"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type AutheliaClient struct {
	AuthorizationPolicy          string   `json:"authorization_policy"`
	ConsentMode                  string   `json:"consent_mode"`
	ClientName                   string   `json:"client_name"`
	GrantTypes                   []string `json:"grant_types"`
	ClientID                     string   `json:"client_id"`
	PreConfiguredConsentDuration string   `json:"pre_configured_consent_duration"`
	Public                       bool     `json:"public"`
	RedirectURIs                 []string `json:"redirect_uris"`
	ResponseTypes                []string `json:"response_types"`
	Scopes                       []string `json:"scopes"`
	ClientSecret                 string   `json:"client_secret,omitempty"`
	SectorIdentifier             string   `json:"sector_identifier,omitempty"`
	UserinfoSignedResponseAlg    string   `json:"userinfo_signed_response_alg"`
}

func (c *AutheliaClient) FillDefaults() {
	if c.ConsentMode == "" {
		c.ConsentMode = "auto"
	}
	if c.GrantTypes == nil {
		if slices.Contains(c.Scopes, "offline_access") {
			c.GrantTypes = []string{"refresh_token", "authorization_code"}
		} else {
			// option 'grant_types' should only have the 'refresh_token' value if the client is also configured with the 'offline_access' scope
			c.GrantTypes = []string{"authorization_code"}
		}
	}
	if c.ResponseTypes == nil {
		c.ResponseTypes = []string{"code"}
	}
	if c.UserinfoSignedResponseAlg == "" {
		c.UserinfoSignedResponseAlg = "none"
	}
	if c.Scopes == nil {
		c.Scopes = []string{"openid", "profile", "email", "groups"}
	}
}

type Authelia struct {
	ChartInfo k8sapp.ChartInfo `json:"helm"`
	ImageInfo k8sapp.ImageInfo `json:"image"`
	Ingress   struct {
		SubDomain string `json:"subDomain"`
	} `json:"ingress"`
	AccessControl map[string]any `json:"accessControl"`
	OIDC          struct {
		Enabled                bool             `json:"enabled"`
		IssuerCertificateChain string           `json:"issuer_certificate_chain"`
		Clients                []AutheliaClient `json:"clients"`
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
		EmailDomain string `json:"emailDomain"`
		Sender      string `json:"sender"`
		Subject     string `json:"subject"`
	} `json:"smtp"`
	Database struct {
		Postgres struct {
			Host     *string `json:"host"`
			Port     *int    `json:"port"`
			Database *string `json:"database"`
			Schema   *string `json:"schema"`
		} `json:"postgres"`
		Redis struct {
			Host *string `json:"host"`
		} `json:"redis"`
	} `json:"database"`
	Assets *struct {
		LogoURL    string `json:"logoURL"`
		FaviconURL string `json:"faviconURL"`
	} `json:"assets"`
	RedirectionSubDomain         string           `json:"redirectionSubDomain"`
	CookieDomains                []map[string]any `json:"cookieDomains"`
	IncludeForwardAuthMiddleware bool             `json:"includeForwardAuthMiddleware"`
}

// https://github.com/authelia/chartrepo/tree/master/charts/authelia

func (props *Authelia) Render(scope kubegogen.Scope) {
	ingressMiddlewares := []k8sapp.NameNamespace{}
	if props.IncludeForwardAuthMiddleware {
		ingressMiddlewares = append(ingressMiddlewares, k8sapp.NameNamespace{Name: "forwardauth-authelia", Namespace: scope.Namespace()})
	}
	ingressMiddlewares = append(ingressMiddlewares, k8sapp.NameNamespace{Name: "chain-authelia", Namespace: scope.Namespace()})
	appProps := &k8sapp.ApplicationProps{
		Name:               "authelia",
		IngressMiddlewares: ingressMiddlewares,
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "authelia",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 9091, ServicePort: 80, Ingress: &k8sapp.ApplicationIngress{Host: props.Ingress.SubDomain + "." + GetDomain(scope)}},
				{Name: "metrics", Port: 9959, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{Path: "/metrics"}},
			},
			Args: []string{
				"--config=/configuration.yaml",
			},
			Command: []string{"authelia"},
			Env: map[string]string{
				"AUTHELIA_SERVER_DISABLE_HEALTHCHECK": "true",
				"TZ":                                  "UTC",
			},
			LivenessProbe:  &corev1.Probe{FailureThreshold: 5, PeriodSeconds: 30, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
			ReadinessProbe: &corev1.Probe{FailureThreshold: 5, PeriodSeconds: 5, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
			StartupProbe:   &corev1.Probe{FailureThreshold: 6, InitialDelaySeconds: 10, PeriodSeconds: 5, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
		}},
		ExternalSecrets:    []k8sapp.ApplicationExternalSecret{},
		EnableServiceLinks: infrahelpers.Ptr(false),
	}
	redirectionURL := "https://" + infrahelpers.Ternary(props.RedirectionSubDomain != "", props.RedirectionSubDomain+".", "") + GetDomain(scope)
	if len(props.CookieDomains) == 0 {
		props.CookieDomains = append(props.CookieDomains, map[string]any{
			"domain": GetDomain(scope),
		})
	}
	for i := range props.CookieDomains {
		if props.CookieDomains[i]["default_redirection_url"] == nil {
			props.CookieDomains[i]["default_redirection_url"] = redirectionURL
		}
		if props.CookieDomains[i]["authelia_url"] == nil {
			props.CookieDomains[i]["authelia_url"] = "https://" + props.Ingress.SubDomain + "." + GetDomain(scope)
		}
	}
	// https://github.com/authelia/chartrepo/blob/master/charts/authelia/templates/configMap.yaml
	configMap := map[string]any{
		"theme": "light",
		"server": map[string]any{
			"address":    "tcp://0.0.0.0:9091/",
			"asset_path": "",
			"headers":    map[string]any{"csp_template": ""},
			"buffers":    map[string]any{"read": 4096, "write": 4096},
			"timeouts":   map[string]any{"read": "6s", "write": "6s", "idle": "30s"},
			"endpoints": map[string]any{
				"enable_pprof":   false,
				"enable_expvars": false,
			},
		},
		"log": map[string]any{
			"level":       "info",
			"format":      "text",
			"file_path":   "",
			"keep_stdout": true,
		},
		"telemetry": map[string]any{
			"metrics": map[string]any{
				"enabled": true,
				"address": "tcp://0.0.0.0:9959",
				"buffers": map[string]any{"read": 4096, "write": 4096},
			},
		},
		"totp": map[string]any{
			"disable":     false,
			"issuer":      GetDomain(scope),
			"algorithm":   "sha1",
			"digits":      6,
			"period":      30,
			"skew":        1,
			"secret_size": 32,
		},
		"webauthn": map[string]any{
			"disable":                           false,
			"display_name":                      "Authelia",
			"attestation_conveyance_preference": "indirect",
			"user_verification":                 "preferred",
			"timeout":                           "60s",
		},
		"ntp": map[string]any{
			"address":               "udp://time.cloudflare.com:123",
			"version":               4,
			"max_desync":            "3s",
			"disable_startup_check": false,
			"disable_failure":       false,
		},
		"regulation": map[string]any{
			"max_retries": 3,
			"find_time":   "2m",
			"ban_time":    "5m",
		},
		"default_2fa_method": "",
		"access_control":     props.AccessControl,
		"session": map[string]any{
			"name":        "authelia_session",
			"secret":      "{{ .SESSION_ENCRYPTION_KEY }}",
			"same_site":   "lax",
			"expiration":  "1h",
			"inactivity":  "5m",
			"remember_me": "1M",
			"cookies":     props.CookieDomains,
			"redis": map[string]any{
				"host":                       props.Database.Redis.Host,
				"port":                       6379,
				"database_index":             0,
				"maximum_active_connections": 8,
				"minimum_idle_connections":   0,
			},
		},
		"storage": map[string]any{
			"postgres": map[string]any{
				"address":  "tcp://" + *props.Database.Postgres.Host + ":" + fmt.Sprintf("%d", *props.Database.Postgres.Port),
				"database": props.Database.Postgres.Database,
				"schema":   props.Database.Postgres.Schema,
				"username": "{{ .POSTGRES_USERNAME }}",
				"timeout":  "5s",
				"password": "{{ .POSTGRES_PASSWORD }}",
			},
			"encryption_key": "{{ .STORAGE_ENCRYPTION_KEY }}",
		},
		"notifier": map[string]any{
			"disable_startup_check": false,
			"smtp": map[string]any{
				"address":  "smtp://{{ .SMTP_HOST }}:{{ .SMTP_PORT }}",
				"password": "{{ .SMTP_PASSWORD }}",
				"timeout":  "5s",
				"username": "{{ .SMTP_USERNAME }}",
				"sender": infrahelpers.UseOrDefault(
					props.SMTP.Sender,
					fmt.Sprintf("Authelia <authelia@%s>", props.SMTP.EmailDomain),
				),
				"identifier":            props.SMTP.EmailDomain,
				"subject":               infrahelpers.UseOrDefault(props.SMTP.Subject, "[authelia] {title}"),
				"startup_check_address": fmt.Sprintf("test@%s", props.SMTP.EmailDomain),
				"disable_html_emails":   false,
				"disable_require_tls":   false,
				"disable_starttls":      false,
				"tls": map[string]any{
					"server_name":     "{{ .SMTP_HOST }}",
					"skip_verify":     false,
					"minimum_version": "TLS1.2",
					"maximum_version": "TLS1.3",
				},
			},
		},
		"authentication_backend": map[string]any{
			"password_reset": map[string]any{
				"disable":    false,
				"custom_url": "",
			},
			// # How often authelia should check if there is an user update in LDAP
			// # refresh_interval: 1m
			// "refresh_interval": "always",
			// https://github.com/nitnelave/lldap/blob/main/example_configs/authelia_config.yml
		},
		"password_policy": map[string]any{
			"standard": map[string]any{
				"enabled":           false,
				"min_length":        8,
				"max_length":        0,
				"require_uppercase": true,
				"require_lowercase": true,
				"require_number":    true,
				"require_special":   true,
			},
			"zxcvbn": map[string]any{
				"enabled":   false,
				"min_score": 0,
			},
		},
		"identity_validation": map[string]any{
			"reset_password": map[string]any{
				"jwt_secret": "{{ .JWT_TOKEN }}",
			},
		},
	}
	secrets := map[string]string{
		"SMTP_PASSWORD":          "SMTP_PASSWORD",
		"SMTP_HOST":              "SMTP_HOST",
		"SMTP_PORT":              "SMTP_PORT",
		"SMTP_USERNAME":          "SMTP_USERNAME",
		"JWT_TOKEN":              "AUTHELIA_JWT_TOKEN",
		"SESSION_ENCRYPTION_KEY": "AUTHELIA_SESSION_ENCRYPTION_KEY",
		"STORAGE_ENCRYPTION_KEY": "AUTHELIA_STORAGE_ENCRYPTION_KEY",
		"POSTGRES_PASSWORD":      "POSTGRES_USER_PASSWORD",
		"POSTGRES_USERNAME":      "POSTGRES_USERNAME",
	}
	if props.OIDC.Enabled {
		for i, client := range props.OIDC.Clients {
			client.FillDefaults()
			props.OIDC.Clients[i] = client
		}
		secrets["OIDC_PRIVATE_KEY"] = "AUTHELIA_OIDC_PRIVATE_KEY"
		secrets["OIDC_HMAC_SECRET"] = "AUTHELIA_OIDC_HMAC_SECRET"
		if configMap["identity_providers"] == nil {
			configMap["identity_providers"] = map[string]any{}
		}
		configMap["identity_providers"].(map[string]any)["oidc"] = map[string]any{
			"enforce_pkce":                 "public_clients_only",
			"enable_pkce_plain_challenge":  false,
			"enable_client_debug_messages": false,
			"minimum_parameter_entropy":    8,
			"hmac_secret":                  "{{ .OIDC_HMAC_SECRET }}",
			"lifespans": map[string]string{
				"access_token":   "1h",
				"id_token":       "1h",
				"refresh_token":  "90m",
				"authorize_code": "1m",
			},
			"cors": map[string]any{
				"endpoints": []string{
					"token",
					"userinfo",
					// below ones may not be needed
					"authorization",
					"revocation",
					"introspection",
				},
				"allowed_origins_from_client_redirect_uris": true,
			},
			"clients": props.OIDC.Clients,
		}
		configMap["identity_providers"].(map[string]any)["oidc"].(map[string]any)["jwks"] = []map[string]any{
			{"key": "{{ .OIDC_PRIVATE_KEY }}", "certificate_chain": props.OIDC.IssuerCertificateChain},
		}
	}
	if props.Assets != nil {
		appProps.Containers[0].ExtraVolumeMounts = append(appProps.Containers[0].ExtraVolumeMounts, corev1.VolumeMount{Name: "assets", MountPath: "/assets"})
		appProps.ExtraVolumes = append(appProps.ExtraVolumes, corev1.Volume{Name: "assets", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}})
		appProps.InitContainers = []k8sapp.ApplicationContainer{{
			Name:  "init",
			Image: k8sapp.ImageInfo{Repository: "docker.io/curlimages/curl", Tag: "8.2.1"},
			Command: []string{
				"sh",
				"-c",
				fmt.Sprintf("sleep 1 && curl %q -o /assets/logo.png && curl %q -o /assets/favicon.ico", props.Assets.LogoURL, props.Assets.FaviconURL),
			},
			ExtraVolumeMounts: []corev1.VolumeMount{{Name: "assets", MountPath: "/assets"}},
		}}
		configMap["server"].(map[string]any)["asset_path"] = "/assets"
	}
	if props.AuthMode == "file" {
		// Note: this might be broken
		configMap["authentication_backend"].(map[string]any)["file"] = map[string]any{
			"path": "/config/users_database.yml",
		}
		appProps.ExternalSecrets = append(appProps.ExternalSecrets, k8sapp.ApplicationExternalSecret{
			Name:              "authelia-users-db",
			RemoteRefs:        map[string]string{"users_database.yml": "AUTHELIA_USERS_DATABASE_YML"},
			MountToContainers: []string{"authelia"},
			MountPath:         "/config/users_database.yml",
			SubPath:           "users_database.yml",
		})
	}
	if props.AuthMode == "ldap" {
		configMap["authentication_backend"].(map[string]any)["ldap"] = map[string]any{
			"implementation":      "custom",
			"address":             props.LDAP.URL,
			"password":            "{{ .LDAP_PASSWORD }}",
			"timeout":             "5s",
			"start_tls":           false,
			"base_dn":             props.LDAP.BaseDN,
			"additional_users_dn": "ou=people",
			// # users_filter: "(&({username_attribute}={input})(objectClass=person))"
			"users_filter":         props.LDAP.UsersFilter,
			"additional_groups_dn": "ou=groups",
			"groups_filter":        props.LDAP.GroupsFilter,
			"attributes": map[string]string{
				"username":     "uid",
				"group_name":   "cn",
				"display_name": props.LDAP.DisplayNameAttribute,
				"mail":         props.LDAP.MailAttribute,
			},
			"user":                             props.LDAP.User + "," + props.LDAP.BaseDN,
			"permit_referrals":                 false,
			"permit_unauthenticated_bind":      false,
			"permit_feature_detection_failure": false,
			"tls": map[string]any{
				"server_name":     "",
				"skip_verify":     false,
				"minimum_version": "TLS1.2",
				"maximum_version": "TLS1.3",
			},
		}
		secrets["LDAP_PASSWORD"] = props.LDAP.PasswordSecretKey
	}

	configMapString := infrahelpers.ToYamlString(configMap)
	configMapString = strings.Replace(configMapString, `"{{ .OIDC_PRIVATE_KEY }}"`, "|\n{{ .OIDC_PRIVATE_KEY | nindent 8 }}", 1)
	appProps.ExternalSecrets = append(appProps.ExternalSecrets, k8sapp.ApplicationExternalSecret{
		Name:       "authelia",
		RemoteRefs: secrets,
		Template: map[string]string{
			"configuration.yaml": configMapString,
		},
		MountToContainers: []string{"authelia"},
		MountName:         "config",
		MountPath:         "/configuration.yaml",
		SubPath:           "configuration.yaml",
		ReadOnly:          true,
	})

	k8sapp.NewApplication(scope, appProps)

	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "forwardauth-authelia"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			ForwardAuth: &traefikv1alpha1.ForwardAuth{
				Address: "http://authelia." + scope.Namespace() + ".svc.cluster.local/api/verify?rd=https://" + props.Ingress.SubDomain + "." + GetDomain(scope) + "/",
				AuthResponseHeaders: []string{
					"Remote-User",
					"Remote-Name",
					"Remote-Email",
					"Remote-Groups",
				},
				TrustForwardHeader: true,
			},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "headers-authelia"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Headers: &dynamic.Headers{
				BrowserXSSFilter:        true,
				CustomFrameOptionsValue: "SAMEORIGIN",
				CustomResponseHeaders: map[string]string{
					"Cache-Control": "no-store",
					"Pragma":        "no-cache",
				},
			},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "chain-authelia-auth"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Chain: &traefikv1alpha1.Chain{
				Middlewares: []traefikv1alpha1.MiddlewareRef{{Name: "forwardauth-authelia", Namespace: scope.Namespace()}},
			},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "chain-authelia"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Chain: &traefikv1alpha1.Chain{
				Middlewares: []traefikv1alpha1.MiddlewareRef{{Name: "headers-authelia", Namespace: scope.Namespace()}},
			},
		},
	})
	scope.AddApiObject(&traefikv1alpha1.TLSOption{
		ObjectMeta: metav1.ObjectMeta{
			Name: "authelia",
			// Annotations: infrahelpers.MergeAnnotations( // is this needed?
			// 	GetCertIssuerAnnotation(scope),
			// 	map[string]string{
			// 		"traefik.ingress.kubernetes.io/router.middlewares": scope.Namespace() + "-chain-authelia@kubernetescrd",
			// 	},
			// ),
		},
		Spec: traefikv1alpha1.TLSOptionSpec{
			CipherSuites: []string{"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384"},
			MaxVersion:   "VersionTLS13",
			MinVersion:   "VersionTLS12",
		},
	})
}

// selector:
//   matchLabels:
//   app.kubernetes.io/instance: authelia
// revisionHistoryLimit: 5
// strategy:
//   rollingUpdate:
//     maxSurge: 25%
//     maxUnavailable: 25%
//   type: RollingUpdate
