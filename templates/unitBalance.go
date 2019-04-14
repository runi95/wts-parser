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
{{- range $index, $element := .UnitBalance }}
C;X1;Y{{ inc $index }};K{{ .UnitBalanceID.String }}
C;X5;K{{ .Level.String }}
C;X6;KFALSE
C;X7;K{{ .Type.String }}
C;X8;K{{ .Goldcost.String }}
C;X9;K{{ .Lumbercost.String }}
C;X10;K{{ .GoldRep.String }}
C;X11;K{{ .LumberRep.String }}
C;X12;K{{ .Fmade.String }}
C;X13;K{{ .Fused.String }}
C;X14;K{{ .Bountydice.String }}
C;X15;K{{ .Bountysides.String }}
C;X16;K{{ .Bountyplus.String }}
C;X17;K{{ .Lumberbountydice.String }}
C;X18;K{{ .Lumberbountysides.String }}
C;X19;K{{ .Lumberbountyplus.String }}
C;X20;K{{ .StockMax.String }}
C;X21;K{{ .StockRegen.String }}
C;X22;K{{ .StockStart.String }}
C;X23;K{{ .HP.String }}
C;X24;K{{ .RealHP.String }}
C;X25;K{{ .RegenHP.String }}
C;X26;K{{ .RegenType.String }}
C;X27;K{{ .ManaN.String }}
C;X28;K{{ .RealM.String }}
C;X29;K{{ .Mana0.String }}
C;X30;K{{ .RegenMana.String }}
C;X31;K{{ .Def.String }}
C;X32;K{{ .DefUp.String }}
C;X33;K{{ .Realdef.String }}
C;X34;K{{ .DefType.String }}
C;X35;K{{ .Spd.String }}
C;X36;K{{ .MinSpd.String }}
C;X37;K{{ .MaxSpd.String }}
C;X38;K{{ .Bldtm.String }}
C;X39;K{{ .Reptm.String }}
C;X40;K{{ .Sight.String }}
C;X41;K{{ .Nsight.String }}
C;X42;K{{ .STR.String }}
C;X43;K{{ .INT.String }}
C;X44;K{{ .AGI.String }}
C;X45;K{{ .STRplus.String }}
C;X46;K{{ .INTplus.String }}
C;X47;K{{ .AGIplus.String }}
C;X48;K{{ .AbilTest.String }}
C;X49;K{{ .Primary.String }}
C;X50;K{{ .Upgrades.String }}
C;X51;K{{ .Tilesets.String }}
C;X52;K{{ .Nbrandom.String }}
C;X53;K{{ .Isbldg.String }}
C;X54;K{{ .PreventPlace.String }}
C;X55;K{{ .RequirePlace.String }}
C;X56;K{{ .Repulse.String }}
C;X57;K{{ .RepulseParam.String }}
C;X58;K{{ .RepulseGroup.String }}
C;X59;K{{ .RepulsePrio.String }}
C;X60;K{{ .Collision.String }}
{{- end }}
E
{{ end }}`
}
