package global

{{- if .HasGlobal }}

import "github.com/KeSilent/study-hub/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}