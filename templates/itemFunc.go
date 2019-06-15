package templates

func GetItemFuncTemplate() string {
	return `{{- define "ItemFunc" }}{{- range . }}{{- if .ItemFuncId.Valid }}[{{ .ItemFuncId.String }}]
{{- if .Art.Valid }}
Art={{ .Art.String }}
{{- end }}
{{- if .Dependencyor.Valid }}
Dependencyor={{ .Dependencyor.String }}
{{- end }}
{{- if .Xpfactor.Valid }}
Xpfactor={{ .Xpfactor.String }}
{{- end }}
{{- if .Buttonpos.Valid }}
Buttonpos={{ .Buttonpos.String }}
{{- end }}
{{- if .Requires.Valid }}
Requires={{ .Requires.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}`
}
