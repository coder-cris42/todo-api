{{- /* Generate a fullname using release and chart name, keep under 63 chars */ -}}
{{- define "todo-api.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
