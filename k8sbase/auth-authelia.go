package k8sbase

import (
	"fmt"
	"log"

	"github.com/blesswinsamuel/infra-base/infrahelpers"
	"github.com/blesswinsamuel/infra-base/kubegogen"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	traefikv1alpha1 "github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/blesswinsamuel/infra-base/k8sapp"
)

type AutheliaClient struct {
	AuthorizationPolicy          string   `json:"authorization_policy"`
	ConsentMode                  string   `json:"consent_mode"`
	Description                  string   `json:"description"`
	GrantTypes                   []string `json:"grant_types"`
	ID                           string   `json:"id"`
	PreConfiguredConsentDuration string   `json:"pre_configured_consent_duration"`
	Public                       bool     `json:"public"`
	RedirectURIs                 []string `json:"redirect_uris"`
	ResponseTypes                []string `json:"response_types"`
	Scopes                       []string `json:"scopes"`
	Secret                       string   `json:"secret,omitempty"`
	SectorIdentifier             string   `json:"sector_identifier"`
	UserInfoSigningAlgorithm     string   `json:"userinfo_signing_algorithm"`
}

func (c *AutheliaClient) FillDefaults() {
	if c.ConsentMode == "" {
		c.ConsentMode = "auto"
	}
	if c.GrantTypes == nil {
		c.GrantTypes = []string{"refresh_token", "authorization_code"}
	}
	if c.ResponseTypes == nil {
		c.ResponseTypes = []string{"code"}
	}
	if c.UserInfoSigningAlgorithm == "" {
		c.UserInfoSigningAlgorithm = "none"
	}
	if c.Scopes == nil {
		c.Scopes = []string{"openid", "profile", "email", "groups"}
	}
}

type AutheliaProps struct {
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
	Assets *struct {
		LogoURL    string `json:"logoURL"`
		FaviconURL string `json:"faviconURL"`
	} `json:"assets"`
	RedirectionSubDomain string `json:"redirectionSubDomain"`
}

// https://github.com/authelia/chartrepo/tree/master/charts/authelia

