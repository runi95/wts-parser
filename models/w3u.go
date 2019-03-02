package models

import (
	"gopkg.in/volatiletech/null.v6"
)

// Unit data
type W3uData struct {
	BaseUnitId   string
	CustomUnitId string

	Udaa null.String // Default Active Ability
	Uhab null.String // Hero Abilities
	Uabi null.String // Abilities
	Utcc null.String // Allow Custom Team Color
	Uble null.String // Blend Time
	Ucbs null.String // Cast Backswing
	Ucpt null.String // Cast Point
	Urun null.String // Run Speed
	Uwal null.String // Walk Speed
	Ubpx null.String // Button PositionX
	Ubpy null.String // Button PositionY
	Udtm null.String // Death Time
	Uept null.String // Elevation Sample Points
	Uerd null.String // Elevation Sample Radius
	Ufrd null.String // Fog Of War Sample Radius
	Uubs null.String // Ground Texture
	Ushr null.String // Has Water Shadow
	Uico null.String // Icon
	Ussi null.String // Score Screen Icon
	Umxp null.String // Max Pitch
	Umxr null.String // Max Roll
	Umdl null.String // Model
	Uver null.String // Model Extra Versions
	Uocc null.String // Occulder Height
	Uori null.String // Orientation Interpolation
	Uisz null.String // Swim Projectile ImpactZ
	Uimz null.String // Projectile ImpactZ
	Ulpx null.String // Projectile LaunchX
	Ulsz null.String // Swim Projectile LaunchZ
	Ulpz null.String // Projectile LaunchZ
	Uprw null.String // Propulsion Window
	Uani null.String // Required Animation Names
	Uaap null.String // Required Animation Attachments
	Ualp null.String // Required Animation Link Names
	Ubpr null.String // Required Bone Names
	Uscb null.String // Scale Projectiles
	Usca null.String // Scale
	Uslz null.String // Selection Z
	Usew null.String // Selection On Water
	Ussc null.String // Selection Scale
	Ushu null.String // Shadow Image
	Ushx null.String // Shadow Image CenterX
	Ushy null.String // Shadow Image CenterY
	Ushh null.String // Shadow Image Height
	Ushw null.String // Shadow Image Width
	Ushb null.String // Shadow Texture
	Uspa null.String // Special Art
	Utaa null.String // Target Art
	Utco null.String // Team Color
	Uclr null.String // Red Tint
	Uclg null.String // Green Tint
	Uclb null.String // Blue Tint
	Ulos null.String // Use Extended Line Of Sight
	Uacq null.String // Acquisition Range
	Uarm null.String // Armor Type
	Ubs1 null.String // Backswing Point1
	Udp1 null.String // Damage Point1
	Ua1f null.String // Area Of Effect Full1
	Ua1h null.String // Area Of Effect Medium1
	Ua1q null.String // Area Of Effect Small1
	Ua1p null.String // Area Of Effect Targets1
	Ua1t null.String // Attack Type1
	Ua1c null.String // Cooldown 1
	Ua1b null.String // Damage Base1
	Uhd1 null.String // Damage Factor Medium1
	Uqd1 null.String // Damage Factor Small1
	Udl1 null.String // Damage Loss Factor1
	Ua1d null.String // Damage Number Of Dice1
	Ua1s null.String // Damage Sides Per Die1
	Usd1 null.String // Damage Spill Distance1
	Usr1 null.String // Damage Spill Radius1
	Udu1 null.String // Damage Upgrade Amount1
	Utc1 null.String // Maximum Targets1
	Uma1 null.String // Projectile Arc1
	Ua1m null.String // Projectile Art1
	Umh1 null.String // Projectile Homing1
	Ua1z null.String // Projectile Speed1
	Ua1r null.String // Range 1
	Urb1 null.String // Range Motion Buffer1
	Uwu1 null.String // Show UI1
	Ua1g null.String // Targets Allowed1
	Ucs1 null.String // Weapon Sound1
	Ua1w null.String // Weapon Type1
	Ubs2 null.String // Backswing Point2
	Udp2 null.String // Damage Point2
	Ua2f null.String // Area Of Effect Full2
	Ua2h null.String // Area Of Effect Medium2
	Ua2q null.String // Area Of Effect Small2
	Ua2p null.String // Area Of Effect Targets2
	Ua2t null.String // Attack Type2
	Ua2c null.String // Cooldown 2
	Ua2b null.String // Damage Base2
	Uhd2 null.String // Damage Factor Medium2
	Uqd2 null.String // Damage Factor Small2
	Udl2 null.String // Damage Loss Factor2
	Ua2d null.String // Damage Number Of Dice2
	Ua2s null.String // Damage Sides Per Die2
	Usd2 null.String // Damage Spill Distance2
	Usr2 null.String // Damage Spill Radius2
	Udu2 null.String // Damage Upgrade Amount2
	Utc2 null.String // Maximum Targets2
	Uma2 null.String // Projectile Arc2
	Ua2m null.String // Projectile Art2
	Umh2 null.String // Projectile Homing2
	Ua2z null.String // Projectile Speed2
	Ua2r null.String // Range 2
	Urb2 null.String // Range Motion Buffer2
	Uwu2 null.String // Show UI2
	Ua2g null.String // Targets Allowed2
	Ucs2 null.String // Weapon Sound2
	Ua2w null.String // Weapon Type2
	Uaen null.String // Attacks Enabled
	Udea null.String // Death Type
	Udef null.String // Defense Base
	Udty null.String // Defense Type
	Udup null.String // Defense Upgrade Bonus
	Uamn null.String // Minimum Attack Range
	Utar null.String // Targeted As
	Udro null.String // Drop Items On Death
	Ucam null.String // Category Campaign
	Uspe null.String // Category Special
	Uhos null.String // Display As Neutral Hostile
	Utss null.String // Has Tileset Specific Data
	Uine null.String // Placeable In Editor
	Util null.String // Tilesets
	Uuch null.String // Use Click Helper
	Urpo null.String // Group Separation Enabled
	Urpg null.String // Group Separation Group Number
	Urpp null.String // Group Separation Parameter
	Urpr null.String // Group Separation Priority
	Umvh null.String // Fly Height
	Umvf null.String // Minimum Height
	Umvs null.String // Speed Base
	Umas null.String // Speed Maximum
	Umis null.String // Speed Minimum
	Umvr null.String // Turn Rate
	Umvt null.String // Move Type
	Uabr null.String // A IPlacement Radius
	Uabt null.String // A IPlacement Type
	Ucol null.String // Collision Size
	Upat null.String // Pathing Map
	Upar null.String // Placement Prevented By
	Upap null.String // Placement Requires
	Upaw null.String // Placement Requires Water Radius
	Ubsl null.String // Build Sound
	Ulfi null.String // Sound Loop Fade In Rate
	Ulfo null.String // Sound Loop Fade Out Rate
	Umsl null.String // Move Sound
	Ursl null.String // Random Sound
	Usnd null.String // Sound Set
	Uagp null.String // Agility Per Level
	Ubld null.String // Build Time
	Uibo null.String // Can Be Built On
	Ucbo null.String // Can Build On
	Ufle null.String // Can Flee
	Ufoo null.String // Food Cost
	Ufma null.String // Food Produced
	Ufor null.String // Formation Rank
	Ubba null.String // Gold Bounty Base
	Ubdi null.String // Gold Bounty Number Of Dice
	Ubsi null.String // Gold Bounty Sides Per Die
	Ugol null.String // Gold Cost
	Uhhd null.String // Hide Hero Death Message
	Uhhb null.String // Hide Hero Interface Icon
	Uhhm null.String // Hide Hero Minimap Display
	Uhom null.String // Hide Minimap Display
	Uhpm null.String // Hit Points Maximum
	Uhpr null.String // Hit Points Regeneration
	Uhrt null.String // Hit Points Regeneration Type
	Uinp null.String // Intelligence Per Level
	Ubdg null.String // Is ABuilding
	Ulev null.String // Level
	Ulba null.String // Lumber Bounty Base
	Ulbd null.String // Lumber Bounty Number Of Dice
	Ulbs null.String // Lumber Bounty Sides Per Die
	Ulum null.String // Lumber Cost
	Umpi null.String // Mana Initial Amount
	Umpm null.String // Mana Maximum
	Umpr null.String // Mana Regeneration
	Unbm null.String // Show Neutral Building Icon
	Unbr null.String // Valid As Random Neutral Building
	Upoi null.String // Point Value
	Upra null.String // Primary Attribute
	Upri null.String // Priority
	Urac null.String // Race
	Ugor null.String // Repair Gold Cost
	Ulur null.String // Repair Lumber Cost
	Urtm null.String // Repair Time
	Usid null.String // Sight Radius Day
	Usin null.String // Sight Radius Night
	Usle null.String // Sleeps
	Uagi null.String // Starting Agility
	Uint null.String // Starting Intelligence
	Ustr null.String // Starting Strength
	Usma null.String // Stock Maximum
	Usrg null.String // Stock Replenish Interval
	Usst null.String // Stock Start Delay
	Ustp null.String // Strength Per Level
	Ucar null.String // Transported Size
	Utyp null.String // Unit Classification
	Udep null.String // Dependency Equivalents
	Urva null.String // Hero Revival Locations
	Umki null.String // Items Made
	Usei null.String // Items Sold
	Ureq null.String // Requirements
	Urqa null.String // Requirements Levels
	Urq1 null.String // Requirements Tier2
	Urq2 null.String // Requirements Tier3
	Urq3 null.String // Requirements Tier4
	Urq4 null.String // Requirements Tier5
	Urq5 null.String // Requirements Tier6
	Urq6 null.String // Requirements Tier7
	Urq7 null.String // Requirements Tier8
	Urq8 null.String // Requirements Tier9
	Urqc null.String // Requirements Tiers Used
	Ubui null.String // Structures Built
	Ures null.String // Researches Available
	Urev null.String // Revives Dead Heroes
	Useu null.String // Units Sold
	Utra null.String // Units Trained
	Uupt null.String // Upgrades To
	Upgr null.String // Upgrades Used
	Ides null.String // Description
	Uhot null.String // Hotkey
	Unam null.String // Name
	Unsf null.String // Name Editor Suffix
	Upro null.String // Proper Names
	Upru null.String // Proper Names Used
	Uawt null.String // Awaken Tooltip
	Utip null.String // Tooltip
	Utub null.String // Ubertip
	Utpr null.String // Revive Tooltip

	// Unknown
	Ucun null.String
	Ucut null.String
}

