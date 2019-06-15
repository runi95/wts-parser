package templates

func GetItemDataTemplate() string {
	return `{{- define "ItemData" }}ID;PWXL;N;E
B;X35;Y{{ .RowCount }};D0
C;X1;Y1;K"itemID"
C;X2;K"comment"
C;X3;K"scriptname"
C;X4;K"version"
C;X5;K"class"
C;X6;K"Level"
C;X7;K"oldLevel"
C;X8;K"abilList"
C;X9;K"cooldownID"
C;X10;K"ignoreCD"
C;X11;K"uses"
C;X12;K"prio"
C;X13;K"usable"
C;X14;K"perishable"
C;X15;K"droppable"
C;X16;K"pawnable"
C;X17;K"sellable"
C;X18;K"pickRandom"
C;X19;K"powerup"
C;X20;K"drop"
C;X21;K"stockMax"
C;X22;K"stockRegen"
C;X23;K"stockStart"
C;X24;K"goldcost"
C;X25;K"lumbercost"
C;X26;K"HP"
C;X27;K"morph"
C;X28;K"armor"
C;X29;K"file"
C;X30;K"scale"
C;X31;K"selSize"
C;X32;K"colorR"
C;X33;K"colorG"
C;X34;K"colorB"
C;X35;K"InBeta"
{{- range $index, $element := .Item }}{{- if .ItemID.Valid }}
C;X1;Y{{ inc $index }};K{{ .ItemID.String }} 
{{- if .Comment.Valid }}
C;X2;K{{ .Comment.String }}
{{- end }}
{{- if .Scriptname.Valid }}
C;X3;K{{ .Scriptname.String }}
{{- end }}
{{- if .Version.Valid }}
C;X4;K{{ .Version.String }}
{{- end }}
{{- if .Class.Valid }}
C;X5;K{{ .Class.String }}
{{- end }}
{{- if .Level.Valid }}
C;X6;K{{ .Level.String }}
{{- end }}
{{- if .OldLevel.Valid }}
C;X7;K{{ .OldLevel.String }}
{{- end }}
{{- if .AbilList.Valid }}
C;X8;K{{ .AbilList.String }}
{{- end }}
{{- if .CooldownID.Valid }}
C;X9;K{{ .CooldownID.String }}
{{- end }}
{{- if .IgnoreCD.Valid }}
C;X10;K{{ .IgnoreCD.String }}
{{- end }}
{{- if .Uses.Valid }}
C;X11;K{{ .Uses.String }}
{{- end }}
{{- if .Prio.Valid }}
C;X12;K{{ .Prio.String }}
{{- end }}
{{- if .Usable.Valid }}
C;X13;K{{ .Usable.String }}
{{- end }}
{{- if .Perishable.Valid }}
C;X14;K{{ .Perishable.String }}
{{- end }}
{{- if .Droppable.Valid }}
C;X15;K{{ .Droppable.String }}
{{- end }}
{{- if .Pawnable.Valid }}
C;X16;K{{ .Pawnable.String }}
{{- end }}
{{- if .Sellable.Valid }}
C;X17;K{{ .Sellable.String }}
{{- end }}
{{- if .PickRandom.Valid }}
C;X18;K{{ .PickRandom.String }}
{{- end }}
{{- if .Powerup.Valid }}
C;X19;K{{ .Powerup.String }}
{{- end }}
{{- if .Drop.Valid }}
C;X20;K{{ .Drop.String }}
{{- end }}
{{- if .StockMax.Valid }}
C;X21;K{{ .StockMax.String }}
{{- end }}
{{- if .StockRegen.Valid }}
C;X22;K{{ .StockRegen.String }}
{{- end }}
{{- if .StockStart.Valid }}
C;X23;K{{ .StockStart.String }}
{{- end }}
{{- if .Goldcost.Valid }}
C;X24;K{{ .Goldcost.String }}
{{- end }}
{{- if .Lumbercost.Valid }}
C;X25;K{{ .Lumbercost.String }}
{{- end }}
{{- if .HP.Valid }}
C;X26;K{{ .HP.String }}
{{- end }}
{{- if .Morph.Valid }}
C;X27;K{{ .Morph.String }}
{{- end }}
{{- if .Armor.Valid }}
C;X28;K{{ .Armor.String }}
{{- end }}
{{- if .File.Valid }}
C;X29;K{{ .File.String }}
{{- end }}
{{- if .Scale.Valid }}
C;X30;K{{ .Scale.String }}
{{- end }}
{{- if .SelSize.Valid }}
C;X31;K{{ .SelSize.String }}
{{- end }}
{{- if .ColorR.Valid }}
C;X32;K{{ .ColorR.String }}
{{- end }}
{{- if .ColorG.Valid }}
C;X33;K{{ .ColorG.String }}
{{- end }}
{{- if .ColorB.Valid }}
C;X34;K{{ .ColorB.String }}
{{- end }}
{{- if .InBeta.Valid }}
C;X35;K{{ .InBeta.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
