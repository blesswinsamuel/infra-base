{{- define "slack.title" -}}
{{- $emoji := "" }}
{{- if eq .Status "firing" }}
{{- $emoji = ":fire:" }}
{{- end }}
{{- if eq .Status "resolved" }}
{{- $emoji = ":white_check_mark:" }}
{{- end }}
{{ $emoji }} [{{ .Status | toUpper }}{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ end }}] {{ .CommonLabels.alertname }} for {{ .CommonLabels.job }}
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
{{- range .Alerts }}
*Alert:* {{ .Annotations.summary }}{{ if .Labels.severity }} - `{{ .Labels.severity }}`{{ end }}
*Description:* {{ .Annotations.description }}
*Graph:* <{{ .GeneratorURL }}|:chart_with_upwards_trend:>
*Details:*
  {{- range .Labels.SortedPairs }}
  • *{{ .Name }}:* `{{ .Value }}`
  {{- end }}
{{ end }}
{{- end -}}

{{/* https://core.telegram.org/bots/update56kabdkb12ibuisabdubodbasbdaosd#html-style */}}
{{/* https://github.com/prometheus/alertmanager/blob/main/docs/notifications.md */}}
{{/* https://github.com/prometheus/alertmanager/blob/a85979e19d24490322d5ce342301d17b0f13dcc5/template/template.go#L170-L195 */}}
{{- define "telegram.message" -}}
{{- $emoji := "" }}
{{- if eq .Status "firing" }}
{{- $emoji = "🔥" }}
{{- end }}
{{- if eq .Status "resolved" }}
{{- $emoji = "✅" }}
{{- end }}
<b>Status</b>: <b>{{.Status | toUpper}} {{$emoji}}</b> <code>{{ .CommonLabels.alertname }}</code> for job <code>{{ .CommonLabels.job }}</code>
{{- range .Alerts }}
  <b>Alert</b>: {{ .Annotations.summary }}
  <b>Description</b>: {{ .Annotations.description }}
  <b>Details</b>:
  {{- range .Labels.SortedPairs }}
    - <b>{{ .Name }}</b>: <code>{{ .Value }}</code>
  {{- end }}
  <b>Graph</b>: <a href="{{ .GeneratorURL }}">📈 Grafana</a>
  <b>Severity</b>: {{ .Labels.severity }}{{ if eq .Labels.severity "warning" }} ⚠️{{ else if eq .Labels.severity "critical" }} 🚨{{ end }}
{{- end }}
{{- end -}}
