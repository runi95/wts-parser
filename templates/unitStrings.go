package templates

func GetUnitStringsTemplate() string {
	return `{{- define "UnitStrings" }}{{- range . }}{{- if .UnitStringId.Valid }}{{- if campaignAndRaceCheck .Race.String .Campaign.String }}[{{ .UnitStringId.String }}]
{{- if .UnitString.Name.Valid }}
Name={{ .UnitString.Name.String }}
{{- end }}
{{- if .Hotkey.Valid }}
Hotkey={{ .Hotkey.String }}
{{- end }}
{{- if .Description.Valid }}
Description={{ .Description.String }}
{{- end }}
{{- if .Tip.Valid }}
Tip={{ .Tip.String }}
{{- end }}
{{- if .Ubertip.Valid }}
Ubertip={{ .Ubertip.String }}
{{- end }}
{{- if .Editorsuffix.Valid }}
EditorSuffix={{ .Editorsuffix.String }}
{{- end }}
{{- if .Propernames.Valid }}
Propernames={{ .Propernames.String }}
{{- end }}
{{- if .Revivetip.Valid }}
Revivetip={{ .Revivetip.String }}
{{- end }}
{{- if .Awakentip.Valid }}
Awakentip={{ .Awakentip.String }}
{{- end }}
{{- if .Casterupgradename.Valid }}
Casterupgradename={{ .Casterupgradename.String }}
{{- end }}
{{- if .Casterupgradetip.Valid }}
Casterupgradetip={{ .Casterupgradetip.String }}
{{- end }}
{{- if .Dependencyor.Valid }}
DependencyOr={{ .Dependencyor.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}
{{- end }}`
}
