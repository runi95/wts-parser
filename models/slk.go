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
	UnitUIID        null.String
	SortUI          null.String
	File            null.String
	FileVerFlags    null.String
	UnitSound       null.String
	TilesetSpecific null.String
	// Name             null.String
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
	ElevPts          null.String
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
	ShadowW          null.String
	ShadowH          null.String
	ShadowX          null.String
	ShadowY          null.String
	ShadowOnWater    null.String
	SelCircOnWater   null.String
	OccH             null.String
	// InBeta           null.String
}

type UnitWeapons struct {
	UnitWeapID null.String
	SortWeap   null.String
	// Sort2        null.String
	// Comment      null.String
	WeapsOn      null.String
	Acquire      null.String
	MinRange     null.String
	Castpt       null.String
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
	RangeN1      null.String
	RngTst       null.String
	RngBuff1     null.String
	AtkType1     null.String
	WeapTp1      null.String
	Cool1        null.String
	Mincool1     null.String
	Dice1        null.String
	Sides1       null.String
	Dmgplus1     null.String
	DmgUp1       null.String
	Mindmg1      null.String
	Avgdmg1      null.String
	Maxdmg1      null.String
	Dmgpt1       null.String // But really float64, unless empty string
	BackSw1      null.String // But really float64, unless empty string
	Farea1       null.String
	Harea1       null.String
	Qarea1       null.String
	Hfact1       null.String
	Qfact1       null.String
	SplashTargs1 null.String
	TargCount1   null.String
	DamageLoss1  null.String
	SpillDist1   null.String
	SpillRadius1 null.String
	DmgUpg       null.String
	Dmod1        null.String
	DPS          null.String
	WeapType2    null.String
	Targs2       null.String
	ShowUI2      null.String
	RangeN2      null.String
	RngTst2      null.String
	RngBuff2     null.String
	AtkType2     null.String
	WeapTp2      null.String
	Cool2        null.String
	Mincool2     null.String
	Dice2        null.String
	Sides2       null.String
	Dmgplus2     null.String
	DmgUp2       null.String
	Mindmg2      null.String
	Avgdmg2      null.String
	Maxdmg2      null.String
	Dmgpt2       null.String // But really float64, unless empty string
	BackSw2      null.String // But really float64, unless empty string
	Farea2       null.String
	Harea2       null.String
	Qarea2       null.String
	Hfact2       null.String
	Qfact2       null.String
	SplashTargs2 null.String
	TargCount2   null.String
	DamageLoss2  null.String
	SpillDist2   null.String
	SpillRadius2 null.String
	InBeta       null.String
}

type UnitBalance struct {
	UnitBalanceID null.String
	SortBalance   null.String
	Sort2         null.String
	// Comment           null.String
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
	// InBeta            null.String
}

type UnitAbilities struct {
	UnitAbilID null.String
	SortAbil   null.String
	// Comment      null.String
	Auto         null.String
	AbilList     null.String
	HeroAbilList null.String
	// InBeta       null.String
}

type SLKUnit struct {
	*UnitUI
	*UnitData
	*UnitBalance
	*UnitWeapons
	*UnitAbilities
	*UnitFunc
	*UnitString
}

type ItemData struct {
	ItemID     null.String
	Comment    null.String
	Scriptname null.String
	Version    null.String
	Class      null.String
	Level      null.String
	OldLevel   null.String
	AbilList   null.String
	CooldownID null.String
	IgnoreCD   null.String
	Uses       null.String
	Prio       null.String
	Usable     null.String
	Perishable null.String
	Droppable  null.String
	Pawnable   null.String
	Sellable   null.String
	PickRandom null.String
	Powerup    null.String
	Drop       null.String
	StockMax   null.String
	StockRegen null.String
	StockStart null.String
	Goldcost   null.String
	Lumbercost null.String
	HP         null.String
	Morph      null.String
	Armor      null.String
	File       null.String
	Scale      null.String
	SelSize    null.String
	ColorR     null.String
	ColorG     null.String
	ColorB     null.String
	InBeta     null.String
}