// Item data
type W3iData struct {
	Iabi null.String
	Iarm null.String
	Iico null.String
	Ubpx null.String
	Ubpy null.String
	Icla null.String
	Iclb null.String
	Iclg null.String
	Iclr null.String
	Icid null.String
	Ides null.String
	Idrp null.String
	Idro null.String
	Ifil null.String
	Igol null.String
	Uhot null.String
	Ihtp null.String
	Iicd null.String
	Ilev null.String
	Ilum null.String
	Imor null.String
	Unam null.String
	Ilvo null.String
	Ipaw null.String
	Iper null.String
	Iprn null.String
	Ipow null.String
	Ipri null.String
	Ureq null.String
	Urqa null.String
	Isca null.String
	Isel null.String
	Issc null.String
	Isto null.String
	Istr null.String
	Isst null.String
	Utip null.String
	Utub null.String
	Iusa null.String
	Iuse null.String
}

// Destructible data
type W3dData struct {
	Bnam null.String
	Bsuf null.String
	Bcat null.String
	Btil null.String
	Btsp null.String
	Bfil null.String
	Blit null.String
	Bflo null.String
	Btxi null.String
	Btxf null.String
	Buch null.String
	Bonc null.String
	Bonw null.String
	Bcpd null.String
	Bwal null.String
	Bclh null.String
	Btar null.String
	Barm null.String
	Bvar null.String
	Bhps null.String
	Boch null.String
	Bflh null.String
	Bfxr null.String
	Bsel null.String
	Bmis null.String
	Bmas null.String
	Bcpr null.String
	Bmap null.String
	Bmar null.String
	Brad null.String
	Bfra null.String
	Bfvi null.String
	Bptx null.String
	Bptd null.String
	Bdsn null.String
	Bshd null.String
	Bsmm null.String
	Bmmr null.String
	Bmmg null.String
	Bmmb null.String
	Bumm null.String
	Bbut null.String
	Bret null.String
	Breg null.String
	Brel null.String
	Busr null.String
	Bvcr null.String
	Bvcg null.String
	Bvcb null.String
	Bgse null.String
	Bgsc null.String
	Bgpm null.String
}

