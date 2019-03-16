package templates

func GetUnitUITemplate() string {
	return `{{- define "UnitUI" }}ID;PWXL;N;E
B;X51;Y{{ .RowCount }};D0
C;X1;Y1;K"unitUIID"
C;X2;K"sortUI"
C;X3;K"file"
C;X4;K"fileVerFlags"
C;X5;K"unitSound"
C;X6;K"tilesetSpecific"
C;X7;K"name"
C;X8;K"unitClass"
C;X9;K"special"
C;X10;K"campaign"
C;X11;K"inEditor"
C;X12;K"hiddenInEditor"
C;X13;K"hostilePal"
C;X14;K"dropItems"
C;X15;K"nbmmIcon"
C;X16;K"useClickHelper"
C;X17;K"hideHeroBar"
C;X18;K"hideHeroMinimap"
C;X19;K"hideHeroDeathMsg"
C;X20;K"hideOnMinimap"
C;X21;K"blend"
C;X22;K"scale"
C;X23;K"scaleBull"
C;X24;K"maxPitch"
C;X25;K"maxRoll"
C;X26;K"elevPts"
C;X27;K"elevRad"
C;X28;K"fogRad"
C;X29;K"walk"
C;X30;K"run"
C;X31;K"selZ"
C;X32;K"weap1"
C;X33;K"weap2"
C;X34;K"teamColor"
C;X35;K"customTeamColor"
C;X36;K"armor"
C;X37;K"modelScale"
C;X38;K"red"
C;X39;K"green"
C;X40;K"blue"
C;X41;K"uberSplat"
C;X42;K"unitShadow"
C;X43;K"buildingShadow"
C;X44;K"shadowW"
C;X45;K"shadowH"
C;X46;K"shadowX"
C;X47;K"shadowY"
C;X48;K"shadowOnWater"
C;X49;K"selCircOnWater"
C;X50;K"occH"
C;X51;K"InBeta"
{{- range $index, $element := .UnitUI }}
C;X1;Y{{ inc $index }};K{{ .UnitUIID.String }}
C;X2;K{{ .SortUI.String }}
C;X3;K{{ .File.String }}
C;X4;K{{ .FileVerFlags.String }}
C;X5;K{{ .UnitSound.String }}
C;X6;K{{ .TilesetSpecific.String }}
C;X7;K{{ .Name.String }}
C;X8;K{{ .UnitClass.String }}
C;X9;K{{ .Special.String }}
C;X10;K{{ .Campaign.String }}
C;X11;K{{ .InEditor.String }}
C;X12;K{{ .HiddenInEditor.String }}
C;X13;K{{ .HostilePal.String }}
C;X14;K{{ .DropItems.String }}
C;X15;K{{ .NbmmIcon.String }}
C;X16;K{{ .UseClickHelper.String }}
C;X17;K{{ .HideHeroBar.String }}
C;X18;K{{ .HideHeroMinimap.String }}
C;X19;K{{ .HideHeroDeathMsg.String }}
C;X20;K{{ .HideOnMinimap.String }}
C;X21;K{{ .Blend.String }}
C;X22;K{{ .Scale.String }}
C;X23;K{{ .ScaleBull.String }}
C;X24;K{{ .MaxPitch.String }}
C;X25;K{{ .MaxRoll.String }}
C;X26;K{{ .ElevPts.String }}
C;X27;K{{ .ElevRad.String }}
C;X28;K{{ .FogRad.String }}
C;X29;K{{ .Walk.String }}
C;X30;K{{ .Run.String }}
C;X31;K{{ .SelZ.String }}
C;X32;K{{ .Weap1.String }}
C;X33;K{{ .Weap2.String }}
C;X34;K{{ .TeamColor.String }}
C;X35;K{{ .CustomTeamColor.String }}
C;X36;K{{ .Armor.String }}
C;X37;K{{ .ModelScale.String }}
C;X38;K{{ .Red.String }}
C;X39;K{{ .Green.String }}
C;X40;K{{ .Blue.String }}
C;X41;K{{ .UberSplat.String }}
C;X42;K{{ .UnitShadow.String }}
C;X43;K{{ .BuildingShadow.String }}
C;X44;K{{ .ShadowW.String }}
C;X45;K{{ .ShadowH.String }}
C;X46;K{{ .ShadowX.String }}
C;X47;K{{ .ShadowY.String }}
C;X48;K{{ .ShadowOnWater.String }}
C;X49;K{{ .SelCircOnWater.String }}
C;X50;K{{ .OccH.String }}
C;X51;K{{ .InBeta.String }}
{{- end }}
E
{{ end }}`
}