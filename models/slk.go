package models

import "gopkg.in/volatiletech/null.v6"

type UnitData struct {
	UnitID             null.String
	Sort               null.String
	Comment            null.String
	Race               null.String
	Prio               null.String
	Threat             null.String
	Valid              null.String
	DeathType          null.String
	Death              null.String
	CanSleep           null.String
	CargoSize          null.String
	Movetp             null.String
	MoveHeight         null.String
	MoveFloor          null.String
	TurnRate           null.String
	PropWin            null.String
	OrientInterp       null.String
	Formation          null.String
	TargType           null.String
	PathTex            null.String
	FatLOS             null.String
	Points             null.String
	BuffType           null.String
	BuffRadius         null.String
	NameCount          null.String
	CanFlee            null.String
	RequireWaterRadius null.String
	IsBuildOn          null.String
	CanBuildOn         null.String
	InBeta             null.String
	Version            null.String
}

type UnitUI struct {
	UnitUIID         null.String
	SortUI           null.String
	File             null.String
	FileVerFlags     null.String
	UnitSound        null.String
	TilesetSpecific  null.String
	Name             null.String
	UnitClass        null.String
	Special          null.String
	Campaign         null.String
	InEditor         null.String
	HiddenInEditor   null.String
	HostilePal       null.String
	DropItems        null.String
	NbmmIcon         null.String
	UseClickHelper   null.String
	HideHeroBar      null.String
	HideHeroMinimap  null.String
	HideHeroDeathMsg null.String
	HideOnMinimap    null.String
	Blend            null.String
	Scale            null.String
	ScaleBull        null.String
	MaxPitch         null.String
	MaxRoll          null.String
	ElevPts          null.String // But really int, unless empty string
	ElevRad          null.String
	FogRad           null.String
	Walk             null.String
	Run              null.String
	SelZ             null.String
	Weap1            null.String
	Weap2            null.String
	TeamColor        null.String
	CustomTeamColor  null.String
	Armor            null.String
	ModelScale       null.String
	Red              null.String
	Green            null.String
	Blue             null.String
	UberSplat        null.String
	UnitShadow       null.String
	BuildingShadow   null.String
	ShadowW          null.String // But really int, unless empty string
	ShadowH          null.String // But really int, unless empty string
	ShadowX          null.String // But really int, unless empty string
	ShadowY          null.String // But really int, unless empty string
	ShadowOnWater    null.String
	SelCircOnWater   null.String
	OccH             null.String
	InBeta           null.String
}

type UnitWeapons struct {
	UnitWeapID   null.String
	SortWeap     null.String
	Sort2        null.String
	Comment      null.String
	WeapsOn      null.String
	Acquire      null.String // But really int, unless empty string
	MinRange     null.String // But really int, unless empty string
	Castpt       null.String // But really int, unless empty string
	Castbsw      null.String
	LaunchX      null.String
	LaunchY      null.String
	LaunchZ      null.String
	LaunchSwimZ  null.String
	ImpactZ      null.String
	ImpactSwimZ  null.String
	WeapType1    null.String
	Targs1       null.String
	ShowUI1      null.String
	RangeN1      null.String // But really int, unless empty string
	RngTst       null.String
	RngBuff1     null.String // But really int, unless empty string
	AtkType1     null.String
	WeapTp1      null.String
	Cool1        null.String // But really int, unless empty string
	Mincool1     null.String
	Dice1        null.String // But really int, unless empty string
	Sides1       null.String // But really int, unless empty string
	Dmgplus1     null.String // But really int, unless empty string
	DmgUp1       null.String
	Mindmg1      null.String // But really int, unless empty string
	Avgdmg1      null.String // But really int, unless empty string
	Maxdmg1      null.String // But really int, unless empty string
	Dmgpt1       null.String // But really float64, unless empty string
	BackSw1      null.String // But really float64, unless empty string
	Farea1       null.String // But really int, unless empty string
	Harea1       null.String // But really int, unless empty string
	Qarea1       null.String // But really int, unless empty string
	Hfact1       null.String
	Qfact1       null.String
	SplashTargs1 null.String
	TargCount1   null.String // But really int, unless empty string
	DamageLoss1  null.String
	SpillDist1   null.String
	SpillRadius1 null.String
	DmgUpg       null.String
	Dmod1        null.String // But really int, unless empty string
	DPS          null.String
	WeapType2    null.String
	Targs2       null.String
	ShowUI2      null.String
	RangeN2      null.String // But really int, unless empty string
	RngTst2      null.String
	RngBuff2     null.String // But really int, unless empty string
	AtkType2     null.String
	WeapTp2      null.String
	Cool2        null.String // But really int, unless empty string
	Mincool2     null.String
	Dice2        null.String // But really int, unless empty string
	Sides2       null.String // But really int, unless empty string
	Dmgplus2     null.String // But really int, unless empty string
	DmgUp2       null.String
	Mindmg2      null.String // But really int, unless empty string
	Avgdmg2      null.String // But really int, unless empty string
	Maxdmg2      null.String // But really int, unless empty string
	Dmgpt2       null.String // But really float64, unless empty string
	BackSw2      null.String // But really float64, unless empty string
	Farea2       null.String // But really int, unless empty string
	Harea2       null.String // But really int, unless empty string
	Qarea2       null.String // But really int, unless empty string
	Hfact2       null.String
	Qfact2       null.String
	SplashTargs2 null.String
	TargCount2   null.String // But really int, unless empty string
	DamageLoss2  null.String
	SpillDist2   null.String
	SpillRadius2 null.String
	InBeta       null.String
}

type UnitBalance struct {
	UnitBalanceID     null.String
	SortBalance       null.String
	Sort2             null.String
	Comment           null.String
	Level             null.String
	Type              null.String
	Goldcost          null.String
	Lumbercost        null.String
	GoldRep           null.String
	LumberRep         null.String
	Fmade             null.String
	Fused             null.String
	Bountydice        null.String
	Bountysides       null.String
	Bountyplus        null.String
	Lumberbountydice  null.String
	Lumberbountysides null.String
	Lumberbountyplus  null.String
	StockMax          null.String
	StockRegen        null.String
	StockStart        null.String
	HP                null.String
	RealHP            null.String
	RegenHP           null.String
	RegenType         null.String
	ManaN             null.String
	RealM             null.String
	Mana0             null.String
	RegenMana         null.String
	Def               null.String
	DefUp             null.String
	Realdef           null.String
	DefType           null.String
	Spd               null.String
	MinSpd            null.String
	MaxSpd            null.String
	Bldtm             null.String
	Reptm             null.String
	Sight             null.String
	Nsight            null.String
	STR               null.String
	INT               null.String
	AGI               null.String
	STRplus           null.String
	INTplus           null.String
	AGIplus           null.String
	AbilTest          null.String
	Primary           null.String
	Upgrades          null.String
	Tilesets          null.String
	Nbrandom          null.String
	Isbldg            null.String
	PreventPlace      null.String
	RequirePlace      null.String
	Repulse           null.String
	RepulseParam      null.String
	RepulseGroup      null.String
	RepulsePrio       null.String
	Collision         null.String
	InBeta            null.String
}

type UnitAbilities struct {
	UnitAbilID   null.String
	SortAbil     null.String
	Comment      null.String
	Auto         null.String
	AbilList     null.String
	HeroAbilList null.String
	InBeta       null.String
}

type SLKUnit struct {
	UnitUI        *UnitUI
	UnitData      *UnitData
	UnitBalance   *UnitBalance
	UnitWeapons   *UnitWeapons
	UnitAbilities *UnitAbilities
}