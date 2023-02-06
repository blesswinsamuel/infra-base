# https://external-secrets.io/v0.5.8/provider-kubernetes/
# https://external-secrets.io/v0.5.8/spec/
{{- define "cluster-base.secrets.clustersecretstore" }}
---
apiVersion: v1
kind: Secret
metadata:
  name: doppler-token-auth-api
  namespace: default
data:
  dopplerToken: '{{ .Values.global.dopplerServiceToken }}'
---
apiVersion: external-secrets.io/v1beta1
kind: ClusterSecretStore
metadata:
  name: '{{ .Values.global.clusterExternalSecretStoreName }}'
spec:
  # controller: doppler  # like ingressClassName definition
  provider:
    doppler:
      auth:
        secretRef:
          dopplerToken:
            name: doppler-token-auth-api
            key: dopplerToken
            namespace: default
{{- end }}
