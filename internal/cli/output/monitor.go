package output

var MonitorListTemplate = `Your Checkmate Monitors on instance {{green .InstanceURL}}
Team ID: {{cyan .TeamID }}

{{if and .TotalMonitors (gt (len .TotalMonitors) 0) -}}
Total Monitors: {{blue (intToString (len .TotalMonitors))}}
{{else}}
There are no monitors
{{- end }}
{{if and .UpMonitors (gt (len .UpMonitors) 0)}}
Up Monitors: {{green (intToString (len .UpMonitors))}}
{{- else}}
{{- end }}`
