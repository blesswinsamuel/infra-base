# # https://github.com/external-secrets/external-secrets/tree/main/deploy/charts/external-secrets
# apiVersion: helm.cattle.io/v1
# kind: HelmChartConfig
# metadata:
#   name: external-secrets
#   namespace: external-secrets
# spec:
#   valuesContent: |-
#     webhook:
#       create: false
#     certController:
#       create: false
