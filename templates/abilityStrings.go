package templates

func GetAbilityStringsTemplate() string {
	return `{{- define "AbilityStrings" }}{{- range . }}{{- if .AbilityStringId.Valid }}{{- if raceCheck .Race.String }}[{{ .AbilityStringId.String }}]
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
{{- if .Untip.Valid }}
Untip={{ .Untip.String }}
{{- end }}
{{- if .Unubertip.Valid }}
Unubertip={{ .Unubertip.String }}
{{- end }}
{{- if .Bufftip.Valid }}
Bufftip={{ .Bufftip.String }}
{{- end }}
{{- if .Buffubertip.Valid }}
Buffubertip={{ .Buffubertip.String }}
{{- end }}
{{- if .Hotkey.Valid }}
Hotkey={{ .Hotkey.String }}
{{- end }}
{{- if .Unhotkey.Valid }}
Unhotkey={{ .Unhotkey.String }}
{{- end }}
{{- if .Researchhotkey.Valid }}
Researchhotkey={{ .Researchhotkey.String }}
{{- end }}
{{- if .Researchtip.Valid }}
Researchtip={{ .Researchtip.String }}
{{- end }}
{{- if .Researchubertip.Valid }}
Researchubertip={{ .Researchubertip.String }}
{{- end }}
{{- if .Globalmessage.Valid }}
Globalmessage={{ .Globalmessage.String }}
{{- end }}
{{- if .Globalsound.Valid }}
Globalsound={{ .Globalsound.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}
{{- end }}`
}
