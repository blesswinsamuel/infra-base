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
{{/* https://prometheus.io/docs/alerting/latest/notifications/ */}}
{{/* https://github.com/prometheus/alertmanager/blob/a85979e19d24490322d5ce342301d17b0f13dcc5/template/template.go#L170-L195 */}}
{{/* https://github.com/prometheus/alertmanager/blob/ca5089d33eabaf03638a083d9a84f08c6de1acfb/template/default.tmpl#L115-L124 */}}
{{/* https://gist.github.com/jidckii/5ac5f8f20368b56de72af70222509b7b */}}
{{ define "__alertmanagerURL" }}{{ .ExternalURL }}/#/alerts?receiver={{ .Receiver | urlquery }}{{ end }}
{{ define "__severityEmoji" }}{{ if eq . "warning" }}⚠️{{ else if eq . "critical" }}🚨{{ else }}{{ . }}{{ end }}{{ end }}
{{ define "__alertStatusEmoji" }}{{ if eq . "firing" }}🔥{{ else if eq . "resolved" }}✅{{ else }}🪪{{ end }}{{ end }}

{{- define "telegram.message.alert.list" -}}
{{- range $i, $alert := . }}
---
{{- if lt $i 1 }}
{{ template "__alertStatusEmoji" $alert.Status }} <b>{{ $alert.Labels.alertname }}</b> {{ template "__severityEmoji" $alert.Labels.severity }}
{{- if $alert.Annotations.summary }}
📝 {{ $alert.Annotations.summary }}
{{- end }}
{{- if $alert.Annotations.description }}
📖 {{ $alert.Annotations.description }}
{{- end }}
{{- if $alert.Annotations.runbook_url }}
📚 <a href="{{ $alert.Annotations.runbook_url }}">Runbook</a>
{{- end }}
🏷 Labels:
{{- range $alert.Labels.SortedPairs }}
{{- if .Value }}
  <i>{{ .Name }}</i>: <code>{{ .Value }}</code>
{{- end }}
{{- end }}
{{- if $alert.GeneratorURL }}
📈 <a href="{{ $alert.GeneratorURL }}">Graph</a> 📈
{{- end }}
{{- else }}
{{ template "__alertStatusEmoji" $alert.Status }} {{ $alert.Labels.alertname }} {{ template "__severityEmoji" $alert.Labels.severity }}
{{- if $alert.Annotations.summary }}
📝 {{ $alert.Annotations.summary }}
{{- end }}
...
{{- end }}
{{- end }}
{{- end -}}

{{- define "telegram.message" -}}
<b>{{ if eq .Status "firing" }}🔥{{ else if eq .Status "resolved" }}✅{{ end }} {{.Status | toUpper}}</b> ({{ .Alerts | len }})
{{- template "telegram.message.alert.list" .Alerts }}
---
💊 <a href="{{ template "__alertmanagerURL" . }}">Alertmanager</a> 💊
{{- end -}}
