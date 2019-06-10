package templates

func GetAbilityMetaDataTemplate() string {
	return `{{- define "AbilityMetaData" }}ID;PWXL;N;E
B;X26;Y{{ .RowCount }};D0
C;X1;Y1;K"ID"
C;X2;K"field"
C;X3;K"slk"
C;X4;K"index"
C;X5;K"repeat"
C;X6;K"data"
C;X7;K"category"
C;X8;K"displayName"
C;X9;K"sort"
C;X10;K"type"
C;X11;K"changeFlags"
C;X12;K"importType"
C;X13;K"stringExt"
C;X14;K"caseSens"
C;X15;K"canBeEmpty"
C;X16;K"minVal"
C;X17;K"maxVal"
C;X18;K"forceNonNeg"
C;X19;K"useUnit"
C;X20;K"useHero"
C;X21;K"useItem"
C;X22;K"useCreep"
C;X23;K"useSpecific"
C;X24;K"notSpecific"
C;X25;K"version"
C;X26;K"section"
{{- range $index, $element := .Data }}{{- if .ID.Valid }}
C;X1;Y{{ inc $index }};K{{ .ID.String }} 
{{- if .Field.Valid }}
C;X2;K{{ .Field.String }}
{{- end }}
{{- if .Slk.Valid }}
C;X3;K{{ .Slk.String }}
{{- end }}
{{- if .Index.Valid }}
C;X4;K{{ .Index.String }}
{{- end }}
{{- if .Repeat.Valid }}
C;X5;K{{ .Repeat.String }}
{{- end }}
{{- if .Data.Valid }}
C;X6;K{{ .Data.String }}
{{- end }}
{{- if .Category.Valid }}
C;X7;K{{ .Category.String }}
{{- end }}
{{- if .DisplayName.Valid }}
C;X8;K{{ .DisplayName.String }}
{{- end }}
{{- if .Sort.Valid }}
C;X9;K{{ .Sort.String }}
{{- end }}
{{- if .Type.Valid }}
C;X10;K{{ .Type.String }}
{{- end }}
{{- if .ChangeFlags.Valid }}
C;X11;K{{ .ChangeFlags.String }}
{{- end }}
{{- if .ImportType.Valid }}
C;X12;K{{ .ImportType.String }}
{{- end }}
{{- if .StringExt.Valid }}
C;X13;K{{ .StringExt.String }}
{{- end }}
{{- if .CaseSens.Valid }}
C;X14;K{{ .CaseSens.String }}
{{- end }}
{{- if .CanBeEmpty.Valid }}
C;X15;K{{ .CanBeEmpty.String }}
{{- end }}
{{- if .MinVal.Valid }}
C;X16;K{{ .MinVal.String }}
{{- end }}
{{- if .MaxVal.Valid }}
C;X17;K{{ .MaxVal.String }}
{{- end }}
{{- if .ForceNonNeg.Valid }}
C;X18;K{{ .ForceNonNeg.String }}
{{- end }}
{{- if .UseUnit.Valid }}
C;X19;K{{ .UseUnit.String }}
{{- end }}
{{- if .UseHero.Valid }}
C;X20;K{{ .UseHero.String }}
{{- end }}
{{- if .UseItem.Valid }}
C;X21;K{{ .UseItem.String }}
{{- end }}
{{- if .UseCreep.Valid }}
C;X22;K{{ .UseCreep.String }}
{{- end }}
{{- if .UseSpecific.Valid }}
C;X23;K{{ .UseSpecific.String }}
{{- end }}
{{- if .NotSpecific.Valid }}
C;X24;K{{ .NotSpecific.String }}
{{- end }}
{{- if .Version.Valid }}
C;X25;K{{ .Version.String }}
{{- end }}
{{- if .Section.Valid }}
C;X26;K{{ .Section.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
