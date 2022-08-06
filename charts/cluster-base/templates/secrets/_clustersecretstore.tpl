# https://external-secrets.io/v0.5.8/provider-kubernetes/
# https://external-secrets.io/v0.5.8/spec/
{{- define "cluster-base.secrets.clustersecretstore" }}
---
apiVersion: external-secrets.io/v1beta1
kind: ClusterSecretStore
metadata:
  name: secretstore
spec:
  provider:
    kubernetes:
      remoteNamespace: default
      server:
        url: "kubernetes.default"
        caProvider:
          type: ConfigMap
          name: kube-root-ca.crt
          key: ca.crt
          namespace: default # only ClusterSecretStore
      auth:
        # serviceAccount:
        #   name: "external-secrets-store-sa"
        #   namespace: "default" # only ClusterSecretStore
        token:
          bearerToken:
            name: external-secrets-store-sa-token
            key: token
            namespace: default
        # cert:
        #   clientCert:
        #     name: "k3s-serving"
        #     key: "tls.crt"
        #     namespace: "kube-system" # only ClusterSecretStore
        #   clientKey:
        #     name: "k3s-serving"
        #     key: "tls.key"
        #     namespace: "kube-system" # only ClusterSecretStore
---
apiVersion: v1
kind: Secret
metadata:
  name: external-secrets-store-sa-token
  annotations:
    kubernetes.io/service-account.name: external-secrets-store-sa
type: kubernetes.io/service-account-token
---
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: default
  name: external-secrets-store-sa
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: external-secrets-store-role
rules:
- apiGroups: [""]
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - authorization.k8s.io
  resources:
  - selfsubjectrulesreviews
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: default
  name: external-secrets-store-role-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: external-secrets-store-role
subjects:
- kind: ServiceAccount
  name: external-secrets-store-sa
  namespace: default
{{- end }}
