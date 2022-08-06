{{/* Merge the local chart values and the common chart defaults */}}
{{- define "cluster-base.values.setup" -}}
  {{- if hasKey .Values "cluster-base" -}}
    {{- $defaultValues := deepCopy (get .Values "cluster-base") -}}
    {{- $userValues := deepCopy (omit .Values "cluster-base") -}}
    {{- $mergedValues := mustMergeOverwrite $defaultValues $userValues -}}
    {{- $_ := set . "Values" (deepCopy $mergedValues) -}}
  {{- end -}}
{{- end -}}