func (w3uData *W3uData) TransformToSLKUnit(SLKUnit *SLKUnit) {
	unitData := SLKUnit.UnitData
	unitUI := SLKUnit.UnitUI
	unitBalance := SLKUnit.UnitBalance
	unitWeapons := SLKUnit.UnitWeapons
	unitAbilities := SLKUnit.UnitAbilities

	unitData.UnitID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	unitData.Sort.SetValid("\"a1\"") // TODO: Figure this out
	unitData.Comment.SetValid("") // TODO: Figure this out
	if w3uData.Urac.Valid {
		unitData.Race.SetValid("\"" + w3uData.Urac.String + "\"")
	}
	unitData.Prio.SetValid("9") // TODO: Figure this out
	unitData.Threat.SetValid("1") // TODO: Figure this out
	unitData.Valid.SetValid("1") // ?
	unitData.DeathType.SetValid("2") // ?
	unitData.Death.SetValid("3") // ?
	unitData.CanSleep.SetValid("0")
	unitData.CargoSize.SetValid("1")
	if w3uData.Umvt.Valid {
		unitData.Movetp.SetValid("\"" + w3uData.Umvt.String + "\"")
	}
	unitData.MoveHeight.SetValid("0") // TODO: Set this value correctly!
	unitData.MoveFloor.SetValid("0") // ?
	if w3uData.Umvr.Valid {
		unitData.TurnRate.SetValid(w3uData.Umvr.String)
	}
	unitData.PropWin.SetValid("60") // ?
	unitData.OrientInterp.SetValid("4") // ?
	unitData.Formation.SetValid("2") // ?
	if w3uData.Utar.Valid {
		unitData.TargType.SetValid("\"" + w3uData.Utar.String + "\"")
	}
	unitData.PathTex.SetValid("\"_\"") // ?
	unitData.FatLOS.SetValid("0") // ?
	unitData.Points.SetValid("100") // TODO: Set this value correctly!
	unitData.BuffType.SetValid("\"_\"") // ?
	unitData.BuffRadius.SetValid("\"-\"") // ?
	unitData.NameCount.SetValid("\"-\"") // ?
	unitData.CanFlee.SetValid("1") // TODO: Set this value correctly!
	unitData.RequireWaterRadius.SetValid("0")
	unitData.IsBuildOn.SetValid("0") // // TODO: Set this value correctly!
	unitData.CanBuildOn.SetValid("0") // TODO: Set this value correctly!
	unitData.InBeta.SetValid("0") // ?
	unitData.Version.SetValid("0") // ?

	unitUI.UnitUIID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	unitUI.SortUI.SetValid("\"a1\"") // TODO: Figure this out
	if w3uData.Umdl.Valid {
		unitUI.File.SetValid(w3uData.Umdl.String)
	}
	unitUI.FileVerFlags.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.TilesetSpecific.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.UnitClass.SetValid("\"UUnit03\"") // TODO: Figure out what this is supposed to be!
	unitUI.Special.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.Campaign.SetValid("0") // Doesn't really matter much
	unitUI.InEditor.SetValid("") // Doesn't really matter much
	unitUI.HiddenInEditor.SetValid("") // Doesn't really matter much
	unitUI.HostilePal.SetValid("\"-\"")// Don't know what this is, defaults to -
	unitUI.DropItems.SetValid("1") // // Don't know what this is, defaults to 1
	unitUI.NbmmIcon.SetValid("\"-\"") // Don't know what this is, defaults to -
	unitUI.UseClickHelper.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.HideHeroBar.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.HideHeroMinimap.SetValid("0") // Doesn't really matter much
	unitUI.HideHeroDeathMsg.SetValid("0") // Doesn't really matter much
	unitUI.Blend.SetValid("0,15") // Don't know what this is, defaults to 0,15
	unitUI.Scale.SetValid("1") // TODO: Set this value correctly!
	unitUI.ScaleBull.SetValid("1") // Don't know what this is, defaults to 1
	unitUI.MaxPitch.SetValid("45") // TODO: Figure this one out
	unitUI.MaxRoll.SetValid("10") // TODO: Figure this one out
	unitUI.ElevPts.SetValid("\"-\"") // Don't know what this is, defaults to -
	unitUI.ElevRad.SetValid("30") // Don't know what this is, defaults to 30
	unitUI.FogRad.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.Walk.SetValid("260") // TODO: Set this value correctly!
	unitUI.Run.SetValid("260") // TODO: Set this value correctly!
	unitUI.SelZ.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.Weap1.SetValid("\"_\"") // Doesn't really matter much
	unitUI.Weap2.SetValid("\"_\"") // Doesn't really matter much
	unitUI.TeamColor.SetValid("-1") // Don't know what this is, defaults to -1
	unitUI.Armor.SetValid("\"Flesh\"") // TODO: Set this value correctly!
	unitUI.ModelScale.SetValid("1") // TODO: Set this value correctly!
	unitUI.Red.SetValid("255") // TODO: Set this value correctly!
	unitUI.Green.SetValid("255") // TODO: Set this value correctly!
	unitUI.Blue.SetValid("255") // TODO: Set this value correctly!
	unitUI.UberSplat.SetValid("\"_\"") // Don't know what this is, defaults to _
	unitUI.UnitShadow.SetValid("\"_\"") // TODO: Set this value correctly (it can be _, Shadow, ShadowFlyer...)
	unitUI.BuildingShadow.SetValid("\"_\"") // TODO: Set this value correctly!
	unitUI.ShadowW.SetValid("190") // TODO: Set this value correctly!
	unitUI.ShadowH.SetValid("190") // TODO: Set this value correctly!
	unitUI.ShadowX.SetValid("75") // TODO: Set this value correctly!
	unitUI.ShadowY.SetValid("75") // TODO: Set this value correctly!
	unitUI.ShadowOnWater.SetValid("1") // TODO: Set this value correctly!
	unitUI.OccH.SetValid("0") // Don't know what this is, defaults to 0
	unitUI.InBeta.SetValid("") // Don't know what this is, defaults to empty string

	if w3uData.Usnd.Valid {
		unitUI.UnitSound.SetValid("\"" + w3uData.Usnd.String + "\"") // Simple default for now
	} else {
		unitUI.UnitSound.SetValid("\"HeroArchMage\"") // Simple default for now
	}

	if w3uData.Unam.Valid {
		unitUI.Name.SetValid("\"" + w3uData.Unam.String + "\"")
	} else {
		unitUI.Name.SetValid("\"archmage\"")
	}

	unitBalance.UnitBalanceID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	unitBalance.SortBalance.SetValid("\"a1\"")
	unitBalance.Sort2.SetValid("\"uher\"")
	unitBalance.Comment.SetValid("")
	unitBalance.Level.SetValid("5")
	unitBalance.Type.SetValid("\"_\"")
	unitBalance.Goldcost.SetValid("425")
	unitBalance.Lumbercost.SetValid("0")
	unitBalance.GoldRep.SetValid("425")
	unitBalance.LumberRep.SetValid("0")
	unitBalance.Fmade.SetValid("\"-\"")
	unitBalance.Fused.SetValid("0")
	unitBalance.Bountydice.SetValid("8")
	unitBalance.Bountysides.SetValid("3")
	unitBalance.Bountyplus.SetValid("30")
	unitBalance.Lumberbountydice.SetValid("0")
	unitBalance.Lumberbountysides.SetValid("0")
	unitBalance.Lumberbountyplus.SetValid("0")
	unitBalance.StockMax.SetValid("3")
	unitBalance.StockRegen.SetValid("30")
	unitBalance.StockStart.SetValid("0")
	unitBalance.HP.SetValid("100")
	unitBalance.RealHP.SetValid("450")
	unitBalance.RegenHP.SetValid("0,25")
	unitBalance.RegenType.SetValid("\"always\"")
	unitBalance.ManaN.SetValid("0")
	unitBalance.RealM.SetValid("285")
	unitBalance.Mana0.SetValid("100")
	unitBalance.RegenMana.SetValid("0,01")
	unitBalance.Def.SetValid("0")
	unitBalance.DefUp.SetValid("0")
	unitBalance.Realdef.SetValid("3,1")
	unitBalance.DefType.SetValid("\"hero\"")
	unitBalance.Spd.SetValid("320")
	unitBalance.MinSpd.SetValid("0")
	unitBalance.MaxSpd.SetValid("0")
	unitBalance.Bldtm.SetValid("55")
	unitBalance.Reptm.SetValid("55")
	unitBalance.Sight.SetValid("1800")
	unitBalance.Nsight.SetValid("800")
	unitBalance.STR.SetValid("\"-\"")
	unitBalance.INT.SetValid("\"-\"")
	unitBalance.AGI.SetValid("\"-\"")
	unitBalance.STRplus.SetValid("\"-\"")
	unitBalance.INTplus.SetValid("\"-\"")
	unitBalance.AGIplus.SetValid("\"-\"")
	unitBalance.AbilTest.SetValid("\"-\"")
	unitBalance.Primary.SetValid("\"_\"")
	unitBalance.Upgrades.SetValid("\"_\"")
	unitBalance.Tilesets.SetValid("*")
	unitBalance.Nbrandom.SetValid("\"-\"")
	unitBalance.Isbldg.SetValid("0")
	unitBalance.PreventPlace.SetValid("\"_\"")
	unitBalance.RequirePlace.SetValid("\"_\"")
	unitBalance.Repulse.SetValid("0")
	unitBalance.RepulseParam.SetValid("0")
	unitBalance.RepulseGroup.SetValid("0")
	unitBalance.RepulsePrio.SetValid("0")
	unitBalance.Collision.SetValid("32")
	unitBalance.InBeta.SetValid("0")

	unitWeapons.UnitWeapID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	unitWeapons.SortWeap.SetValid("\"a1\"") // ?
	unitWeapons.Sort2.SetValid("\"uher\"") // ?
	unitWeapons.Comment.SetValid("") // ?
	unitWeapons.WeapsOn.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.Acquire.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.MinRange.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Castpt.SetValid("\"-\"") // ?
	unitWeapons.Castbsw.SetValid("0.51") // ?
	unitWeapons.LaunchX.SetValid("0") // ?
	unitWeapons.LaunchY.SetValid("0") // ?
	unitWeapons.LaunchZ.SetValid("60") // ?
	unitWeapons.LaunchSwimZ.SetValid("0") // ?
	unitWeapons.ImpactZ.SetValid("120") // ?
	unitWeapons.ImpactSwimZ.SetValid("0") // ?
	unitWeapons.WeapType1.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.Targs1.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.ShowUI1.SetValid("1") // ?
	unitWeapons.RangeN1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.RngTst.SetValid("\"-\"") // ?
	unitWeapons.RngBuff1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.AtkType1.SetValid("\"normal\"") // TODO: Set this value correctly!
	unitWeapons.WeapTp1.SetValid("\"missile\"") // TODO: Set this value correctly!
	unitWeapons.Cool1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Mincool1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dice1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Sides1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dmgplus1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.DmgUp1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Mindmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Avgdmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Maxdmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dmgpt1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.BackSw1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Farea1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Harea1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Qarea1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Hfact1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Qfact1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.SplashTargs1.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.TargCount1.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.DamageLoss1.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.SpillDist1.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.SpillRadius1.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.DmgUp1.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.Dmod1.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.DPS.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.WeapType2.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.Targs2.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.ShowUI2.SetValid("1") // ?
	unitWeapons.RangeN2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.RngTst2.SetValid("\"-\"") // ?
	unitWeapons.RngBuff2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.AtkType2.SetValid("\"normal\"") // TODO: Set this value correctly!
	unitWeapons.WeapTp2.SetValid("\"missile\"") // TODO: Set this value correctly!
	unitWeapons.Cool2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Mincool2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dice2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Sides2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dmgplus2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.DmgUp2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Mindmg2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Avgdmg2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Maxdmg2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Dmgpt2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.BackSw2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Farea2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Harea2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Qarea2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Hfact2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.Qfact2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.SplashTargs2.SetValid("\"_\"") // TODO: Set this value correctly!
	unitWeapons.TargCount2.SetValid("\"-\"") // TODO: Set this value correctly!
	unitWeapons.DamageLoss2.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.SpillDist2.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.SpillRadius2.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.InBeta.SetValid("0") // ?

	unitAbilities.UnitAbilID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	unitAbilities.SortAbil.SetValid("")
	unitAbilities.Comment.SetValid("")
	unitAbilities.Auto.SetValid("\"_\"") // We never want to enable auto
	unitAbilities.InBeta.SetValid("")

	if w3uData.Uhab.Valid {
		unitAbilities.HeroAbilList.SetValid("\"" + w3uData.Uhab.String + "\"")
	} else {
		unitAbilities.HeroAbilList.SetValid("") // Our map doesn't have any heroes
	}

	if w3uData.Uabi.Valid {
		unitAbilities.AbilList.SetValid("\"" + w3uData.Uabi.String + "\"")
	} else {
		unitAbilities.AbilList.SetValid("\"_\"")
	}
}