func (props *AutheliaProps) Chart(scope kubegogen.Construct) kubegogen.Construct {
	appProps := &k8sapp.ApplicationProps{
		Name:               "authelia",
		IngressMiddlewares: []k8sapp.NameNamespace{{Name: "chain-authelia", Namespace: k8sapp.GetNamespaceContext(scope)}},
		Containers: []k8sapp.ApplicationContainer{{
			Name:  "authelia",
			Image: props.ImageInfo,
			Ports: []k8sapp.ContainerPort{
				{Name: "http", Port: 9091, ServicePort: 80, Ingress: &k8sapp.ApplicationIngress{Host: "authelia." + GetDomain(scope)}},
				{Name: "metrics", Port: 9959, PrometheusScrape: &k8sapp.ApplicationPrometheusScrape{Path: "/metrics"}},
			},
			Args: []string{
				"--config=/configuration.yaml",
			},
			Command: []string{"authelia"},
			Env: map[string]string{
				"AUTHELIA_SERVER_DISABLE_HEALTHCHECK":                      "true",
				"AUTHELIA_JWT_SECRET_FILE":                                 "/secrets/JWT_TOKEN",
				"AUTHELIA_SESSION_SECRET_FILE":                             "/secrets/SESSION_ENCRYPTION_KEY",
				"AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD_FILE":       "/secrets/LDAP_PASSWORD",
				"AUTHELIA_NOTIFIER_SMTP_PASSWORD_FILE":                     "/secrets/SMTP_PASSWORD",
				"AUTHELIA_STORAGE_ENCRYPTION_KEY_FILE":                     "/secrets/STORAGE_ENCRYPTION_KEY",
				"AUTHELIA_STORAGE_POSTGRES_PASSWORD_FILE":                  "/secrets/STORAGE_PASSWORD",
				"AUTHELIA_IDENTITY_PROVIDERS_OIDC_HMAC_SECRET_FILE":        "/secrets/OIDC_HMAC_SECRET",
				"AUTHELIA_IDENTITY_PROVIDERS_OIDC_ISSUER_PRIVATE_KEY_FILE": "/secrets/OIDC_PRIVATE_KEY",
				"TZ": "UTC",
			},
			ExtraVolumeMounts: []corev1.VolumeMount{
				{Name: "secrets", MountPath: "/secrets", ReadOnly: true},
			},
			LivenessProbe:  &corev1.Probe{FailureThreshold: 5, PeriodSeconds: 30, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
			ReadinessProbe: &corev1.Probe{FailureThreshold: 5, PeriodSeconds: 5, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
			StartupProbe:   &corev1.Probe{FailureThreshold: 6, InitialDelaySeconds: 10, PeriodSeconds: 5, SuccessThreshold: 1, TimeoutSeconds: 5, ProbeHandler: corev1.ProbeHandler{HTTPGet: &corev1.HTTPGetAction{Path: "/api/health", Port: intstr.FromString("http"), Scheme: "HTTP"}}},
		}},
		ExtraVolumes: []corev1.Volume{
			{Name: "secrets", VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: "authelia", Items: []corev1.KeyToPath{
				{Key: "JWT_TOKEN", Path: "JWT_TOKEN"},
				{Key: "SESSION_ENCRYPTION_KEY", Path: "SESSION_ENCRYPTION_KEY"},
				{Key: "STORAGE_ENCRYPTION_KEY", Path: "STORAGE_ENCRYPTION_KEY"},
				{Key: "STORAGE_PASSWORD", Path: "STORAGE_PASSWORD"},
				{Key: "LDAP_PASSWORD", Path: "LDAP_PASSWORD"},
				{Key: "SMTP_PASSWORD", Path: "SMTP_PASSWORD"},
				{Key: "OIDC_PRIVATE_KEY", Path: "OIDC_PRIVATE_KEY"},
				{Key: "OIDC_HMAC_SECRET", Path: "OIDC_HMAC_SECRET"},
			}}}},
		},
		ExternalSecrets:    []k8sapp.ApplicationExternalSecret{},
		EnableServiceLinks: infrahelpers.Ptr(false),
	}
	// https://github.com/authelia/chartrepo/blob/master/charts/authelia/templates/configMap.yaml
	configMap := map[string]any{
		"theme": "light",
		"server": map[string]any{
			"host":           "0.0.0.0",
			"port":           9091,
			"asset_path":     "",
			"headers":        map[string]any{"csp_template": ""},
			"buffers":        map[string]any{"read": 4096, "write": 4096},
			"timeouts":       map[string]any{"read": "6s", "write": "6s", "idle": "30s"},
			"enable_pprof":   false,
			"enable_expvars": false,
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
			"issuer":      "home.bless.win",
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
			"address":               "time.cloudflare.com:123",
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
		"default_redirection_url": "https://" +
			infrahelpers.Ternary(props.RedirectionSubDomain != "", props.RedirectionSubDomain+".", "") +
			GetDomain(scope),
		"default_2fa_method": "",
		"access_control":     props.AccessControl,
		"session": map[string]any{
			"name":                 "authelia_session",
			"domain":               GetDomain(scope),
			"same_site":            "lax",
			"expiration":           "1h",
			"inactivity":           "5m",
			"remember_me_duration": "1M",
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
				"host":     props.Database.Postgres.Host,
				"port":     props.Database.Postgres.Port,
				"database": props.Database.Postgres.Database,
				"schema":   props.Database.Postgres.Schema,
				"username": props.Database.Postgres.Username,
				"timeout":  "5s",
			},
		},
		"notifier": map[string]any{
			"disable_startup_check": false,
			"smtp": map[string]any{
				"host":     props.SMTP.Host,
				"port":     props.SMTP.Port,
				"timeout":  "5s",
				"username": props.SMTP.Username,
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
					"server_name":     props.SMTP.Host,
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
		"identity_providers": map[string]any{},
	}
	if props.OIDC.Enabled {
		for i, client := range props.OIDC.Clients {
			client.FillDefaults()
			props.OIDC.Clients[i] = client
		}
		configMap["identity_providers"].(map[string]any)["oidc"] = map[string]any{
			"access_token_lifespan":        "1h",
			"authorize_code_lifespan":      "1m",
			"id_token_lifespan":            "1h",
			"refresh_token_lifespan":       "90m",
			"enforce_pkce":                 "public_clients_only",
			"enable_pkce_plain_challenge":  false,
			"enable_client_debug_messages": false,
			"minimum_parameter_entropy":    8,
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
		if props.OIDC.IssuerCertificateChain != "" {
			configMap["identity_providers"].(map[string]any)["oidc"].(map[string]any)["issuer_certificate_chain"] = props.OIDC.IssuerCertificateChain
		}
	}
	if props.Assets != nil {
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
			"url":                 props.LDAP.URL,
			"timeout":             "5s",
			"start_tls":           false,
			"base_dn":             props.LDAP.BaseDN,
			"username_attribute":  "uid",
			"additional_users_dn": "ou=people",
			// # users_filter: "(&({username_attribute}={input})(objectClass=person))"
			"users_filter":                     props.LDAP.UsersFilter,
			"additional_groups_dn":             "ou=groups",
			"groups_filter":                    props.LDAP.GroupsFilter,
			"group_name_attribute":             "cn",
			"mail_attribute":                   props.LDAP.MailAttribute,
			"display_name_attribute":           props.LDAP.DisplayNameAttribute,
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
	}

	appProps.ConfigMaps = append(appProps.ConfigMaps, k8sapp.ApplicationConfigMap{
		Name:              "authelia",
		Data:              map[string]string{"configuration.yaml": infrahelpers.ToYamlString(configMap)},
		MountToContainers: []string{"authelia"},
		MountName:         "config",
		MountPath:         "/configuration.yaml",
		SubPath:           "configuration.yaml",
		ReadOnly:          true,
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
	appProps.ExternalSecrets = append(appProps.ExternalSecrets, k8sapp.ApplicationExternalSecret{
		Name:       "authelia",
		RemoteRefs: secrets,
	})

	app := k8sapp.NewApplicationChart(scope, "authelia", appProps)
	// k8sapp.NewHelm(chart, "helm", &k8sapp.HelmProps{
	// 	ChartInfo:     props.ChartInfo,
	// 	ReleaseName:   "authelia",
	// 	Namespace:     chart.Namespace(),
	// 	PatchResource: patchResource,
	// 	Values: map[string]any{
	// 		"domain": GetDomain(scope),
	// 		"pod":    pod,
	// 		"ingress": map[string]any{
	// 			"enabled":   true,
	// 			"subdomain": props.Ingress.SubDomain,
	// 			"traefikCRD": map[string]any{
	// 				"enabled":             true,
	// 				"disableIngressRoute": true,
	// 			},
	// 			"annotations": infrahelpers.MergeAnnotations(
	// 				GetCertIssuerAnnotation(scope),
	// 				map[string]string{
	// 					"traefik.ingress.kubernetes.io/router.middlewares": k8sapp.GetNamespaceContext(scope) + "-chain-authelia@kubernetescrd",
	// 				},
	// 			),
	// 		},
	// 		"secret": map[string]any{
	// 			"existingSecret": "authelia",
	// 		},
	// 		"configMap": configMap,
	// 	},
	// })

	k8sapp.NewK8sObject(app, "forwardauth-authelia", changeApiVersion(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "forwardauth-authelia"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			ForwardAuth: &traefikv1alpha1.ForwardAuth{
				Address: "http://authelia." + k8sapp.GetNamespaceContext(app) + ".svc.cluster.local/api/verify?rd=https://" + props.Ingress.SubDomain + "." + GetDomain(scope) + "/",
				AuthResponseHeaders: []string{
					"Remote-User",
					"Remote-Name",
					"Remote-Email",
					"Remote-Groups",
				},
				TrustForwardHeader: true,
			},
		},
	}))

	k8sapp.NewK8sObject(app, "headers-authelia", changeApiVersion(&traefikv1alpha1.Middleware{
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
	}))

	k8sapp.NewK8sObject(app, "chain-authelia-auth", changeApiVersion(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "chain-authelia-auth"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Chain: &traefikv1alpha1.Chain{
				Middlewares: []traefikv1alpha1.MiddlewareRef{{Name: "forwardauth-authelia", Namespace: k8sapp.GetNamespaceContext(app)}},
			},
		},
	}))

	k8sapp.NewK8sObject(app, "chain-authelia", changeApiVersion(&traefikv1alpha1.Middleware{
		ObjectMeta: metav1.ObjectMeta{Name: "chain-authelia"},
		Spec: traefikv1alpha1.MiddlewareSpec{
			Chain: &traefikv1alpha1.Chain{
				Middlewares: []traefikv1alpha1.MiddlewareRef{{Name: "headers-authelia", Namespace: k8sapp.GetNamespaceContext(app)}},
			},
		},
	}))

	k8sapp.NewK8sObject(app, "authelia-tlsoption", changeApiVersion(&traefikv1alpha1.TLSOption{
		ObjectMeta: metav1.ObjectMeta{
			Name: "authelia",
			Annotations: infrahelpers.MergeAnnotations( // is this needed?
				GetCertIssuerAnnotation(scope),
				map[string]string{
					"traefik.ingress.kubernetes.io/router.middlewares": k8sapp.GetNamespaceContext(scope) + "-chain-authelia@kubernetescrd",
				},
			),
		},
		Spec: traefikv1alpha1.TLSOptionSpec{
			CipherSuites: []string{"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384"},
			MaxVersion:   "VersionTLS13",
			MinVersion:   "VersionTLS12",
		},
	}))

	return app
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

// to check diff
func changeApiVersion(v runtime.Object) runtime.Object {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(v)
	if err != nil {
		log.Fatalf("ToUnstructured: %v", err)
	}
	unstructuredObj["apiVersion"] = "traefik.containo.us/v1alpha1"
	switch v.(type) {
	case *traefikv1alpha1.Middleware:
		unstructuredObj["kind"] = "Middleware"
	case *traefikv1alpha1.TLSOption:
		unstructuredObj["kind"] = "TLSOption"
	}
	return &unstructured.Unstructured{Object: unstructuredObj}
}