type SLKItem struct {
	*ItemData
	*ItemFunc
	*ItemString
}

type AbilityMetaData struct {
	ID          null.String
	Field       null.String
	Slk         null.String
	Index       null.String
	Repeat      null.String
	Data        null.String
	Category    null.String
	DisplayName null.String
	Sort        null.String
	Type        null.String
	ChangeFlags null.String
	ImportType  null.String
	StringExt   null.String
	CaseSens    null.String
	CanBeEmpty  null.String
	MinVal      null.String
	MaxVal      null.String
	ForceNonNeg null.String
	UseUnit     null.String
	UseHero     null.String
	UseItem     null.String
	UseCreep    null.String
	UseSpecific null.String
	NotSpecific null.String
	Version     null.String
	Section     null.String
}

type AbilityData struct {
	Alias       null.String
	Code        null.String
	Comments    null.String
	Version     null.String
	UseInEditor null.String
	Hero        null.String
	Item        null.String
	Sort        null.String
	Race        null.String
	CheckDep    null.String
	Levels      null.String
	ReqLevel    null.String
	LevelSkip   null.String
	Priority    null.String
	Targs1      null.String
	Cast1       null.String
	Dur1        null.String
	HeroDur1    null.String
	Cool1       null.String
	Cost1       null.String
	Area1       null.String
	Rng1        null.String
	DataA1      null.String
	DataB1      null.String
	DataC1      null.String
	DataD1      null.String
	DataE1      null.String
	DataF1      null.String
	DataG1      null.String
	DataH1      null.String
	DataI1      null.String
	UnitID1     null.String
	BuffID1     null.String
	EfctID1     null.String
	Targs2      null.String
	Cast2       null.String
	Dur2        null.String
	HeroDur2    null.String
	Cool2       null.String
	Cost2       null.String
	Area2       null.String
	Rng2        null.String
	DataA2      null.String
	DataB2      null.String
	DataC2      null.String
	DataD2      null.String
	DataE2      null.String
	DataF2      null.String
	DataG2      null.String
	DataH2      null.String
	DataI2      null.String
	UnitID2     null.String
	BuffID2     null.String
	EfctID2     null.String
	Targs3      null.String
	Cast3       null.String
	Dur3        null.String
	HeroDur3    null.String
	Cool3       null.String
	Cost3       null.String
	Area3       null.String
	Rng3        null.String
	DataA3      null.String
	DataB3      null.String
	DataC3      null.String
	DataD3      null.String
	DataE3      null.String
	DataF3      null.String
	DataG3      null.String
	DataH3      null.String
	DataI3      null.String
	UnitID3     null.String
	BuffID3     null.String
	EfctID3     null.String
	Targs4      null.String
	Cast4       null.String
	Dur4        null.String
	HeroDur4    null.String
	Cool4       null.String
	Cost4       null.String
	Area4       null.String
	Rng4        null.String
	DataA4      null.String
	DataB4      null.String
	DataC4      null.String
	DataD4      null.String
	DataE4      null.String
	DataF4      null.String
	DataG4      null.String
	DataH4      null.String
	DataI4      null.String
	UnitID4     null.String
	BuffID4     null.String
	EfctID4     null.String
	InBeta      null.String
}

type LevelDependentData struct {
	Targs   null.String
	Cast    null.String
	Dur     null.String
	HeroDur null.String
	Cool    null.String
	Cost    null.String
	Area    null.String
	Rng     null.String
	DataA   null.String
	DataB   null.String
	DataC   null.String
	DataD   null.String
	DataE   null.String
	DataF   null.String
	DataG   null.String
	DataH   null.String
	DataI   null.String
	UnitID  null.String
	BuffID  null.String
	EfctID  null.String
}

type SLKAbility struct {
	*AbilityData
	*AbilityFunc
	*AbilityString
}
