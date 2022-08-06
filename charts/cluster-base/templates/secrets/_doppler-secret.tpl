{{- define "cluster-base.secrets.dopplersecret" }}
---
apiVersion: secrets.doppler.com/v1alpha1
kind: DopplerSecret
metadata:
  name: doppler-secrets # DopplerSecret Name
  namespace: doppler-operator-system
spec:
  tokenSecret: # Kubernetes service token secret (namespace defaults to doppler-operator-system)
    name: doppler-service-token
  managedSecret: # Kubernetes managed secret (will be created if does not exist)
    name: doppler-secrets
    namespace: default # Should match the namespace of deployments that will use the secret
---
apiVersion: v1
kind: Secret
metadata:
  name: doppler-service-token
  namespace: doppler-operator-system
data:
  serviceToken: {{ .Values.global.dopplerServiceToken }}
{{- end }}
