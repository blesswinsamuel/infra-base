{{- define "cluster-base.system.helm-ops" -}}
{{- with .Values.system.helmOps }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: helm-ops
  namespace: '{{ tpl .namespace $ }}'
data:
  hooks.yaml: |
    - id: redeploy-webhook
      execute-command: "/config/redeploy.sh"
      command-working-directory: "/config"
  redeploy.sh: |
    #!/bin/bash
    set -ex

    HELM_CHART_DIR=$GIT_REPO_PATH/{{ .helmChartPath }}

    helm -n {{ .helmReleaseNamespace }} dependency build $HELM_CHART_DIR
    helm -n {{ .helmReleaseNamespace }} diff upgrade {{ .helmReleaseName }} $HELM_CHART_DIR --three-way-merge
    helm -n {{ .helmReleaseNamespace }} upgrade --install {{ .helmReleaseName }} $HELM_CHART_DIR
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: helm-ops
  namespace: '{{ tpl .namespace $ }}'
  annotations:
    reloader.stakater.com/search: "true"
spec:
  selector:
    matchLabels:
      app: helm-ops
  replicas: 1
  template:
    metadata:
      labels:
        app: helm-ops
    spec:
      serviceAccountName: helm-ops
      volumes:
        - name: helm-chart-git
          emptyDir: {}
        - name: config
          configMap:
            defaultMode: 0755
            name: helm-ops
        - name: git-ssh-key
          secret:
            secretName: helm-ops
            defaultMode: 0600
      containers:
        - name: webhook
          image: blesswinsamuel/docker-helm:main
          ports:
            - containerPort: 80
          command:
            - webhook 
            - -hooks=/config/hooks.yaml
            - -verbose
            - -port=80
            - -hotreload
          env:
            {{- range $key, $val := .env }}
            - name: {{ $key }}
              value: {{ $val | quote }}
            {{- end}}
            {{- range $key, $val := .envFromSecret }}
            - name: {{ $key }}
              valueFrom:
                secretKeyRef:
                  name: {{ $val.name }}
                  key: {{ $val.key }}
            {{- end }}
            - name: GIT_REPO_PATH
              value: /repo/helm-chart-git/current
          volumeMounts:
            - name: helm-chart-git
              mountPath: /repo
              # readOnly: true
            - name: config
              mountPath: /config
              readOnly: true
        - name: git-sync
          image: blesswinsamuel/git-sync:main
          # imagePullPolicy: Always
          # command: ['sleep', '1000000']
          args:
            - --repo={{ tpl .gitRepo.url $ }}
            - --branch={{ .gitRepo.branch }}
            - --ssh=true
            - --ssh-key-file=/ssh-key/git-private-key
            - --ssh-known-hosts=false
            - --ssh-known-hosts-file=/ssh-key/known_hosts
            - --depth=1
            # - --max-sync-failures=5
            # - --wait=60  # --period=60s
            - --dest=current
            - --root=/repo/helm-chart-git
            - --webhook-url=http://localhost/hooks/redeploy-webhook
            - --webhook-timeout=120s
          volumeMounts:
            - name: helm-chart-git
              mountPath: /repo
            - name: git-ssh-key
              mountPath: /ssh-key
# ---
# apiVersion: v1
# kind: Service
# metadata:
#   name: helm-ops
#   namespace: '{{ tpl .namespace $ }}'
#   labels:
#     app: helm-ops
# spec:
#   ports:
#     - port: 80
#       protocol: TCP
#   selector:
#     app: helm-ops
---
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: helm-ops
  namespace: '{{ tpl .namespace $ }}'
  annotations:
    reloader.stakater.com/match: "true"
spec:
  refreshInterval: 2m
  secretStoreRef:
    name: secretstore
    kind: ClusterSecretStore
  target:
    name: helm-ops
  data:
  - secretKey: git-private-key
    remoteRef: { key: doppler-secrets, property: BLESS_STACK_DEPLOY_KEY }
  - secretKey: known_hosts
    remoteRef: { key: doppler-secrets, property: GITHUB_KNOWN_HOSTS }
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: helm-ops
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: helm-ops
  namespace: '{{ tpl .namespace $ }}'
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: helm-ops
  namespace: '{{ tpl .namespace $ }}'
{{- end }}
{{- end }}
