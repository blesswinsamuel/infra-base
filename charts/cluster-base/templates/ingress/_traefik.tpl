# https://github.com/traefik/traefik-helm-chart/blob/master/traefik
# https://github.com/k3s-io/k3s/blob/master/manifests/traefik.yaml
{{ define "cluster-base.ingress.traefik" }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChartConfig
metadata:
  name: traefik
  namespace: kube-system
spec:
  valuesContent: |-
    image:
      name: traefik
      tag: v2.9.8
    ingressClass:
      enabled: true
      isDefaultClass: true
    ingressRoute:
      dashboard:
        enabled: false
    providers:
      kubernetesCRD:
        allowCrossNamespace: true
    additionalArguments:
      - --accesslog=true
      - --accesslog.format=json
      - --log.format=json
    ports:
      web:
        redirectTo: websecure
        forwardedHeaders:
          trustedIPs:
            {{- $.Values.traefik.trustedIPs | toYaml | nindent 12 }}
        proxyProtocol:
          trustedIPs:
            {{- $.Values.traefik.trustedIPs | toYaml | nindent 12 }}
      websecure:
        forwardedHeaders:
          trustedIPs:
            {{- $.Values.traefik.trustedIPs | toYaml | nindent 12 }}
        proxyProtocol:
          trustedIPs:
            {{- $.Values.traefik.trustedIPs | toYaml | nindent 12 }}
        ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
        # works only from traefik v3
        # asDefault: true
      websecure:
        middlewares:
          - {{ include "cluster-base.ingress.router-middleware.auth" $ }}
    service:
      spec:
        externalTrafficPolicy: Local  # So that traefik gets the real IP - https://github.com/k3s-io/k3s/discussions/2997#discussioncomment-413904
---
# # https://github.com/traefik/traefik/issues/5571#issuecomment-539393453 - affects wss in goatcounter
# apiVersion: traefik.containo.us/v1alpha1
# kind: Middleware
# metadata:
#   name: ssl-header
#   namespace: kube-system
# spec:
#   headers:
#     customRequestHeaders:
#       X-Forwarded-Proto: https,wss
# ---
{{- with .Values.traefik.dashboardIngress }}
{{- if .enabled }}
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: traefik-dashboard
  namespace: kube-system
spec:
  dnsNames:
    - "{{ .subDomain }}.{{ tpl $.Values.global.domain $ }}"
  secretName: traefik-dashboard-tls
  issuerRef:
    name: "{{ $.Values.global.clusterCertIssuerName }}"
    kind: ClusterIssuer
---
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: traefik-dashboard-external
  namespace: kube-system
spec:
  entryPoints:
    - websecure
  routes:
    - match: Host(`{{ .subDomain }}.{{ tpl $.Values.global.domain $ }}`) && (PathPrefix(`/dashboard`) || PathPrefix(`/api`))
      kind: Rule
      services:
        - name: api@internal
          kind: TraefikService
      middlewares:
        {{- if eq $.Values.global.internetAuthType "basic-auth" }}
        - name: traefik-basic-auth
          namespace: kube-system
        {{- else if eq $.Values.global.internetAuthType "traefik-forward-auth" }}
        - name: traefik-forward-auth
          namespace: auth
        {{- else if eq $.Values.global.internetAuthType "authelia" }}
        - name: forwardauth-authelia
          namespace: auth
        {{- end }}
  tls:
    secretName: traefik-dashboard-tls
    domains:
      - main: "{{ .subDomain }}.{{ tpl $.Values.global.domain $ }}"
{{- end }}
{{- end }}
---
{{- if .Values.traefik.middlewares.basicAuth.enabled }}
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: traefik-basic-auth
  namespace: kube-system
spec:
  basicAuth:
    secret: traefik-basic-auth
    removeHeader: true
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: traefik-basic-auth
  namespace: kube-system
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: secretstore
    kind: ClusterSecretStore
  target:
    name: traefik-basic-auth
  data:
  - secretKey: users
    remoteRef:
      key: TRAEFIK_BASIC_AUTH_USERS
{{- end }}
---
{{- if .Values.traefik.middlewares.stripPrefix.enabled }}
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: traefik-strip-prefix
  namespace: kube-system
spec:
  stripPrefixRegex:
    regex:
    - ^/[^/]+
{{- end }}
{{- end }}
