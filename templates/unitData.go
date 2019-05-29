package templates

func GetUnitDataTemplate() string {
	return `{{- define "UnitData" }}ID;PWXL;N;E
B;X32;Y{{ .RowCount }};D0
C;X1;Y1;K"unitID"
C;X2;K"sort"
C;X3;K"comment(s)"
C;X4;K"race"
C;X5;K"prio"
C;X6;K"threat"
C;X7;K"valid"
C;X8;K"deathType"
C;X9;K"death"
C;X10;K"canSleep"
C;X11;K"cargoSize"
C;X12;K"movetp"
C;X13;K"moveHeight"
C;X14;K"moveFloor"
C;X15;K"turnRate"
C;X16;K"propWin"
C;X17;K"orientInterp"
C;X18;K"formation"
C;X19;K"targType"
C;X20;K"pathTex"
C;X21;K"fatLOS"
C;X22;K"points"
C;X23;K"buffType"
C;X24;K"buffRadius"
C;X25;K"nameCount"
C;X26;K"canFlee"
C;X27;K"requireWaterRadius"
C;X28;K"isBuildOn"
C;X29;K"canBuildOn"
C;X30;K"InBeta"
C;X31;K"version"
{{- range $index, $element := .Unit }}{{- if .UnitID.Valid }}
C;X1;Y{{ inc $index }};K{{ .UnitID.String }} 
{{- if .Race.Valid }}
C;X4;K{{ .Race.String }}
{{- end }} 
{{- if .Prio.Valid }}
C;X5;K{{ .Prio.String }}
{{- end }} 
{{- if .Threat.Valid }}
C;X6;K{{ .Threat.String }}
{{- end }} 
{{- if .Valid.Valid }}
C;X7;K{{ .Valid.String }}
{{- end }} 
{{- if .DeathType.Valid }}
C;X8;K{{ .DeathType.String }}
{{- end }} 
{{- if .Death.Valid }}
C;X9;K{{ .Death.String }}
{{- end }} 
{{- if .CanSleep.Valid }}
C;X10;K{{ .CanSleep.String }}
{{- end }} 
{{- if .CargoSize.Valid }}
C;X11;K{{ .CargoSize.String }}
{{- end }} 
{{- if .Movetp.Valid }}
C;X12;K{{ .Movetp.String }}
{{- end }} 
{{- if .MoveHeight.Valid }}
C;X13;K{{ .MoveHeight.String }}
{{- end }} 
{{- if .MoveFloor.Valid }}
C;X14;K{{ .MoveFloor.String }}
{{- end }} 
{{- if .TurnRate.Valid }}
C;X15;K{{ .TurnRate.String }}
{{- end }} 
{{- if .PropWin.Valid }}
C;X16;K{{ .PropWin.String }}
{{- end }} 
{{- if .OrientInterp.Valid }}
C;X17;K{{ .OrientInterp.String }}
{{- end }} 
{{- if .Formation.Valid }}
C;X18;K{{ .Formation.String }}
{{- end }} 
{{- if .TargType.Valid }}
C;X19;K{{ .TargType.String }}
{{- end }} 
{{- if .PathTex.Valid }}
C;X20;K{{ .PathTex.String }}
{{- end }} 
{{- if .FatLOS.Valid }}
C;X21;K{{ .FatLOS.String }}
{{- end }} 
{{- if .Points.Valid }}
C;X22;K{{ .Points.String }}
{{- end }} 
{{- if .BuffType.Valid }}
C;X23;K{{ .BuffType.String }}
{{- end }} 
{{- if .BuffRadius.Valid }}
C;X24;K{{ .BuffRadius.String }}
{{- end }} 
{{- if .NameCount.Valid }}
C;X25;K{{ .NameCount.String }}
{{- end }} 
{{- if .CanFlee.Valid }}
C;X26;K{{ .CanFlee.String }}
{{- end }} 
{{- if .RequireWaterRadius.Valid }}
C;X27;K{{ .RequireWaterRadius.String }}
{{- end }} 
{{- if .IsBuildOn.Valid }}
C;X28;K{{ .IsBuildOn.String }}
{{- end }} 
{{- if .CanBuildOn.Valid }}
C;X29;K{{ .CanBuildOn.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
