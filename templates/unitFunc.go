package templates

func GetUnitFuncTemplate() string {
	return `{{- define "UnitFunc" }}{{- range . }}{{- if .UnitFuncId.Valid }}{{- if campaignAndRaceCheck .Race.String .Campaign.String }}[{{ .UnitFuncId.String }}]
{{- if .Art.Valid }}
Art={{ .Art.String }}
{{- end }}
{{- if .Specialart.Valid }}
Specialart={{ .Specialart.String }}
{{- end }}
{{- if .Casterupgradeart.Valid }}
Casterupgradeart={{ .Casterupgradeart.String }}
{{- end }}
{{- if .Targetart.Valid }}
Targetart={{ .Targetart.String }}
{{- end }}
{{- if .Scorescreenicon.Valid }}
ScoreScreenIcon={{ .Scorescreenicon.String }}
{{- end }}
{{- if .Buttonpos.Valid }}
Buttonpos={{ .Buttonpos.String }}
{{- end }}
{{- if .Missileart.Valid }}
Missileart={{ .Missileart.String }}
{{- end }}
{{- if .Missilearc.Valid }}
Missilearc={{ .Missilearc.String }}
{{- end }}
{{- if .Missilespeed.Valid }}
Missilespeed={{ .Missilespeed.String }}
{{- end }}
{{- if .Missilehoming.Valid }}
MissileHoming={{ .Missilehoming.String }}
{{- end }}
{{- if .Trains.Valid }}
Trains={{ .Trains.String }}
{{- end }}
{{- if .Builds.Valid }}
Builds={{ .Builds.String }}
{{- end }}
{{- if .Researches.Valid }}
Builds={{ .Researches.String }}
{{- end }}
{{- if .Upgrade.Valid }}
Upgrade={{ .Upgrade.String }}
{{- end }}
{{- if .Sellunits.Valid }}
Sellunits={{ .Sellunits.String }}
{{- end }}
{{- if .Requires.Valid }}
Requires={{ .Requires.String }}
{{- end }}
{{- if .Requires1.Valid }}
Requires1={{ .Requires1.String }}
{{- end }}
{{- if .Requires2.Valid }}
Requires2={{ .Requires2.String }}
{{- end }}
{{- if .Requirescount.Valid }}
Requirescount={{ .Requirescount.String }}
{{- end }}
{{- if .Buildingsoundlabel.Valid }}
BuildingSoundLabel={{ .Buildingsoundlabel.String }}
{{- end }}
{{- if .Movementsoundlabel.Valid }}
MovementSoundLabel={{ .Movementsoundlabel.String }}
{{- end }}
{{- if .Loopingsoundfadein.Valid }}
LoopingSoundFadeIn={{ .Loopingsoundfadein.String }}
{{- end }}
{{- if .Loopingsoundfadeout.Valid }}
LoopingSoundFadeOut={{ .Loopingsoundfadeout.String }}
{{- end }}
{{- if .Revive.Valid }}
Revive={{ .Revive.String }}
{{- end }}
{{- if .Animprops.Valid }}
Animprops={{ .Animprops.String }}
{{- end }}
{{- if .Attachmentanimprops.Valid }}
Attachmentanimprops={{ .Attachmentanimprops.String }}
{{- end }}
{{- if .Dependencyor.Valid }}
DependencyOr={{ .Dependencyor.String }}
{{- end }}
{{- if .Makeitems.Valid }}
Makeitems={{ .Makeitems.String }}
{{- end }}
{{- if .Attachmentlinkprops.Valid }}
Attachmentlinkprops={{ .Attachmentlinkprops.String }}
{{- end }}
{{- if .Boneprops.Valid }}
Boneprops={{ .Boneprops.String }}
{{- end }}
{{- if .Sellitems.Valid }}
Sellitems={{ .Sellitems.String }}
{{- end }}
{{- if .Randomsoundlabel.Valid }}
Randomsoundlabel={{ .Randomsoundlabel.String }}
{{- end }}
{{ end }}
{{- end }}
{{- end }}
{{- end }}`
}
