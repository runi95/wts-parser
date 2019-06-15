package templates

func GetAbilityFuncTemplate() string {
	return `{{- define "AbilityFunc" }}{{- range . }}{{- if .AbilityFuncId.Valid }}{{- if raceCheck .Race.String }}[{{ .AbilityFuncId.String }}]
{{- if .Art.Valid }}
Art={{ .Art.String }}
{{- end }}
{{- if .Unart.Valid }}
Unart={{ .Unart.String }}
{{- end }}
{{- if .Researchart.Valid }}
Researchart={{ .Researchart.String }}
{{- end }}
{{- if .Buttonpos.Valid }}
Buttonpos={{ .Buttonpos.String }}
{{- end }}
{{- if .Buffart.Valid }}
Buffart={{ .Buffart.String }}
{{- end }}
{{- if .Unbuttonpos.Valid }}
Unbuttonpos={{ .Unbuttonpos.String }}
{{- end }}
{{- if .Researchbuttonpos.Valid }}
Researchbuttonpos={{ .Researchbuttonpos.String }}
{{- end }}
{{- if .Specialart.Valid }}
Specialart={{ .Specialart.String }}
{{- end }}
{{- if .Effectsound.Valid }}
Effectsound={{ .Effectsound.String }}
{{- end }}
{{- if .Effectsoundlooped.Valid }}
Effectsoundlooped={{ .Effectsoundlooped.String }}
{{- end }}
{{- if .Effectart.Valid }}
Effectart={{ .Effectart.String }}
{{- end }}
{{- if .Casterart.Valid }}
Casterart={{ .Casterart.String }}
{{- end }}
{{- if .Targetart.Valid }}
Targetart={{ .Targetart.String }}
{{- end }}
{{- if .Missilearc.Valid }}
Missilearc={{ .Missilearc.String }}
{{- end }}
{{- if .Missileart.Valid }}
Missileart={{ .Missileart.String }}
{{- end }}
{{- if .Missilespeed.Valid }}
Missilespeed={{ .Missilespeed.String }}
{{- end }}
{{- if .Missilehoming.Valid }}
Missilehoming={{ .Missilehoming.String }}
{{- end }}
{{- if .Requires.Valid }}
Requires={{ .Requires.String }}
{{- end }}
{{- if .Requiresamount.Valid }}
Requiresamount={{ .Requiresamount.String }}
{{- end }}
{{- if .Order.Valid }}
Order={{ .Order.String }}
{{- end }}
{{- if .Unorder.Valid }}
Unorder={{ .Unorder.String }}
{{- end }}
{{- if .Orderon.Valid }}
Orderon={{ .Orderon.String }}
{{- end }}
{{- if .Orderoff.Valid }}
Orderoff={{ .Orderoff.String }}
{{- end }}
{{- if .Animnames.Valid }}
Animnames={{ .Animnames.String }}
{{- end }}
{{- if .Targetattach.Valid }}
Targetattach={{ .Targetattach.String }}
{{- end }}
{{- if .Casterattach.Valid }}
Casterattach={{ .Casterattach.String }}
{{- end }}
{{- if .Lightningeffect.Valid }}
Lightningeffect={{ .Lightningeffect.String }}
{{- end }}
{{- if .Spelldetail.Valid }}
Spelldetail={{ .Spelldetail.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}
{{- end }}`
}
