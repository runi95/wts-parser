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
{{- range $index, $element := .Unit }}{{- if .UnitUIID.Valid }}
C;X1;Y{{ inc $index }};K{{ .UnitUIID.String }} 
{{- if .File.Valid }}
C;X3;K{{ .File.String }}
{{- end }} 
{{- if .FileVerFlags.Valid }}
C;X4;K{{ .FileVerFlags.String }}
{{- end }} 
{{- if .UnitSound.Valid }}
C;X5;K{{ .UnitSound.String }}
{{- end }} 
{{- if .TilesetSpecific.Valid }}
C;X6;K{{ .TilesetSpecific.String }}
{{- end }} 
{{- if .Name.Valid }}
C;X7;K{{ .Name.String }}
{{- end }} 
{{- if .UnitClass.Valid }}
C;X8;K{{ .UnitClass.String }}
{{- end }} 
{{- if .Special.Valid }}
C;X9;K{{ .Special.String }}
{{- end }} 
{{- if .Campaign.Valid }}
C;X10;K{{ .Campaign.String }}
{{- end }} 
{{- if .InEditor.Valid }}
C;X11;K{{ .InEditor.String }}
{{- end }} 
{{- if .HiddenInEditor.Valid }}
C;X12;K{{ .HiddenInEditor.String }}
{{- end }} 
{{- if .HostilePal.Valid }}
C;X13;K{{ .HostilePal.String }}
{{- end }} 
{{- if .DropItems.Valid }}
C;X14;K{{ .DropItems.String }}
{{- end }} 
{{- if .NbmmIcon.Valid }}
C;X15;K{{ .NbmmIcon.String }}
{{- end }} 
{{- if .UseClickHelper.Valid }}
C;X16;K{{ .UseClickHelper.String }}
{{- end }} 
{{- if .HideHeroBar.Valid }}
C;X17;K{{ .HideHeroBar.String }}
{{- end }} 
{{- if .HideHeroMinimap.Valid }}
C;X18;K{{ .HideHeroMinimap.String }}
{{- end }} 
{{- if .HideHeroDeathMsg.Valid }}
C;X19;K{{ .HideHeroDeathMsg.String }}
{{- end }} 
{{- if .HideOnMinimap.Valid }}
C;X20;K{{ .HideOnMinimap.String }}
{{- end }} 
{{- if .Blend.Valid }}
C;X21;K{{ .Blend.String }}
{{- end }} 
{{- if .Scale.Valid }}
C;X22;K{{ .Scale.String }}
{{- end }} 
{{- if .ScaleBull.Valid }}
C;X23;K{{ .ScaleBull.String }}
{{- end }} 
{{- if .MaxPitch.Valid }}
C;X24;K{{ .MaxPitch.String }}
{{- end }} 
{{- if .MaxRoll.Valid }}
C;X25;K{{ .MaxRoll.String }}
{{- end }} 
{{- if .ElevPts.Valid }}
C;X26;K{{ .ElevPts.String }}
{{- end }} 
{{- if .ElevRad.Valid }}
C;X27;K{{ .ElevRad.String }}
{{- end }} 
{{- if .FogRad.Valid }}
C;X28;K{{ .FogRad.String }}
{{- end }} 
{{- if .Walk.Valid }}
C;X29;K{{ .Walk.String }}
{{- end }} 
{{- if .Run.Valid }}
C;X30;K{{ .Run.String }}
{{- end }} 
{{- if .SelZ.Valid }}
C;X31;K{{ .SelZ.String }}
{{- end }} 
{{- if .Weap1.Valid }}
C;X32;K{{ .Weap1.String }}
{{- end }} 
{{- if .Weap2.Valid }}
C;X33;K{{ .Weap2.String }}
{{- end }} 
{{- if .TeamColor.Valid }}
C;X34;K{{ .TeamColor.String }}
{{- end }} 
{{- if .CustomTeamColor.Valid }}
C;X35;K{{ .CustomTeamColor.String }}
{{- end }} 
{{- if .Armor.Valid }}
C;X36;K{{ .Armor.String }}
{{- end }} 
{{- if .ModelScale.Valid }}
C;X37;K{{ .ModelScale.String }}
{{- end }} 
{{- if .Red.Valid }}
C;X38;K{{ .Red.String }}
{{- end }} 
{{- if .Green.Valid }}
C;X39;K{{ .Green.String }}
{{- end }} 
{{- if .Blue.Valid }}
C;X40;K{{ .Blue.String }}
{{- end }} 
{{- if .UberSplat.Valid }}
C;X41;K{{ .UberSplat.String }}
{{- end }} 
{{- if .UnitShadow.Valid }}
C;X42;K{{ .UnitShadow.String }}
{{- end }} 
{{- if .BuildingShadow.Valid }}
C;X43;K{{ .BuildingShadow.String }}
{{- end }} 
{{- if .ShadowW.Valid }}
C;X44;K{{ .ShadowW.String }}
{{- end }} 
{{- if .ShadowH.Valid }}
C;X45;K{{ .ShadowH.String }}
{{- end }} 
{{- if .ShadowX.Valid }}
C;X46;K{{ .ShadowX.String }}
{{- end }} 
{{- if .ShadowY.Valid }}
C;X47;K{{ .ShadowY.String }}
{{- end }} 
{{- if .ShadowOnWater.Valid }}
C;X48;K{{ .ShadowOnWater.String }}
{{- end }} 
{{- if .SelCircOnWater.Valid }}
C;X49;K{{ .SelCircOnWater.String }}
{{- end }} 
{{- if .OccH.Valid }}
C;X50;K{{ .OccH.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
