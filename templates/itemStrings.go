package templates

func GetItemStringsTemplate() string {
	return `{{- define "ItemStrings" }}{{- range . }}{{- if .ItemStringId.Valid }}[{{ .ItemStringId.String }}]
{{- if .Name.Valid }}
Name={{ .Name.String }}
{{- end }}
{{- if .Editorsuffix.Valid }}
Editorsuffix={{ .Editorsuffix.String }}
{{- end }}
{{- if .Tip.Valid }}
Tip={{ .Tip.String }}
{{- end }}
{{- if .Ubertip.Valid }}
Ubertip={{ .Ubertip.String }}
{{- end }}
{{- if .Hotkey.Valid }}
Hotkey={{ .Hotkey.String }}
{{- end }}
{{- if .Description.Valid }}
Description={{ .Description.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}`
}
