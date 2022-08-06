{{- define "cluster-base.secrets.dockercreds" -}}
{{- with .Values.externalSecrets.regcred }}
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: regcred
  namespace: '{{ tpl .namespace $ }}'
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: secretstore
    kind: ClusterSecretStore
  target:
    template:
      type: kubernetes.io/dockerconfigjson
      data:
        .dockerconfigjson: |
          {
            "auths": {
              "{{`{{ .registry }}`}}": {
                "auth": {{`"{{ (printf "%s:%s" .username .password) | b64enc }}"`}}
              }
            }
          }
    name: regcred
    creationPolicy: Owner
  data:
  - secretKey: registry
    remoteRef:
      key: doppler-secrets
      property: {{ .keyPrefix }}CONTAINER_REGISTRY_URL
  - secretKey: username
    remoteRef:
      key: doppler-secrets
      property: {{ .keyPrefix }}CONTAINER_REGISTRY_USERNAME
  - secretKey: password
    remoteRef:
      key: doppler-secrets
      property: {{ .keyPrefix }}CONTAINER_REGISTRY_PASSWORD
{{- end -}}
{{- end -}}
