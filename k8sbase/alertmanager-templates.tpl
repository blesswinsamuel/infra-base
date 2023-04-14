{{- define "slack.title" -}}
    [{{ .Status | toUpper }}{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ end }}] {{ .CommonLabels.alertname }} for {{ .CommonLabels.job }}
    {{- if gt (len .CommonLabels) (len .GroupLabels) -}}
      {{" "}}(
      {{- with .CommonLabels.Remove .GroupLabels.Names }}
        {{- range $index, $label := .SortedPairs -}}
          {{ if $index }}, {{ end }}
          {{- $label.Name }}="{{ $label.Value -}}"
        {{- end }}
      {{- end -}}
      )
    {{- end }}
{{- end -}}

{{- define "slack.text" -}}
    {{ range .Alerts -}}
    *Alert:* {{ .Annotations.summary }}{{ if .Labels.severity }} - `{{ .Labels.severity }}`{{ end }}

    *Description:* {{ .Annotations.description }}

    *Graph:* <{{ .GeneratorURL }}|:chart_with_upwards_trend:>

    *Details:*
      {{ range .Labels.SortedPairs }} • *{{ .Name }}:* `{{ .Value }}`
      {{ end }}
    {{ end }}
{{- end -}}

{{/* https://core.telegram.org/bots/update56kabdkb12ibuisabdubodbasbdaosd#html-style */}}
{{/* https://github.com/prometheus/alertmanager/blob/main/docs/notifications.md */}}
{{- define "telegram.message" -}}
  {{- if eq .Status "firing" }}
  <b>Status</b>: <b>{{.Status | toUpper}} 🔥</b> <code>{{ .CommonLabels.alertname }}</code> for job <code>{{ .CommonLabels.job }}</code>
  {{- end }}
  {{- if eq .Status "resolved" }}
  <b>Status</b>: <b>{{.Status | toUpper}} ✅</b> <code>{{ .CommonLabels.alertname }}</code> for job <code>{{ .CommonLabels.job }}</code>
  {{- end }}
  {{- range .Alerts }}
  <b>Alert</b>: {{ .Annotations.summary | reReplaceAll "&" "&amp;" | reReplaceAll "<" "&lt;" | reReplaceAll ">" "&gt;" }}
  <b>Description</b>: {{ .Annotations.description | reReplaceAll "&" "&amp;" | reReplaceAll "<" "&lt;" | reReplaceAll ">" "&gt;" }}
  <b>Details</b>:
  {{ range .Labels.SortedPairs }}   - <b>{{ .Name }}</b>: <code>{{ .Value }}</code>
  {{ end -}}
  <b>Graph</b>: <a href="{{ .GeneratorURL | reReplaceAll "\"" "%22" }}">📈 Grafana</a>
  <b>Severity</b>: {{ .Labels.severity }}{{ if eq .Labels.severity "warning" }} ⚠️{{ else if eq .Labels.severity "critical" }} 🚨{{ end }}
  {{- end }}
{{- end -}}
