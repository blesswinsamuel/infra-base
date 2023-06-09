package middlewares_traefikio


// Headers holds the headers middleware configuration.
//
// This middleware manages the requests and responses headers. More info: https://doc.traefik.io/traefik/v2.10/middlewares/http/headers/#customrequestheaders
type MiddlewareSpecHeaders struct {
	// AccessControlAllowCredentials defines whether the request can include user credentials.
	AccessControlAllowCredentials *bool `field:"optional" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// AccessControlAllowHeaders defines the Access-Control-Request-Headers values sent in preflight response.
	AccessControlAllowHeaders *[]*string `field:"optional" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// AccessControlAllowMethods defines the Access-Control-Request-Method values sent in preflight response.
	AccessControlAllowMethods *[]*string `field:"optional" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// AccessControlAllowOriginList is a list of allowable origins.
	//
	// Can also be a wildcard origin "*".
	AccessControlAllowOriginList *[]*string `field:"optional" json:"accessControlAllowOriginList" yaml:"accessControlAllowOriginList"`
	// AccessControlAllowOriginListRegex is a list of allowable origins written following the Regular Expression syntax (https://golang.org/pkg/regexp/).
	AccessControlAllowOriginListRegex *[]*string `field:"optional" json:"accessControlAllowOriginListRegex" yaml:"accessControlAllowOriginListRegex"`
	// AccessControlExposeHeaders defines the Access-Control-Expose-Headers values sent in preflight response.
	AccessControlExposeHeaders *[]*string `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// AccessControlMaxAge defines the time that a preflight request may be cached.
	AccessControlMaxAge *float64 `field:"optional" json:"accessControlMaxAge" yaml:"accessControlMaxAge"`
	// AddVaryHeader defines whether the Vary header is automatically added/updated when the AccessControlAllowOriginList is set.
	AddVaryHeader *bool `field:"optional" json:"addVaryHeader" yaml:"addVaryHeader"`
	// AllowedHosts defines the fully qualified list of allowed domain names.
	AllowedHosts *[]*string `field:"optional" json:"allowedHosts" yaml:"allowedHosts"`
	// BrowserXSSFilter defines whether to add the X-XSS-Protection header with the value 1;
	//
	// mode=block.
	BrowserXssFilter *bool `field:"optional" json:"browserXssFilter" yaml:"browserXssFilter"`
	// ContentSecurityPolicy defines the Content-Security-Policy header value.
	ContentSecurityPolicy *string `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// ContentTypeNosniff defines whether to add the X-Content-Type-Options header with the nosniff value.
	ContentTypeNosniff *bool `field:"optional" json:"contentTypeNosniff" yaml:"contentTypeNosniff"`
	// CustomBrowserXSSValue defines the X-XSS-Protection header value.
	//
	// This overrides the BrowserXssFilter option.
	CustomBrowserXssValue *string `field:"optional" json:"customBrowserXssValue" yaml:"customBrowserXssValue"`
	// CustomFrameOptionsValue defines the X-Frame-Options header value.
	//
	// This overrides the FrameDeny option.
	CustomFrameOptionsValue *string `field:"optional" json:"customFrameOptionsValue" yaml:"customFrameOptionsValue"`
	// CustomRequestHeaders defines the header names and values to apply to the request.
	CustomRequestHeaders *map[string]*string `field:"optional" json:"customRequestHeaders" yaml:"customRequestHeaders"`
	// CustomResponseHeaders defines the header names and values to apply to the response.
	CustomResponseHeaders *map[string]*string `field:"optional" json:"customResponseHeaders" yaml:"customResponseHeaders"`
	// Deprecated: use PermissionsPolicy instead.
	FeaturePolicy *string `field:"optional" json:"featurePolicy" yaml:"featurePolicy"`
	// ForceSTSHeader defines whether to add the STS header even when the connection is HTTP.
	ForceStsHeader *bool `field:"optional" json:"forceStsHeader" yaml:"forceStsHeader"`
	// FrameDeny defines whether to add the X-Frame-Options header with the DENY value.
	FrameDeny *bool `field:"optional" json:"frameDeny" yaml:"frameDeny"`
	// HostsProxyHeaders defines the header keys that may hold a proxied hostname value for the request.
	HostsProxyHeaders *[]*string `field:"optional" json:"hostsProxyHeaders" yaml:"hostsProxyHeaders"`
	// IsDevelopment defines whether to mitigate the unwanted effects of the AllowedHosts, SSL, and STS options when developing.
	//
	// Usually testing takes place using HTTP, not HTTPS, and on localhost, not your production domain. If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false.
	IsDevelopment *bool `field:"optional" json:"isDevelopment" yaml:"isDevelopment"`
	// PermissionsPolicy defines the Permissions-Policy header value.
	//
	// This allows sites to control browser features.
	PermissionsPolicy *string `field:"optional" json:"permissionsPolicy" yaml:"permissionsPolicy"`
	// PublicKey is the public key that implements HPKP to prevent MITM attacks with forged certificates.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// ReferrerPolicy defines the Referrer-Policy header value.
	//
	// This allows sites to control whether browsers forward the Referer header to other sites.
	ReferrerPolicy *string `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// Deprecated: use RedirectRegex instead.
	SslForceHost *bool `field:"optional" json:"sslForceHost" yaml:"sslForceHost"`
	// Deprecated: use RedirectRegex instead.
	SslHost *string `field:"optional" json:"sslHost" yaml:"sslHost"`
	// SSLProxyHeaders defines the header keys with associated values that would indicate a valid HTTPS request.
	//
	// It can be useful when using other proxies (example: "X-Forwarded-Proto": "https").
	SslProxyHeaders *map[string]*string `field:"optional" json:"sslProxyHeaders" yaml:"sslProxyHeaders"`
	// Deprecated: use EntryPoint redirection or RedirectScheme instead.
	SslRedirect *bool `field:"optional" json:"sslRedirect" yaml:"sslRedirect"`
	// Deprecated: use EntryPoint redirection or RedirectScheme instead.
	SslTemporaryRedirect *bool `field:"optional" json:"sslTemporaryRedirect" yaml:"sslTemporaryRedirect"`
	// STSIncludeSubdomains defines whether the includeSubDomains directive is appended to the Strict-Transport-Security header.
	StsIncludeSubdomains *bool `field:"optional" json:"stsIncludeSubdomains" yaml:"stsIncludeSubdomains"`
	// STSPreload defines whether the preload flag is appended to the Strict-Transport-Security header.
	StsPreload *bool `field:"optional" json:"stsPreload" yaml:"stsPreload"`
	// STSSeconds defines the max-age of the Strict-Transport-Security header.
	//
	// If set to 0, the header is not set.
	StsSeconds *float64 `field:"optional" json:"stsSeconds" yaml:"stsSeconds"`
}

