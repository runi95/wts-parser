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
{{- range $index, $element := .UnitData }}
C;X1;Y{{ inc $index }};K{{ .UnitID.String }}
C;X4;K{{ .Race.String }}
C;X5;K{{ .Prio.String }}
C;X6;K{{ .Threat.String }}
C;X7;K{{ .Valid.String }}
C;X8;K{{ .DeathType.String }}
C;X9;K{{ .Death.String }}
C;X10;K{{ .CanSleep.String }}
C;X11;K{{ .CargoSize.String }}
C;X12;K{{ .Movetp.String }}
C;X13;K{{ .MoveHeight.String }}
C;X14;K{{ .MoveFloor.String }}
C;X15;K{{ .TurnRate.String }}
C;X16;K{{ .PropWin.String }}
C;X17;K{{ .OrientInterp.String }}
C;X18;K{{ .Formation.String }}
C;X19;K{{ .TargType.String }}
C;X20;K{{ .PathTex.String }}
C;X21;K{{ .FatLOS.String }}
C;X22;K{{ .Points.String }}
C;X23;K{{ .BuffType.String }}
C;X24;K{{ .BuffRadius.String }}
C;X25;K{{ .NameCount.String }}
C;X26;K{{ .CanFlee.String }}
C;X27;K{{ .RequireWaterRadius.String }}
C;X28;K{{ .IsBuildOn.String }}
C;X29;K{{ .CanBuildOn.String }}
{{- end }}
E
{{ end }}`
}
