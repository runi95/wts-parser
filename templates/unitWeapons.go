package templates

func GetUnitWeaponsTemplate() string {
	return `{{- define "UnitWeapons" }}ID;PWXL;N;E
B;X79;Y{{ .RowCount }};D0
C;X1;Y1;K"unitWeapID"
C;X2;K"sortWeap"
C;X3;K"sort2"
C;X4;K"comment(s)"
C;X5;K"weapsOn"
C;X6;K"acquire"
C;X7;K"minRange"
C;X8;K"castpt"
C;X9;K"castbsw"
C;X10;K"launchX"
C;X11;K"launchY"
C;X12;K"launchZ"
C;X13;K"launchSwimZ"
C;X14;K"impactZ"
C;X15;K"impactSwimZ"
C;X16;K"weapType1"
C;X17;K"targs1"
C;X18;K"showUI1"
C;X19;K"rangeN1"
C;X20;K"RngTst"
C;X21;K"RngBuff1"
C;X22;K"atkType1"
C;X23;K"weapTp1"
C;X24;K"cool1"
C;X25;K"mincool1"
C;X26;K"dice1"
C;X27;K"sides1"
C;X28;K"dmgplus1"
C;X29;K"dmgUp1"
C;X30;K"mindmg1"
C;X31;K"avgdmg1"
C;X32;K"maxdmg1"
C;X33;K"dmgpt1"
C;X34;K"backSw1"
C;X35;K"Farea1"
C;X36;K"Harea1"
C;X37;K"Qarea1"
C;X38;K"Hfact1"
C;X39;K"Qfact1"
C;X40;K"splashTargs1"
C;X41;K"targCount1"
C;X42;K"damageLoss1"
C;X43;K"spillDist1"
C;X44;K"spillRadius1"
C;X45;K"DmgUpg"
C;X46;K"dmod1"
C;X47;K"DPS"
C;X48;K"weapType2"
C;X49;K"targs2"
C;X50;K"showUI2"
C;X51;K"rangeN2"
C;X52;K"RngTst2"
C;X53;K"RngBuff2"
C;X54;K"atkType2"
C;X55;K"weapTp2"
C;X56;K"cool2"
C;X57;K"mincool2"
C;X58;K"dice2"
C;X59;K"sides2"
C;X60;K"dmgplus2"
C;X61;K"dmgUp2"
C;X62;K"mindmg2"
C;X63;K"avgdmg2"
C;X64;K"maxdmg2"
C;X65;K"dmgpt2"
C;X66;K"backSw2"
C;X67;K"Farea2"
C;X68;K"Harea2"
C;X69;K"Qarea2"
C;X70;K"Hfact2"
C;X71;K"Qfact2"
C;X72;K"splashTargs2"
C;X73;K"targCount2"
C;X74;K"damageLoss2"
C;X75;K"spillDist2"
C;X76;K"spillRadius2"
C;X77;K"InBeta"
{{- range $index, $element := .UnitWeapons }}
C;X1;Y{{ inc $index }};K{{ .UnitWeapID.String }}
C;X2;K{{ .SortWeap.String }}
C;X3;K{{ .Sort2.String }}
C;X4;K{{ .Comment.String }}
C;X5;K{{ .WeapsOn.String }}
C;X6;K{{ .Acquire.String }}
C;X7;K{{ .MinRange.String }}
C;X8;K{{ .Castpt.String }}
C;X9;K{{ .Castbsw.String }}
C;X10;K{{ .LaunchX.String }}
C;X11;K{{ .LaunchY.String }}
C;X12;K{{ .LaunchZ.String }}
C;X13;K{{ .LaunchSwimZ.String }}
C;X14;K{{ .ImpactZ.String }}
C;X15;K{{ .ImpactSwimZ.String }}
C;X16;K{{ .WeapType1.String }}
C;X17;K{{ .Targs1.String }}
C;X18;K{{ .ShowUI1.String }}
C;X19;K{{ .RangeN1.String }}
C;X20;K{{ .RngTst.String }}
C;X21;K{{ .RngBuff1.String }}
C;X22;K{{ .AtkType1.String }}
C;X23;K{{ .WeapTp1.String }}
C;X24;K{{ .Cool1.String }}
C;X25;K{{ .Mincool1.String }}
C;X26;K{{ .Dice1.String }}
C;X27;K{{ .Sides1.String }}
C;X28;K{{ .Dmgplus1.String }}
C;X29;K{{ .DmgUp1.String }}
C;X30;K{{ .Mindmg1.String }}
C;X31;K{{ .Avgdmg1.String }}
C;X32;K{{ .Maxdmg1.String }}
C;X33;K{{ .Dmgpt1.String }}
C;X34;K{{ .BackSw1.String }}
C;X35;K{{ .Farea1.String }}
C;X36;K{{ .Harea1.String }}
C;X37;K{{ .Qarea1.String }}
C;X38;K{{ .Hfact1.String }}
C;X39;K{{ .Qfact1.String }}
C;X40;K{{ .SplashTargs1.String }}
C;X41;K{{ .TargCount1.String }}
C;X42;K{{ .DamageLoss1.String }}
C;X43;K{{ .SpillDist1.String }}
C;X44;K{{ .SpillRadius1.String }}
C;X45;K{{ .DmgUpg.String }}
C;X46;K{{ .Dmod1.String }}
C;X47;K{{ .DPS.String }}
C;X48;K{{ .WeapType2.String }}
C;X49;K{{ .Targs2.String }}
C;X50;K{{ .ShowUI2.String }}
C;X51;K{{ .RangeN2.String }}
C;X52;K{{ .RngTst2.String }}
C;X53;K{{ .RngBuff2.String }}
C;X54;K{{ .AtkType2.String }}
C;X55;K{{ .WeapTp2.String }}
C;X56;K{{ .Cool2.String }}
C;X57;K{{ .Mincool2.String }}
C;X58;K{{ .Dice2.String }}
C;X59;K{{ .Sides2.String }}
C;X60;K{{ .Dmgplus2.String }}
C;X61;K{{ .DmgUp2.String }}
C;X62;K{{ .Mindmg2.String }}
C;X63;K{{ .Avgdmg2.String }}
C;X64;K{{ .Maxdmg2.String }}
C;X65;K{{ .Dmgpt2.String }}
C;X66;K{{ .BackSw2.String }}
C;X67;K{{ .Farea2.String }}
C;X68;K{{ .Harea2.String }}
C;X69;K{{ .Qarea2.String }}
C;X70;K{{ .Hfact2.String }}
C;X71;K{{ .Qfact2.String }}
C;X72;K{{ .SplashTargs2.String }}
C;X73;K{{ .TargCount2.String }}
C;X74;K{{ .DamageLoss2.String }}
C;X75;K{{ .SpillDist2.String }}
C;X76;K{{ .SpillRadius2.String }}
C;X77;K{{ .InBeta.String }}
{{- end }}
E
{{ end }}`
}
