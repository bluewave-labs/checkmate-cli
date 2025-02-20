package output

var MonitorListTemplate = `Your Checkmate Monitors on instance {{green .InstanceURL}}
Team ID: {{cyan .TeamID }}

{{if and .TotalMonitors (gt (len .TotalMonitors) 0) -}}
Total Monitors: {{blue (intToString (len .TotalMonitors))}}
{{else}}
There are no monitors
{{- end -}}

{{if and .UpMonitors (gt (len .UpMonitors) 0)}}
Up Monitors: {{green (intToString (len .UpMonitors))}}
{{range .UpMonitors}}
- {{ .Name -}} ({{ .URL }})
{{- end}}
{{- else}}
{{- end }}

{{if and .DownMonitors (gt (len .DownMonitors) 0)}}
Down Monitors: {{red (intToString (len .DownMonitors))}}
{{range .DownMonitors}}
- {{ .Name -}} ({{ .URL }})
{{- end}}
{{- else}}
{{- end }}
`
