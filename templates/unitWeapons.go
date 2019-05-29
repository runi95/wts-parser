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
{{- range $index, $element := .Unit }}{{- if .UnitWeapID.Valid }}
C;X1;Y{{ inc $index }};K{{ .UnitWeapID.String }} 
{{- if .WeapsOn.Valid }}
C;X5;K{{ .WeapsOn.String }}
{{- end }} 
{{- if .Acquire.Valid }}
C;X6;K{{ .Acquire.String }}
{{- end }} 
{{- if .MinRange.Valid }}
C;X7;K{{ .MinRange.String }}
{{- end }} 
{{- if .Castpt.Valid }}
C;X8;K{{ .Castpt.String }}
{{- end }} 
{{- if .Castbsw.Valid }}
C;X9;K{{ .Castbsw.String }}
{{- end }} 
{{- if .LaunchX.Valid }}
C;X10;K{{ .LaunchX.String }}
{{- end }} 
{{- if .LaunchY.Valid }}
C;X11;K{{ .LaunchY.String }}
{{- end }} 
{{- if .LaunchZ.Valid }}
C;X12;K{{ .LaunchZ.String }}
{{- end }} 
{{- if .LaunchSwimZ.Valid }}
C;X13;K{{ .LaunchSwimZ.String }}
{{- end }} 
{{- if .ImpactZ.Valid }}
C;X14;K{{ .ImpactZ.String }}
{{- end }} 
{{- if .ImpactSwimZ.Valid }}
C;X15;K{{ .ImpactSwimZ.String }}
{{- end }} 
{{- if .WeapType1.Valid }}
C;X16;K{{ .WeapType1.String }}
{{- end }} 
{{- if .Targs1.Valid }}
C;X17;K{{ .Targs1.String }}
{{- end }} 
{{- if .ShowUI1.Valid }}
C;X18;K{{ .ShowUI1.String }}
{{- end }} 
{{- if .RangeN1.Valid }}
C;X19;K{{ .RangeN1.String }}
{{- end }} 
{{- if .RngTst.Valid }}
C;X20;K{{ .RngTst.String }}
{{- end }} 
{{- if .RngBuff1.Valid }}
C;X21;K{{ .RngBuff1.String }}
{{- end }} 
{{- if .AtkType1.Valid }}
C;X22;K{{ .AtkType1.String }}
{{- end }} 
{{- if .WeapTp1.Valid }}
C;X23;K{{ .WeapTp1.String }}
{{- end }} 
{{- if .Cool1.Valid }}
C;X24;K{{ .Cool1.String }}
{{- end }} 
{{- if .Mincool1.Valid }}
C;X25;K{{ .Mincool1.String }}
{{- end }} 
{{- if .Dice1.Valid }}
C;X26;K{{ .Dice1.String }}
{{- end }} 
{{- if .Sides1.Valid }}
C;X27;K{{ .Sides1.String }}
{{- end }} 
{{- if .Dmgplus1.Valid }}
C;X28;K{{ .Dmgplus1.String }}
{{- end }} 
{{- if .DmgUp1.Valid }}
C;X29;K{{ .DmgUp1.String }}
{{- end }} 
{{- if .Mindmg1.Valid }}
C;X30;K{{ .Mindmg1.String }}
{{- end }} 
{{- if .Avgdmg1.Valid }}
C;X31;K{{ .Avgdmg1.String }}
{{- end }} 
{{- if .Maxdmg1.Valid }}
C;X32;K{{ .Maxdmg1.String }}
{{- end }} 
{{- if .Dmgpt1.Valid }}
C;X33;K{{ .Dmgpt1.String }}
{{- end }} 
{{- if .BackSw1.Valid }}
C;X34;K{{ .BackSw1.String }}
{{- end }} 
{{- if .Farea1.Valid }}
C;X35;K{{ .Farea1.String }}
{{- end }} 
{{- if .Harea1.Valid }}
C;X36;K{{ .Harea1.String }}
{{- end }} 
{{- if .Qarea1.Valid }}
C;X37;K{{ .Qarea1.String }}
{{- end }} 
{{- if .Hfact1.Valid }}
C;X38;K{{ .Hfact1.String }}
{{- end }} 
{{- if .Qfact1.Valid }}
C;X39;K{{ .Qfact1.String }}
{{- end }} 
{{- if .SplashTargs1.Valid }}
C;X40;K{{ .SplashTargs1.String }}
{{- end }} 
{{- if .TargCount1.Valid }}
C;X41;K{{ .TargCount1.String }}
{{- end }} 
{{- if .DamageLoss1.Valid }}
C;X42;K{{ .DamageLoss1.String }}
{{- end }} 
{{- if .SpillDist1.Valid }}
C;X43;K{{ .SpillDist1.String }}
{{- end }} 
{{- if .SpillRadius1.Valid }}
C;X44;K{{ .SpillRadius1.String }}
{{- end }} 
{{- if .DmgUpg.Valid }}
C;X45;K{{ .DmgUpg.String }}
{{- end }} 
{{- if .Dmod1.Valid }}
C;X46;K{{ .Dmod1.String }}
{{- end }} 
{{- if .DPS.Valid }}
C;X47;K{{ .DPS.String }}
{{- end }} 
{{- if .WeapType2.Valid }}
C;X48;K{{ .WeapType2.String }}
{{- end }} 
{{- if .Targs2.Valid }}
C;X49;K{{ .Targs2.String }}
{{- end }} 
{{- if .ShowUI2.Valid }}
C;X50;K{{ .ShowUI2.String }}
{{- end }} 
{{- if .RangeN2.Valid }}
C;X51;K{{ .RangeN2.String }}
{{- end }} 
{{- if .RngTst2.Valid }}
C;X52;K{{ .RngTst2.String }}
{{- end }} 
{{- if .RngBuff2.Valid }}
C;X53;K{{ .RngBuff2.String }}
{{- end }} 
{{- if .AtkType2.Valid }}
C;X54;K{{ .AtkType2.String }}
{{- end }} 
{{- if .WeapTp2.Valid }}
C;X55;K{{ .WeapTp2.String }}
{{- end }} 
{{- if .Cool2.Valid }}
C;X56;K{{ .Cool2.String }}
{{- end }} 
{{- if .Mincool2.Valid }}
C;X57;K{{ .Mincool2.String }}
{{- end }} 
{{- if .Dice2.Valid }}
C;X58;K{{ .Dice2.String }}
{{- end }} 
{{- if .Sides2.Valid }}
C;X59;K{{ .Sides2.String }}
{{- end }} 
{{- if .Dmgplus2.Valid }}
C;X60;K{{ .Dmgplus2.String }}
{{- end }} 
{{- if .DmgUp2.Valid }}
C;X61;K{{ .DmgUp2.String }}
{{- end }} 
{{- if .Mindmg2.Valid }}
C;X62;K{{ .Mindmg2.String }}
{{- end }} 
{{- if .Avgdmg2.Valid }}
C;X63;K{{ .Avgdmg2.String }}
{{- end }} 
{{- if .Maxdmg2.Valid }}
C;X64;K{{ .Maxdmg2.String }}
{{- end }} 
{{- if .Dmgpt2.Valid }}
C;X65;K{{ .Dmgpt2.String }}
{{- end }} 
{{- if .BackSw2.Valid }}
C;X66;K{{ .BackSw2.String }}
{{- end }} 
{{- if .Farea2.Valid }}
C;X67;K{{ .Farea2.String }}
{{- end }} 
{{- if .Harea2.Valid }}
C;X68;K{{ .Harea2.String }}
{{- end }} 
{{- if .Qarea2.Valid }}
C;X69;K{{ .Qarea2.String }}
{{- end }} 
{{- if .Hfact2.Valid }}
C;X70;K{{ .Hfact2.String }}
{{- end }} 
{{- if .Qfact2.Valid }}
C;X71;K{{ .Qfact2.String }}
{{- end }} 
{{- if .SplashTargs2.Valid }}
C;X72;K{{ .SplashTargs2.String }}
{{- end }} 
{{- if .TargCount2.Valid }}
C;X73;K{{ .TargCount2.String }}
{{- end }} 
{{- if .DamageLoss2.Valid }}
C;X74;K{{ .DamageLoss2.String }}
{{- end }} 
{{- if .SpillDist2.Valid }}
C;X75;K{{ .SpillDist2.String }}
{{- end }} 
{{- if .SpillRadius2.Valid }}
C;X76;K{{ .SpillRadius2.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
