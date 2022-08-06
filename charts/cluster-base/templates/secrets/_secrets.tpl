{{ define "cluster-base.secrets" }}
  {{- include "cluster-base.secrets.clustersecretstore" . }}
  {{- include "cluster-base.secrets.dopplersecret" . }}
  {{- include "cluster-base.secrets.dockercreds" . }}
{{ end }}
