# https://github.com/k8s-at-home/charts/tree/master/charts/stable/traefik-forward-auth
# https://github.com/k8s-at-home/library-charts/tree/main/charts/stable/common
# https://github.com/thomseddon/traefik-forward-auth
{{ define "cluster-base.auth.traefik-forward-auth" }}
{{- with .Values.traefikForwardAuth }}
{{- if .enabled }}
---
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: traefik-forward-auth
  namespace: '{{ include "cluster-base.namespace.helm-chart" $ }}'
spec:
  repo: https://k8s-at-home.com/charts/
  chart: traefik-forward-auth
  version: "2.2.2"
  targetNamespace: {{ tpl .namespace $ }}
  valuesContent: |-
    image:
      tag: 2.2.0{{ .imageTagSuffix }}
    controller:
      annotations:
        reloader.stakater.com/search: "true"
    ingress:
      main:
        enabled: true
        annotations:
          {{ include "cluster-base.ingress.annotation.cert-issuer" $ }}
          traefik.ingress.kubernetes.io/router.middlewares: auth-traefik-forward-auth@kubernetescrd
        hosts:
        - host: {{ .ingress.subDomain }}.{{ $.Values.global.domain }}
          paths:
          - path: /
        tls:
        - secretName: traefik-forward-auth-tls
          hosts:
            - {{ .ingress.subDomain }}.{{ $.Values.global.domain }}
    args:
      {{- .args | toYaml | nindent 6 }}
    env:
      WHITELIST: {{ .whilelist }}
      LOG_FORMAT: json
      LOG_LEVEL: info
      AUTH_HOST: {{ .ingress.subDomain }}.{{ $.Values.global.domain }}
      COOKIE_DOMAIN: {{ $.Values.global.domain }}
      PROVIDERS_GOOGLE_CLIENT_ID:
        valueFrom:
          secretKeyRef:
            name: traefik-forward-auth
            key: PROVIDERS_GOOGLE_CLIENT_ID
      PROVIDERS_GOOGLE_CLIENT_SECRET:
        valueFrom:
          secretKeyRef:
            name: traefik-forward-auth
            key: PROVIDERS_GOOGLE_CLIENT_SECRET
      # PROVIDERS_GOOGLE_PROMPT: 

      SECRET:
        valueFrom:
          secretKeyRef:
            name: traefik-forward-auth
            key: SECRET
    middleware:
      nameOverride: traefik-forward-auth
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: traefik-forward-auth
  namespace: {{ tpl .namespace $ }}
  annotations:
    reloader.stakater.com/match: "true"
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: '{{ tpl $.Values.global.clusterExternalSecretStoreName $ }}'
    kind: ClusterSecretStore
  data:
  - secretKey: PROVIDERS_GOOGLE_CLIENT_SECRET
    remoteRef:
      key: AUTH_GOOGLE_CLIENT_SECRET
  - secretKey: PROVIDERS_GOOGLE_CLIENT_ID
    remoteRef:
      key: AUTH_GOOGLE_CLIENT_ID
  - secretKey: SECRET
    remoteRef:
      key: AUTH_COOKIE_SECRET
{{- end }}
{{- end }}
{{ end }}
