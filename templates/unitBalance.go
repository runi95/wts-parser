package templates

func GetUnitBalanceTemplate() string {
	return `{{- define "UnitBalance" }}ID;PWXL;N;E
B;X61;Y{{ .RowCount }};D0
C;X1;Y1;K"unitBalanceID"
C;X2;K"sortBalance"
C;X3;K"sort2"
C;X4;K"comment(s)"
C;X5;K"level"
C;X7;K"type"
C;X8;K"goldcost"
C;X9;K"lumbercost"
C;X10;K"goldRep"
C;X11;K"lumberRep"
C;X12;K"fmade"
C;X13;K"fused"
C;X14;K"bountydice"
C;X15;K"bountysides"
C;X16;K"bountyplus"
C;X17;K"lumberbountydice"
C;X18;K"lumberbountysides"
C;X19;K"lumberbountyplus"
C;X20;K"stockMax"
C;X21;K"stockRegen"
C;X22;K"stockStart"
C;X23;K"HP"
C;X24;K"realHP"
C;X25;K"regenHP"
C;X26;K"regenType"
C;X27;K"manaN"
C;X28;K"realM"
C;X29;K"mana0"
C;X30;K"regenMana"
C;X31;K"def"
C;X32;K"defUp"
C;X33;K"realdef"
C;X34;K"defType"
C;X35;K"spd"
C;X36;K"minSpd"
C;X37;K"maxSpd"
C;X38;K"bldtm"
C;X39;K"reptm"
C;X40;K"sight"
C;X41;K"nsight"
C;X42;K"STR"
C;X43;K"INT"
C;X44;K"AGI"
C;X45;K"STRplus"
C;X46;K"INTplus"
C;X47;K"AGIplus"
C;X48;K"abilTest"
C;X49;K"Primary"
C;X50;K"upgrades"
C;X51;K"tilesets"
C;X52;K"nbrandom"
C;X53;K"isbldg"
C;X54;K"preventPlace"
C;X55;K"requirePlace"
C;X56;K"repulse"
C;X57;K"repulseParam"
C;X58;K"repulseGroup"
C;X59;K"repulsePrio"
C;X60;K"collision"
C;X61;K"InBeta"
{{- range $index, $element := .Unit }}{{- if .UnitBalanceID.Valid }}
C;X1;Y{{ inc $index }};K{{ .UnitBalanceID.String }} 
{{- if .Level.Valid }}
C;X5;K{{ .Level.String }}
{{- end }}
C;X6;KFALSE 
{{- if .Type.Valid }}
C;X7;K{{ .Type.String }}
{{- end }} 
{{- if .Goldcost.Valid }}
C;X8;K{{ .Goldcost.String }}
{{- end }} 
{{- if .Lumbercost.Valid }}
C;X9;K{{ .Lumbercost.String }}
{{- end }} 
{{- if .GoldRep.Valid }}
C;X10;K{{ .GoldRep.String }}
{{- end }} 
{{- if .LumberRep.Valid }}
C;X11;K{{ .LumberRep.String }}
{{- end }} 
{{- if .Fmade.Valid }}
C;X12;K{{ .Fmade.String }}
{{- end }} 
{{- if .Fused.Valid }}
C;X13;K{{ .Fused.String }}
{{- end }} 
{{- if .Bountydice.Valid }}
C;X14;K{{ .Bountydice.String }}
{{- end }} 
{{- if .Bountysides.Valid }}
C;X15;K{{ .Bountysides.String }}
{{- end }} 
{{- if .Bountyplus.Valid }}
C;X16;K{{ .Bountyplus.String }}
{{- end }} 
{{- if .Lumberbountydice.Valid }}
C;X17;K{{ .Lumberbountydice.String }}
{{- end }} 
{{- if .Lumberbountysides.Valid }}
C;X18;K{{ .Lumberbountysides.String }}
{{- end }} 
{{- if .Lumberbountyplus.Valid }}
C;X19;K{{ .Lumberbountyplus.String }}
{{- end }} 
{{- if .StockMax.Valid }}
C;X20;K{{ .StockMax.String }}
{{- end }} 
{{- if .StockRegen.Valid }}
C;X21;K{{ .StockRegen.String }}
{{- end }} 
{{- if .StockStart.Valid }}
C;X22;K{{ .StockStart.String }}
{{- end }} 
{{- if .HP.Valid }}
C;X23;K{{ .HP.String }}
{{- end }} 
{{- if .RealHP.Valid }}
C;X24;K{{ .RealHP.String }}
{{- end }} 
{{- if .RegenHP.Valid }}
C;X25;K{{ .RegenHP.String }}
{{- end }} 
{{- if .RegenType.Valid }}
C;X26;K{{ .RegenType.String }}
{{- end }} 
{{- if .ManaN.Valid }}
C;X27;K{{ .ManaN.String }}
{{- end }} 
{{- if .RealM.Valid }}
C;X28;K{{ .RealM.String }}
{{- end }} 
{{- if .Mana0.Valid }}
C;X29;K{{ .Mana0.String }}
{{- end }} 
{{- if .RegenMana.Valid }}
C;X30;K{{ .RegenMana.String }}
{{- end }} 
{{- if .Def.Valid }}
C;X31;K{{ .Def.String }}
{{- end }} 
{{- if .DefUp.Valid }}
C;X32;K{{ .DefUp.String }}
{{- end }} 
{{- if .Realdef.Valid }}
C;X33;K{{ .Realdef.String }}
{{- end }} 
{{- if .DefType.Valid }}
C;X34;K{{ .DefType.String }}
{{- end }} 
{{- if .Spd.Valid }}
C;X35;K{{ .Spd.String }}
{{- end }} 
{{- if .MinSpd.Valid }}
C;X36;K{{ .MinSpd.String }}
{{- end }} 
{{- if .MaxSpd.Valid }}
C;X37;K{{ .MaxSpd.String }}
{{- end }} 
{{- if .Bldtm.Valid }}
C;X38;K{{ .Bldtm.String }}
{{- end }} 
{{- if .Reptm.Valid }}
C;X39;K{{ .Reptm.String }}
{{- end }} 
{{- if .Sight.Valid }}
C;X40;K{{ .Sight.String }}
{{- end }} 
{{- if .Nsight.Valid }}
C;X41;K{{ .Nsight.String }}
{{- end }} 
{{- if .STR.Valid }}
C;X42;K{{ .STR.String }}
{{- end }} 
{{- if .INT.Valid }}
C;X43;K{{ .INT.String }}
{{- end }} 
{{- if .AGI.Valid }}
C;X44;K{{ .AGI.String }}
{{- end }} 
{{- if .STRplus.Valid }}
C;X45;K{{ .STRplus.String }}
{{- end }} 
{{- if .INTplus.Valid }}
C;X46;K{{ .INTplus.String }}
{{- end }} 
{{- if .AGIplus.Valid }}
C;X47;K{{ .AGIplus.String }}
{{- end }} 
{{- if .AbilTest.Valid }}
C;X48;K{{ .AbilTest.String }}
{{- end }} 
{{- if .Primary.Valid }}
C;X49;K{{ .Primary.String }}
{{- end }} 
{{- if .Upgrades.Valid }}
C;X50;K{{ .Upgrades.String }}
{{- end }} 
{{- if .Tilesets.Valid }}
C;X51;K{{ .Tilesets.String }}
{{- end }} 
{{- if .Nbrandom.Valid }}
C;X52;K{{ .Nbrandom.String }}
{{- end }} 
{{- if .Isbldg.Valid }}
C;X53;K{{ .Isbldg.String }}
{{- end }} 
{{- if .PreventPlace.Valid }}
C;X54;K{{ .PreventPlace.String }}
{{- end }} 
{{- if .RequirePlace.Valid }}
C;X55;K{{ .RequirePlace.String }}
{{- end }} 
{{- if .Repulse.Valid }}
C;X56;K{{ .Repulse.String }}
{{- end }} 
{{- if .RepulseParam.Valid }}
C;X57;K{{ .RepulseParam.String }}
{{- end }} 
{{- if .RepulseGroup.Valid }}
C;X58;K{{ .RepulseGroup.String }}
{{- end }} 
{{- if .RepulsePrio.Valid }}
C;X59;K{{ .RepulsePrio.String }}
{{- end }} 
{{- if .Collision.Valid }}
C;X60;K{{ .Collision.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
