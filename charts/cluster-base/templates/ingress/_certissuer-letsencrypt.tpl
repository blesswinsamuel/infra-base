{{ define "cluster-base.ingress.certissuer" }}
{{- with .Values.certIssuer }}
---
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
  {{- with .namespace }}
  namespace: {{ tpl . $ }}
  {{- end }}
spec:
  acme:
    email: {{ .email }}
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
      {{- if eq .solver "http" }}
      - http01:
          ingress:
            class: traefik
      {{- end }}
      {{- if eq .solver "dns" }}
      - dns01:
          cloudflare:
            email: {{ .email }}
            apiTokenSecretRef:
              name: cloudflare-api-token
              key: api-token
      {{- end }}
---
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-staging
  {{- with .namespace }}
  namespace: {{ tpl . $ }}
  {{- end }}
spec:
  acme:
    email: {{ .email }}
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-staging
    solvers:
      {{- if eq .solver "http" }}
      - http01:
          ingress:
            class: traefik
      {{- end }}
      {{- if eq .solver "dns" }}
      - dns01:
          cloudflare:
            email: {{ .email }}
            apiTokenSecretRef:
              name: cloudflare-api-token
              key: api-token
      {{- end }}
---
{{- if eq .solver "dns" }}
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: cloudflare-api-token
  # https://cert-manager.io/docs/faq/cluster-resource/
  namespace: cert-manager
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  target:
    name: cloudflare-api-token
  data:
  - secretKey: api-token
    remoteRef:
      key: CLOUDFLARE_API_TOKEN
{{- end }}
{{- end }}
{{- end }}
