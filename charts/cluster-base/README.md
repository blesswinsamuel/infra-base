# cluster-base

## Debugging

```
helm template cloudlab /Volumes/BleSSD/Projects/blesswinsamuel/bless-stack/cloudlab \
  | yq e 'select(.metadata.name == "traefik-forward-auth" and .kind == "HelmChart")' \
  | yq e '.spec.valuesContent' \
  | helm template --namespace auth --repo https://k8s-at-home.com/charts/ --version 2.1.2 traefik-forward-auth traefik-forward-auth --values - --debug
```
