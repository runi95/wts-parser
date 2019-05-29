package templates

func GetUnitAbilitiesTemplate() string {
	return `{{- define "UnitAbilities" }}ID;PWXL;N;E
B;X7;Y{{ .RowCount }};D0
C;X1;Y1;K"unitAbilID"
C;X2;K"sortAbil"
C;X3;K"comment(s)"
C;X4;K"auto"
C;X5;K"abilList"
C;X6;K"heroAbilList"
C;X7;K"InBeta"
{{- range $index, $element := .Unit }}{{- if .UnitAbilID.Valid }}
C;X1;Y{{ inc $index }};K{{ .UnitAbilID.String }} 
{{- if .Auto.Valid }}
C;X4;K{{ .Auto.String }}
{{- end }} 
{{- if .AbilList.Valid }}
C;X5;K{{ .AbilList.String }}
{{- end }} 
{{- if .HeroAbilList.Valid }}
C;X6;K{{ .HeroAbilList.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}