package templates

func GetAbilityDataTemplate() string {
	return `{{- define "AbilityData" }}ID;PWXL;N;E
B;X95;Y{{ .RowCount }};D0
C;X1;Y1;K"alias"
C;X2;K"code"
C;X3;K"comments"
C;X4;K"version"
C;X5;K"useInEditor"
C;X6;K"hero"
C;X7;K"item"
C;X8;K"sort"
C;X9;K"race"
C;X10;K"checkDep"
C;X11;K"levels"
C;X12;K"reqLevel"
C;X13;K"levelSkip"
C;X14;K"priority"
C;X15;K"targs1"
C;X16;K"Cast1"
C;X17;K"Dur1"
C;X18;K"HeroDur1"
C;X19;K"Cool1"
C;X20;K"Cost1"
C;X21;K"Area1"
C;X22;K"Rng1"
C;X23;K"DataA1"
C;X24;K"DataB1"
C;X25;K"DataC1"
C;X26;K"DataD1"
C;X27;K"DataE1"
C;X28;K"DataF1"
C;X29;K"DataG1"
C;X30;K"DataH1"
C;X31;K"DataI1"
C;X32;K"UnitID1"
C;X33;K"BuffID1"
C;X34;K"EfctID1"
C;X35;K"targs2"
C;X36;K"Cast2"
C;X37;K"Dur2"
C;X38;K"HeroDur2"
C;X39;K"Cool2"
C;X40;K"Cost2"
C;X41;K"Area2"
C;X42;K"Rng2"
C;X43;K"DataA2"
C;X44;K"DataB2"
C;X45;K"DataC2"
C;X46;K"DataD2"
C;X47;K"DataE2"
C;X48;K"DataF2"
C;X49;K"DataG2"
C;X50;K"DataH2"
C;X51;K"DataI2"
C;X52;K"UnitID2"
C;X53;K"BuffID2"
C;X54;K"EfctID2"
C;X55;K"targs3"
C;X56;K"Cast3"
C;X57;K"Dur3"
C;X58;K"HeroDur3"
C;X59;K"Cool3"
C;X60;K"Cost3"
C;X61;K"Area3"
C;X62;K"Rng3"
C;X63;K"DataA3"
C;X64;K"DataB3"
C;X65;K"DataC3"
C;X66;K"DataD3"
C;X67;K"DataE3"
C;X68;K"DataF3"
C;X69;K"DataG3"
C;X70;K"DataH3"
C;X71;K"DataI3"
C;X72;K"UnitID3"
C;X73;K"BuffID3"
C;X74;K"EfctID3"
C;X75;K"targs4"
C;X76;K"Cast4"
C;X77;K"Dur4"
C;X78;K"HeroDur4"
C;X79;K"Cool4"
C;X80;K"Cost4"
C;X81;K"Area4"
C;X82;K"Rng4"
C;X83;K"DataA4"
C;X84;K"DataB4"
C;X85;K"DataC4"
C;X86;K"DataD4"
C;X87;K"DataE4"
C;X88;K"DataF4"
C;X89;K"DataG4"
C;X90;K"DataH4"
C;X91;K"DataI4"
C;X92;K"UnitID4"
C;X93;K"BuffID4"
C;X94;K"EfctID4"
C;X95;K"InBeta"
{{- range $index, $element := .Ability }}{{- if .Alias.Valid }}
C;X1;Y{{ inc $index }};K{{ .Alias.String }}
{{- if .Code.Valid }}
C;X2;K{{ .Code.String }}
{{- end }}
{{- if .Comments.Valid }}
C;X3;K{{ .Comments.String }}
{{- end }}
{{- if .Version.Valid }}
C;X4;K{{ .Version.String }}
{{- end }}
{{- if .UseInEditor.Valid }}
C;X5;K{{ .UseInEditor.String }}
{{- end }}
{{- if .Hero.Valid }}
C;X6;K{{ .Hero.String }}
{{- end }}
{{- if .Item.Valid }}
C;X7;K{{ .Item.String }}
{{- end }}
{{- if .Sort.Valid }}
C;X8;K{{ .Sort.String }}
{{- end }}
{{- if .Race.Valid }}
C;X9;K{{ .Race.String }}
{{- end }}
{{- if .CheckDep.Valid }}
C;X10;K{{ .CheckDep.String }}
{{- end }}
{{- if .Levels.Valid }}
C;X11;K{{ .Levels.String }}
{{- end }}
{{- if .ReqLevel.Valid }}
C;X12;K{{ .ReqLevel.String }}
{{- end }}
{{- if .LevelSkip.Valid }}
C;X13;K{{ .LevelSkip.String }}
{{- end }}
{{- if .Priority.Valid }}
C;X14;K{{ .Priority.String }}
{{- end }}
{{- if .Targs1.Valid }}
C;X15;K{{ .Targs1.String }}
{{- end }}
{{- if .Cast1.Valid }}
C;X16;K{{ .Cast1.String }}
{{- end }}
{{- if .Dur1.Valid }}
C;X17;K{{ .Dur1.String }}
{{- end }}
{{- if .HeroDur1.Valid }}
C;X18;K{{ .HeroDur1.String }}
{{- end }}
{{- if .Cool1.Valid }}
C;X19;K{{ .Cool1.String }}
{{- end }}
{{- if .Cost1.Valid }}
C;X20;K{{ .Cost1.String }}
{{- end }}
{{- if .Area1.Valid }}
C;X21;K{{ .Area1.String }}
{{- end }}
{{- if .Rng1.Valid }}
C;X22;K{{ .Rng1.String }}
{{- end }}
{{- if .DataA1.Valid }}
C;X23;K{{ .DataA1.String }}
{{- end }}
{{- if .DataB1.Valid }}
C;X24;K{{ .DataB1.String }}
{{- end }}
{{- if .DataC1.Valid }}
C;X25;K{{ .DataC1.String }}
{{- end }}
{{- if .DataD1.Valid }}
C;X26;K{{ .DataD1.String }}
{{- end }}
{{- if .DataE1.Valid }}
C;X27;K{{ .DataE1.String }}
{{- end }}
{{- if .DataF1.Valid }}
C;X28;K{{ .DataF1.String }}
{{- end }}
{{- if .DataG1.Valid }}
C;X29;K{{ .DataG1.String }}
{{- end }}
{{- if .DataH1.Valid }}
C;X30;K{{ .DataH1.String }}
{{- end }}
{{- if .DataI1.Valid }}
C;X31;K{{ .DataI1.String }}
{{- end }}
{{- if .UnitID1.Valid }}
C;X32;K{{ .UnitID1.String }}
{{- end }}
{{- if .BuffID1.Valid }}
C;X33;K{{ .BuffID1.String }}
{{- end }}
{{- if .EfctID1.Valid }}
C;X34;K{{ .EfctID1.String }}
{{- end }}
{{- if .Targs2.Valid }}
C;X35;K{{ .Targs2.String }}
{{- end }}
{{- if .Cast2.Valid }}
C;X36;K{{ .Cast2.String }}
{{- end }}
{{- if .Dur2.Valid }}
C;X37;K{{ .Dur2.String }}
{{- end }}
{{- if .HeroDur2.Valid }}
C;X38;K{{ .HeroDur2.String }}
{{- end }}
{{- if .Cool2.Valid }}
C;X39;K{{ .Cool2.String }}
{{- end }}
{{- if .Cost2.Valid }}
C;X40;K{{ .Cost2.String }}
{{- end }}
{{- if .Area2.Valid }}
C;X41;K{{ .Area2.String }}
{{- end }}
{{- if .Rng2.Valid }}
C;X42;K{{ .Rng2.String }}
{{- end }}
{{- if .DataA2.Valid }}
C;X43;K{{ .DataA2.String }}
{{- end }}
{{- if .DataB2.Valid }}
C;X44;K{{ .DataB2.String }}
{{- end }}
{{- if .DataC2.Valid }}
C;X45;K{{ .DataC2.String }}
{{- end }}
{{- if .DataD2.Valid }}
C;X46;K{{ .DataD2.String }}
{{- end }}
{{- if .DataE2.Valid }}
C;X47;K{{ .DataE2.String }}
{{- end }}
{{- if .DataF2.Valid }}
C;X48;K{{ .DataF2.String }}
{{- end }}
{{- if .DataG2.Valid }}
C;X49;K{{ .DataG2.String }}
{{- end }}
{{- if .DataH2.Valid }}
C;X50;K{{ .DataH2.String }}
{{- end }}
{{- if .DataI2.Valid }}
C;X51;K{{ .DataI2.String }}
{{- end }}
{{- if .UnitID2.Valid }}
C;X52;K{{ .UnitID2.String }}
{{- end }}
{{- if .BuffID2.Valid }}
C;X53;K{{ .BuffID2.String }}
{{- end }}
{{- if .EfctID2.Valid }}
C;X54;K{{ .EfctID2.String }}
{{- end }}
{{- if .Targs3.Valid }}
C;X55;K{{ .Targs3.String }}
{{- end }}
{{- if .Cast3.Valid }}
C;X56;K{{ .Cast3.String }}
{{- end }}
{{- if .Dur3.Valid }}
C;X57;K{{ .Dur3.String }}
{{- end }}
{{- if .HeroDur3.Valid }}
C;X58;K{{ .HeroDur3.String }}
{{- end }}
{{- if .Cool3.Valid }}
C;X59;K{{ .Cool3.String }}
{{- end }}
{{- if .Cost3.Valid }}
C;X60;K{{ .Cost3.String }}
{{- end }}
{{- if .Area3.Valid }}
C;X61;K{{ .Area3.String }}
{{- end }}
{{- if .Rng3.Valid }}
C;X62;K{{ .Rng3.String }}
{{- end }}
{{- if .DataA3.Valid }}
C;X63;K{{ .DataA3.String }}
{{- end }}
{{- if .DataB3.Valid }}
C;X64;K{{ .DataB3.String }}
{{- end }}
{{- if .DataC3.Valid }}
C;X65;K{{ .DataC3.String }}
{{- end }}
{{- if .DataD3.Valid }}
C;X66;K{{ .DataD3.String }}
{{- end }}
{{- if .DataE3.Valid }}
C;X67;K{{ .DataE3.String }}
{{- end }}
{{- if .DataF3.Valid }}
C;X68;K{{ .DataF3.String }}
{{- end }}
{{- if .DataG3.Valid }}
C;X69;K{{ .DataG3.String }}
{{- end }}
{{- if .DataH3.Valid }}
C;X70;K{{ .DataH3.String }}
{{- end }}
{{- if .DataI3.Valid }}
C;X71;K{{ .DataI3.String }}
{{- end }}
{{- if .UnitID3.Valid }}
C;X72;K{{ .UnitID3.String }}
{{- end }}
{{- if .BuffID3.Valid }}
C;X73;K{{ .BuffID3.String }}
{{- end }}
{{- if .EfctID3.Valid }}
C;X74;K{{ .EfctID3.String }}
{{- end }}
{{- if .Targs4.Valid }}
C;X75;K{{ .Targs4.String }}
{{- end }}
{{- if .Cast4.Valid }}
C;X76;K{{ .Cast4.String }}
{{- end }}
{{- if .Dur4.Valid }}
C;X77;K{{ .Dur4.String }}
{{- end }}
{{- if .HeroDur4.Valid }}
C;X78;K{{ .HeroDur4.String }}
{{- end }}
{{- if .Cool4.Valid }}
C;X79;K{{ .Cool4.String }}
{{- end }}
{{- if .Cost4.Valid }}
C;X80;K{{ .Cost4.String }}
{{- end }}
{{- if .Area4.Valid }}
C;X81;K{{ .Area4.String }}
{{- end }}
{{- if .Rng4.Valid }}
C;X82;K{{ .Rng4.String }}
{{- end }}
{{- if .DataA4.Valid }}
C;X83;K{{ .DataA4.String }}
{{- end }}
{{- if .DataB4.Valid }}
C;X84;K{{ .DataB4.String }}
{{- end }}
{{- if .DataC4.Valid }}
C;X85;K{{ .DataC4.String }}
{{- end }}
{{- if .DataD4.Valid }}
C;X86;K{{ .DataD4.String }}
{{- end }}
{{- if .DataE4.Valid }}
C;X87;K{{ .DataE4.String }}
{{- end }}
{{- if .DataF4.Valid }}
C;X88;K{{ .DataF4.String }}
{{- end }}
{{- if .DataG4.Valid }}
C;X89;K{{ .DataG4.String }}
{{- end }}
{{- if .DataH4.Valid }}
C;X90;K{{ .DataH4.String }}
{{- end }}
{{- if .DataI4.Valid }}
C;X91;K{{ .DataI4.String }}
{{- end }}
{{- if .UnitID4.Valid }}
C;X92;K{{ .UnitID4.String }}
{{- end }}
{{- if .BuffID4.Valid }}
C;X93;K{{ .BuffID4.String }}
{{- end }}
{{- if .EfctID4.Valid }}
C;X94;K{{ .EfctID4.String }}
{{- end }}
{{- if .InBeta.Valid }}
C;X95;K{{ .InBeta.String }}
{{- end }}
{{- end }}
{{- end }}
E
{{- end }}`
}
