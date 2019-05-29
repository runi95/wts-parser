package models

import (
	"gopkg.in/volatiletech/null.v6"
	"strings"
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
	Ulpy null.String // Projectile LaunchY
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
	Ulos null.String // Use Extended Line Of Sight // TODO: Find usage
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
	Udu1 null.String // Damage Upgrade Amount1 // TODO: Find usage
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
	Uabr null.String // A IPlacement Radius // TODO: Find usage
	Uabt null.String // A IPlacement Type // TODO: Find usage
	Ucol null.String // Collision Size
	Upat null.String // Pathing Map
	Upar null.String // Placement Prevented By
	Upap null.String // Placement Requires
	Upaw null.String // Placement Requires Water Radius
	Ubsl null.String // Build Sound
	Ulfi null.String // Sound Loop Fade In Rate // TODO: Find usage
	Ulfo null.String // Sound Loop Fade Out Rate // TODO: Find usage
	Umsl null.String // Move Sound // TODO: Find usage
	Ursl null.String // Random Sound // TODO: Find usage
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
	Ubdg null.String // Is A Building
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
	Udep null.String // Dependency Equivalents // TODO: Find usage
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
	// unitData.Sort.SetValid("\"a1\"") // TODO: Figure this out
	// unitData.Comment.SetValid("")    // TODO: Figure this out
	if w3uData.Urac.Valid {
		unitData.Race.SetValid("\"" + w3uData.Urac.String + "\"")
	}
	if w3uData.Upri.Valid {
		unitData.Prio.SetValid(w3uData.Upri.String)
	}
	// unitData.Threat.SetValid("1")    // TODO: Figure this out
	if w3uData.Unbr.Valid { // Here I assume unitData.Valid = valid as random neutral unit
		unitData.Valid.SetValid(w3uData.Unbr.String)
	}
	if w3uData.Udea.Valid {
		unitData.DeathType.SetValid(w3uData.Udea.String)
	}
	if w3uData.Udtm.Valid { // Here I assume unitData.Death = death time
		unitData.Death.SetValid(w3uData.Udtm.String)
	}
	if w3uData.Usle.Valid {
		unitData.CanSleep.SetValid(w3uData.Usle.String)
	}
	if w3uData.Ucar.Valid {
		unitData.CargoSize.SetValid(w3uData.Ucar.String)
	}
	if w3uData.Umvt.Valid {
		unitData.Movetp.SetValid("\"" + w3uData.Umvt.String + "\"")
	}
	if w3uData.Umvh.Valid {
		unitData.MoveHeight.SetValid(w3uData.Umvh.String)
	}
	if w3uData.Umvf.Valid {
		unitData.MoveFloor.SetValid(w3uData.Umvf.String)
	}
	if w3uData.Umvr.Valid {
		unitData.TurnRate.SetValid(w3uData.Umvr.String)
	}
	if w3uData.Uprw.Valid {
		unitData.PropWin.SetValid(w3uData.Uprw.String)
	}
	if w3uData.Uori.Valid {
		unitData.OrientInterp.SetValid(w3uData.Uori.String)
	}
	if w3uData.Ufor.Valid {
		unitData.Formation.SetValid(w3uData.Ufor.String)
	}
	if w3uData.Utar.Valid {
		unitData.TargType.SetValid("\"" + w3uData.Utar.String + "\"")
	}
	if w3uData.Upat.Valid {
		unitData.PathTex.SetValid(w3uData.Upat.String)
	}
	// unitData.FatLOS.SetValid("0")         // TODO: Set this value correctly!
	if w3uData.Upoi.Valid {
		unitData.Points.SetValid(w3uData.Upoi.String)
	}
	// unitData.BuffType.SetValid("\"_\"")   // TODO: Set this value correctly!
	// unitData.BuffRadius.SetValid("\"-\"") // TODO: Set this value correctly!
	// unitData.NameCount.SetValid("\"-\"")  // TODO: Set this value correctly!
	if w3uData.Ufle.Valid {
		unitData.CanFlee.SetValid(w3uData.Ufle.String)
	}
	if w3uData.Upaw.Valid {
		unitData.RequireWaterRadius.SetValid(w3uData.Upaw.String)
	}
	if w3uData.Uibo.Valid {
		unitData.IsBuildOn.SetValid(w3uData.Uibo.String)
	}
	if w3uData.Ucbo.Valid {
		unitData.CanBuildOn.SetValid(w3uData.Ucbo.String)
	}
	// unitData.InBeta.SetValid("0")  // TODO: Set this value correctly!
	// unitData.Version.SetValid("0") // TODO: Set this value correctly!

	unitUI.UnitUIID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	// unitUI.SortUI.SetValid("\"a1\"") // TODO: Figure this out
	if w3uData.Umdl.Valid {
		unitUI.File.SetValid("\"" + w3uData.Umdl.String + "\"")
	}
	// unitUI.FileVerFlags.SetValid("0")        // TODO: Set this value correctly!
	if w3uData.Utss.Valid {
		unitUI.TilesetSpecific.SetValid(w3uData.Utss.String)
	}
	// unitUI.UnitClass.SetValid("\"UUnit03\"") // TODO: Set this value correctly!
	if w3uData.Uspe.Valid {
		unitUI.Special.SetValid(w3uData.Uspe.String)
	}
	/*
	if w3uData.Ucam.Valid {
		unitUI.Campaign.SetValid(w3uData.Ucam.String)
	}
	*/
	unitUI.Campaign.SetValid("1") // This is permanently set to 1 because we want everything to be campaign
	// unitUI.InEditor.SetValid("")             // TODO: Set this value correctly!
	// unitUI.HiddenInEditor.SetValid("")       // TODO: Set this value correctly!
	// unitUI.HostilePal.SetValid("\"-\"")      // TODO: Set this value correctly!
	if w3uData.Udro.Valid {
		unitUI.DropItems.SetValid(w3uData.Udro.String)
	}
	// unitUI.NbmmIcon.SetValid("\"-\"")        // TODO: Set this value correctly!
	// unitUI.UseClickHelper.SetValid("0")      // TODO: Set this value correctly!
	// unitUI.HideHeroBar.SetValid("0")         // TODO: Set this value correctly!
	if w3uData.Uhhm.Valid {
		unitUI.HideHeroMinimap.SetValid(w3uData.Uhhm.String)
	}
	if w3uData.Uhhd.Valid {
		unitUI.HideHeroDeathMsg.SetValid(w3uData.Uhhd.String)
	}
	if w3uData.Uhom.Valid {
		unitUI.HideOnMinimap.SetValid(w3uData.Uhom.String)
	}
	// unitUI.Blend.SetValid("0,15")            // TODO: Set this value correctly!
	if w3uData.Ussc.Valid {
		unitUI.Scale.SetValid(w3uData.Ussc.String)
	}
	// unitUI.ScaleBull.SetValid("1")   // TODO: Set this value correctly!
	if w3uData.Umxp.Valid {
		unitUI.MaxPitch.SetValid(w3uData.Umxp.String)
	}
	if w3uData.Umxr.Valid {
		unitUI.MaxRoll.SetValid(w3uData.Umxr.String)
	}
	if w3uData.Uept.Valid {
		unitUI.ElevPts.SetValid(w3uData.Uept.String)
	}
	if w3uData.Uerd.Valid {
		unitUI.ElevRad.SetValid(w3uData.Uerd.String)
	}
	if w3uData.Ufrd.Valid {
		unitUI.FogRad.SetValid(w3uData.Ufrd.String)
	}
	if w3uData.Uwal.Valid {
		unitUI.Walk.SetValid(w3uData.Uwal.String)
	}
	if w3uData.Urun.Valid {
		unitUI.Run.SetValid(w3uData.Urun.String)
	}
	// unitUI.SelZ.SetValid("0")        // TODO: Set this value correctly!
	// unitUI.Weap1.SetValid("\"_\"")   // TODO: Set this value correctly! This can be _ or MetalLightChop etc...
	// unitUI.Weap2.SetValid("\"_\"")   // TODO: Set this value correctly! This can be _ or MetalLightChop etc...
	if w3uData.Utcc.Valid {
		unitUI.CustomTeamColor.SetValid(w3uData.Utcc.String)
	}
	if w3uData.Utco.Valid {
		unitUI.TeamColor.SetValid(w3uData.Utco.String)
	}
	if w3uData.Uarm.Valid {
		unitUI.Armor.SetValid(w3uData.Uarm.String)
	}
	if w3uData.Usca.Valid {
		unitUI.ModelScale.SetValid(w3uData.Usca.String)
	}
	if w3uData.Uclr.Valid {
		unitUI.Red.SetValid(w3uData.Uclr.String)
	}
	if w3uData.Uclg.Valid {
		unitUI.Green.SetValid(w3uData.Uclg.String)
	}
	if w3uData.Uclb.Valid {
		unitUI.Blue.SetValid(w3uData.Uclb.String)
	}
	if w3uData.Uubs.Valid {
		unitUI.UberSplat.SetValid(w3uData.Uubs.String)
	}
	if w3uData.Ushu.Valid {
		unitUI.UnitShadow.SetValid(w3uData.Ushu.String)
	}
	if w3uData.Ushb.Valid {
		unitUI.BuildingShadow.SetValid(w3uData.Ushb.String)
	}
	if w3uData.Ushw.Valid {
		unitUI.ShadowW.SetValid(w3uData.Ushw.String)
	}
	if w3uData.Ushh.Valid {
		unitUI.ShadowH.SetValid(w3uData.Ushh.String)
	}
	if w3uData.Ushx.Valid {
		unitUI.ShadowX.SetValid(w3uData.Ushx.String)
	}
	if w3uData.Ushy.Valid {
		unitUI.ShadowY.SetValid(w3uData.Ushy.String)
	}
	if w3uData.Ushr.Valid {
		unitUI.ShadowOnWater.SetValid(w3uData.Ushr.String)
	}
	// unitUI.OccH.SetValid("0")               // TODO: Set this value correctly!
	// unitUI.InBeta.SetValid("")              // TODO: Set this value correctly!
	if w3uData.Usnd.Valid {
		unitUI.UnitSound.SetValid("\"" + w3uData.Usnd.String + "\"")
	}
	/*
	if w3uData.Unam.Valid {
		unitUI.Name.SetValid("\"" + w3uData.Unam.String + "\"")
	}
	*/

	unitBalance.UnitBalanceID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	// unitBalance.SortBalance.SetValid("\"a1\"") // TODO: Set this value correctly!
	// unitBalance.Sort2.SetValid("\"uher\"") // TODO: Set this value correctly!
	// unitBalance.Comment.SetValid("") // TODO: Set this value correctly!
	if w3uData.Ulev.Valid {
		unitBalance.Level.SetValid(w3uData.Ulev.String)
	}
	if w3uData.Utyp.Valid {
		unitBalance.Type.SetValid(w3uData.Utyp.String)
	}
	if w3uData.Ugol.Valid {
		unitBalance.Goldcost.SetValid(w3uData.Ugol.String)
	}
	if w3uData.Ulum.Valid {
		unitBalance.Lumbercost.SetValid(w3uData.Ulum.String)
	}
	if w3uData.Ugor.Valid {
		unitBalance.GoldRep.SetValid(w3uData.Ugor.String)
	}
	if w3uData.Ulur.Valid {
		unitBalance.LumberRep.SetValid(w3uData.Ulur.String)
	}
	if w3uData.Ufma.Valid {
		unitBalance.Fmade.SetValid(w3uData.Ufma.String)
	}
	if w3uData.Ufoo.Valid {
		unitBalance.Fused.SetValid(w3uData.Ufoo.String)
	}
	if w3uData.Ubdi.Valid {
		unitBalance.Bountydice.SetValid(w3uData.Ubdi.String)
	}
	if w3uData.Ubsi.Valid {
		unitBalance.Bountysides.SetValid(w3uData.Ubsi.String)
	}
	if w3uData.Ubba.Valid {
		unitBalance.Bountyplus.SetValid(w3uData.Ubba.String)
	}
	if w3uData.Ulbd.Valid {
		unitBalance.Lumberbountydice.SetValid(w3uData.Ulbd.String)
	}
	if w3uData.Ulbs.Valid {
		unitBalance.Lumberbountysides.SetValid(w3uData.Ulbs.String)
	}
	if w3uData.Ulba.Valid {
		unitBalance.Lumberbountyplus.SetValid(w3uData.Ulba.String)
	}
	if w3uData.Usma.Valid {
		unitBalance.StockMax.SetValid(w3uData.Usma.String)
	}
	if w3uData.Usrg.Valid {
		unitBalance.StockRegen.SetValid(w3uData.Usrg.String)
	}
	if w3uData.Usst.Valid {
		unitBalance.StockStart.SetValid(w3uData.Usst.String)
	}
	if w3uData.Uhpm.Valid {
		unitBalance.HP.SetValid(w3uData.Uhpm.String)
		unitBalance.RealHP.SetValid(w3uData.Uhpm.String) // TODO: This needs be HP + 25 * STR
	}
	// unitBalance.RealHP.SetValid("450") // TODO: Set this value correctly!
	if w3uData.Uhpr.Valid {
		unitBalance.RegenHP.SetValid(w3uData.Uhpr.String)
	}
	if w3uData.Uhrt.Valid {
		unitBalance.RegenType.SetValid(w3uData.Uhrt.String)
	}
	if w3uData.Umpm.Valid {
		unitBalance.ManaN.SetValid(w3uData.Umpm.String)
	}
	// unitBalance.RealM.SetValid("285") // TODO: Calculate this mana amount
	if w3uData.Umpi.Valid {
		unitBalance.Mana0.SetValid(w3uData.Umpi.String)
	}
	if w3uData.Umpr.Valid {
		unitBalance.RegenMana.SetValid(w3uData.Umpr.String)
	}
	if w3uData.Udef.Valid {
		unitBalance.Def.SetValid(w3uData.Udef.String)
		unitBalance.Realdef.SetValid(w3uData.Udef.String) // TODO: Calculate this value by putting base defense + some% * AGI
	}
	if w3uData.Udup.Valid {
		unitBalance.DefUp.SetValid(w3uData.Udup.String)
	}
	// unitBalance.Realdef.SetValid("3,1") // TODO: Set this value correctly!
	if w3uData.Udty.Valid {
		unitBalance.DefType.SetValid(w3uData.Udty.String)
	}
	if w3uData.Umvs.Valid {
		unitBalance.Spd.SetValid(w3uData.Umvs.String)
	}
	if w3uData.Umis.Valid {
		unitBalance.MinSpd.SetValid(w3uData.Umis.String)
	}
	if w3uData.Umas.Valid {
		unitBalance.MaxSpd.SetValid(w3uData.Umas.String)
	}
	if w3uData.Ubld.Valid {
		unitBalance.Bldtm.SetValid(w3uData.Ubld.String)
	}
	if w3uData.Urtm.Valid {
		unitBalance.Reptm.SetValid(w3uData.Urtm.String)
	}
	if w3uData.Usid.Valid {
		unitBalance.Sight.SetValid(w3uData.Usid.String)
	}
	if w3uData.Usin.Valid {
		unitBalance.Nsight.SetValid(w3uData.Usin.String)
	}
	if w3uData.Ustr.Valid {
		unitBalance.STR.SetValid(w3uData.Ustr.String)
	}
	if w3uData.Uint.Valid {
		unitBalance.INT.SetValid(w3uData.Uint.String)
	}
	if w3uData.Uagi.Valid {
		unitBalance.AGI.SetValid(w3uData.Uagi.String)
	}
	if w3uData.Ustp.Valid {
		unitBalance.STRplus.SetValid(w3uData.Ustp.String)
	}
	if w3uData.Uinp.Valid {
		unitBalance.INTplus.SetValid(w3uData.Uinp.String)
	}
	if w3uData.Uagp.Valid {
		unitBalance.AGIplus.SetValid(w3uData.Uagp.String)
	}
	// unitBalance.AbilTest.SetValid("\"-\"") // TODO: Set this value correctly!
	if w3uData.Upra.Valid {
		unitBalance.Primary.SetValid(w3uData.Upra.String)
	}
	// unitBalance.Upgrades.SetValid("\"_\"") // TODO: Set this value correctly!
	if w3uData.Util.Valid {
		unitBalance.Tilesets.SetValid(w3uData.Util.String)
	}
	// unitBalance.Nbrandom.SetValid("\"-\"") // TODO: Set this value correctly!
	if w3uData.Ubdg.Valid {
		unitBalance.Isbldg.SetValid(w3uData.Ubdg.String)
	}
	// TODO: Figure this out, the base SLK files are inverted on this!
	if w3uData.Upar.Valid {
		/*
		split := strings.Split(w3uData.Upar.String, ",")
		for i := range split {
			switch split[i] {
			case "unbuildable":
				split[i] = "buildable"
			case "unwalkable":
				split[i] = "walkable"
			case "buildable":
				split[i] = "unbuildable"
			case "walkable":
				split[i] = "unwalkable"
			}
		}

		newStr := "\""
		for i := range split {
			if i == 0 {
				newStr += split[i]
			} else {
				newStr += "," + split[i]
			}
		}
		newStr += "\""
		*/
		unitBalance.RequirePlace.SetValid(w3uData.Upar.String)
	}
	if w3uData.Upap.Valid {
		/*
		split := strings.Split(w3uData.Upar.String, ",")
		for i := range split {
			switch split[i] {
			case "unbuildable":
				split[i] = "buildable"
			case "unwalkable":
				split[i] = "walkable"
			case "buildable":
				split[i] = "unbuildable"
			case "walkable":
				split[i] = "unwalkable"
			}
		}

		newStr := "\""
		for i := range split {
			if i == 0 {
				newStr += split[i]
			} else {
				newStr += "," + split[i]
			}
		}
		newStr += "\""
		*/
		unitBalance.PreventPlace.SetValid(w3uData.Upap.String)
	}
	if w3uData.Urpo.Valid {
		unitBalance.Repulse.SetValid(w3uData.Urpo.String)
	}
	if w3uData.Urpp.Valid {
		unitBalance.RepulseParam.SetValid(w3uData.Urpp.String)
	}
	if w3uData.Urpg.Valid {
		unitBalance.RepulseGroup.SetValid(w3uData.Urpg.String)
	}
	if w3uData.Urpr.Valid {
		unitBalance.RepulsePrio.SetValid(w3uData.Urpr.String)
	}
	if w3uData.Ucol.Valid {
		unitBalance.Collision.SetValid(w3uData.Ucol.String)
	}
	// unitBalance.InBeta.SetValid("0") // TODO: Set this value correctly!
	unitWeapons.UnitWeapID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	// unitWeapons.SortWeap.SetValid("\"a1\"") // TODO: Set this value correctly!
	// unitWeapons.Sort2.SetValid("\"uher\"") // TODO: Set this value correctly!
	// unitWeapons.Comment.SetValid("") // TODO: Set this value correctly!
	if w3uData.Uaen.Valid {
		unitWeapons.WeapsOn.SetValid(w3uData.Uaen.String)
	}
	if w3uData.Uacq.Valid {
		unitWeapons.Acquire.SetValid(w3uData.Uacq.String)
	}
	if w3uData.Uamn.Valid {
		unitWeapons.MinRange.SetValid(w3uData.Uamn.String)
	}
	if w3uData.Ucpt.Valid {
		unitWeapons.Castpt.SetValid(w3uData.Ucpt.String)
	}
	if w3uData.Ucbs.Valid {
		unitWeapons.Castbsw.SetValid(w3uData.Ucbs.String)
	}
	if w3uData.Ulpx.Valid {
		unitWeapons.LaunchX.SetValid(w3uData.Ulpx.String)
	}
	if w3uData.Ulpy.Valid {
		unitWeapons.LaunchY.SetValid(w3uData.Ulpy.String)
	}
	if w3uData.Ulpz.Valid {
		unitWeapons.LaunchZ.SetValid(w3uData.Ulpz.String)
	}
	if w3uData.Ulsz.Valid {
		unitWeapons.LaunchSwimZ.SetValid(w3uData.Ulsz.String)
	}
	if w3uData.Uimz.Valid {
		unitWeapons.ImpactZ.SetValid(w3uData.Uimz.String)
	}
	if w3uData.Uisz.Valid {
		unitWeapons.ImpactSwimZ.SetValid(w3uData.Uisz.String)
	}
	if w3uData.Ucs1.Valid {
		unitWeapons.WeapType1.SetValid(w3uData.Ucs1.String)
	}
	if w3uData.Ua1g.Valid {
		unitWeapons.Targs1.SetValid("\"" + w3uData.Ua1g.String + "\"")
	}
	if w3uData.Uwu1.Valid {
		unitWeapons.ShowUI1.SetValid(w3uData.Uwu1.String)
	}
	if w3uData.Ua1r.Valid {
		unitWeapons.RangeN1.SetValid(w3uData.Ua1r.String)
	}
	// unitWeapons.RngTst.SetValid("\"-\"") // TODO: Set this value correctly!
	if w3uData.Urb1.Valid {
		unitWeapons.RngBuff1.SetValid(w3uData.Urb1.String)
	}
	if w3uData.Ua1t.Valid {
		unitWeapons.AtkType1.SetValid("\"" + w3uData.Ua1t.String + "\"")
	}
	if w3uData.Ua1w.Valid {
		unitWeapons.WeapTp1.SetValid("\"" + w3uData.Ua1w.String + "\"")
	}
	if w3uData.Ua1c.Valid {
		unitWeapons.Cool1.SetValid(w3uData.Ua1c.String)
	}
	// unitWeapons.Mincool1.SetValid("\"-\"")      // TODO: Set this value correctly! (seems like it's always equal to -)
	if w3uData.Ua1d.Valid {
		unitWeapons.Dice1.SetValid(w3uData.Ua1d.String)
	}
	if w3uData.Ua1s.Valid {
		unitWeapons.Sides1.SetValid(w3uData.Ua1s.String)
	}
	if w3uData.Ua1b.Valid {
		unitWeapons.Dmgplus1.SetValid(w3uData.Ua1b.String)
	}
	// unitWeapons.DmgUp1.SetValid("\"-\"")  // TODO: Set this value correctly! (seems like it's always equal to -)
	// unitWeapons.Mindmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	// unitWeapons.Avgdmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	// unitWeapons.Maxdmg1.SetValid("\"-\"") // TODO: Set this value correctly!
	if w3uData.Udp1.Valid {
		unitWeapons.Dmgpt1.SetValid(w3uData.Udp1.String)
	}
	if w3uData.Ubs1.Valid {
		unitWeapons.BackSw1.SetValid(w3uData.Ubs1.String)
	}
	if w3uData.Ua1f.Valid {
		unitWeapons.Farea1.SetValid(w3uData.Ua1f.String)
	}
	if w3uData.Ua1h.Valid {
		unitWeapons.Harea1.SetValid(w3uData.Ua1h.String)
	}
	if w3uData.Ua1q.Valid {
		unitWeapons.Qarea1.SetValid(w3uData.Ua1q.String)
	}
	if w3uData.Uhd1.Valid {
		unitWeapons.Hfact1.SetValid(w3uData.Uhd1.String)
	}
	if w3uData.Uqd1.Valid {
		unitWeapons.Qfact1.SetValid(w3uData.Uqd1.String)
	}
	if w3uData.Ua1p.Valid {
		unitWeapons.SplashTargs1.SetValid("\"" + w3uData.Ua1p.String + "\"")
	}
	if w3uData.Utc1.Valid {
		unitWeapons.TargCount1.SetValid(w3uData.Utc1.String)
	}
	if w3uData.Udl1.Valid {
		unitWeapons.DamageLoss1.SetValid(w3uData.Udl1.String)
	}
	if w3uData.Usd1.Valid {
		unitWeapons.SpillDist1.SetValid(w3uData.Usd1.String)
	}
	if w3uData.Usr1.Valid {
		unitWeapons.SpillRadius1.SetValid(w3uData.Usr1.String)
	}
	if w3uData.Udu1.Valid {
		unitWeapons.DmgUp1.SetValid(w3uData.Udu1.String)
	}
	// unitWeapons.Dmod1.SetValid("0")             // TODO: Set this value correctly!
	// unitWeapons.DPS.SetValid("0")               // TODO: Set this value correctly!
	if w3uData.Ucs2.Valid {
		unitWeapons.WeapType2.SetValid(w3uData.Ucs2.String)
	}
	if w3uData.Ua2g.Valid {
		unitWeapons.Targs2.SetValid("\"" + w3uData.Ua2g.String + "\"")
	}
	if w3uData.Uwu2.Valid {
		unitWeapons.ShowUI2.SetValid(w3uData.Uwu2.String)
	}
	if w3uData.Ua2r.Valid {
		unitWeapons.RangeN2.SetValid(w3uData.Ua2r.String)
	}
	// unitWeapons.RngTst2.SetValid("\"-\"")       // TODO: Set this value correctly!
	if w3uData.Urb2.Valid {
		unitWeapons.RngBuff2.SetValid(w3uData.Urb2.String)
	}
	if w3uData.Ua2t.Valid {
		unitWeapons.AtkType2.SetValid("\"" + w3uData.Ua2t.String + "\"")
	}
	if w3uData.Ua2w.Valid {
		unitWeapons.WeapTp2.SetValid(w3uData.Ua2w.String)
	}
	if w3uData.Ua2c.Valid {
		unitWeapons.Cool2.SetValid(w3uData.Ua2c.String)
	}
	// unitWeapons.Mincool2.SetValid("\"-\"")      // TODO: Set this value correctly! (seems like it's always equal to -)
	if w3uData.Ua2d.Valid {
		unitWeapons.Dice2.SetValid(w3uData.Ua2d.String)
	}
	if w3uData.Ua2s.Valid {
		unitWeapons.Sides2.SetValid(w3uData.Ua2s.String)
	}
	if w3uData.Ua2b.Valid {
		unitWeapons.Dmgplus2.SetValid(w3uData.Ua2b.String)
	}
	if w3uData.Udu2.Valid {
		unitWeapons.DmgUp2.SetValid(w3uData.Udu2.String)
	}
	// unitWeapons.Mindmg2.SetValid("\"-\"")       // TODO: Set this value correctly!
	// unitWeapons.Avgdmg2.SetValid("\"-\"")       // TODO: Set this value correctly!
	// unitWeapons.Maxdmg2.SetValid("\"-\"")       // TODO: Set this value correctly!
	if w3uData.Udp2.Valid {
		unitWeapons.Dmgpt2.SetValid(w3uData.Udp2.String)
	}
	if w3uData.Ubs2.Valid {
		unitWeapons.BackSw2.SetValid(w3uData.Ubs2.String)
	}
	if w3uData.Ua2f.Valid {
		unitWeapons.Farea2.SetValid(w3uData.Ua2f.String)
	}
	if w3uData.Ua2h.Valid {
		unitWeapons.Harea2.SetValid(w3uData.Ua2h.String)
	}
	if w3uData.Ua2q.Valid {
		unitWeapons.Qarea2.SetValid(w3uData.Ua2q.String)
	}
	if w3uData.Uhd2.Valid {
		unitWeapons.Hfact2.SetValid(w3uData.Uhd2.String)
	}
	if w3uData.Uqd2.Valid {
		unitWeapons.Qfact2.SetValid(w3uData.Uqd2.String)
	}
	if w3uData.Ua2p.Valid {
		unitWeapons.SplashTargs2.SetValid("\"" + w3uData.Ua2p.String + "\"")
	}
	if w3uData.Utc2.Valid {
		unitWeapons.TargCount2.SetValid(w3uData.Utc2.String)
	}
	if w3uData.Udl2.Valid {
		unitWeapons.DamageLoss2.SetValid(w3uData.Udl2.String)
	}
	if w3uData.Usd2.Valid {
		unitWeapons.SpillDist2.SetValid(w3uData.Usd2.String)
	}
	if w3uData.Usr2.Valid {
		unitWeapons.SpillRadius2.SetValid(w3uData.Usr2.String)
	}
	// unitWeapons.InBeta.SetValid("0")            // TODO: Set this value correctly!

	unitAbilities.UnitAbilID.SetValid("\"" + w3uData.CustomUnitId + "\"")
	// unitAbilities.SortAbil.SetValid("") // TODO: Set this value correctly!
	// unitAbilities.Comment.SetValid("") // TODO: Set this value correctly!
	// unitAbilities.Auto.SetValid("\"_\"") // TODO: Set this value correctly!
	// unitAbilities.InBeta.SetValid("") // TODO: Set this value correctly!
	if w3uData.Uhab.Valid {
		unitAbilities.HeroAbilList.SetValid("\"" + w3uData.Uhab.String + "\"")
	}
	if w3uData.Uabi.Valid {
		unitAbilities.AbilList.SetValid("\"" + w3uData.Uabi.String + "\"")
	}
}

func (w3uData *W3uData) TransformToUnitFunc(unitFunc *UnitFunc) {
	unitFunc.UnitFuncId.SetValid(w3uData.CustomUnitId)

	if w3uData.Uma1.Valid || w3uData.Uma2.Valid {
		if w3uData.Uma1.Valid && w3uData.Uma2.Valid {
			unitFunc.Missilearc.SetValid(w3uData.Uma1.String + "," + w3uData.Uma2.String)
		} else if w3uData.Uma1.Valid {
			unitFunc.Missilearc.SetValid(w3uData.Uma1.String)
		} else if w3uData.Uma2.Valid {
			unitFunc.Missilearc.SetValid(w3uData.Uma2.String)
		}
	}
	if w3uData.Ua1m.Valid || w3uData.Ua2m.Valid {
		if w3uData.Ua1m.Valid && w3uData.Ua2m.Valid {
			unitFunc.Missileart.SetValid(w3uData.Ua1m.String + "," + w3uData.Ua2m.String)
		} else if w3uData.Ua1m.Valid {
			unitFunc.Missileart.SetValid(w3uData.Ua1m.String)
		} else if w3uData.Ua2m.Valid {
			unitFunc.Missileart.SetValid(w3uData.Ua2m.String)
		}
	}
	if w3uData.Umh1.Valid || w3uData.Umh2.Valid {
		if w3uData.Umh1.Valid && w3uData.Umh2.Valid {
			unitFunc.Missilehoming.SetValid(w3uData.Umh1.String + "," + w3uData.Umh2.String)
		} else if w3uData.Umh1.Valid {
			unitFunc.Missilehoming.SetValid(w3uData.Umh1.String)
		} else if w3uData.Umh2.Valid {
			unitFunc.Missilehoming.SetValid(w3uData.Umh2.String)
		}
	}
	if w3uData.Ua1z.Valid || w3uData.Ua2z.Valid {
		if w3uData.Ua1z.Valid && w3uData.Ua2z.Valid {
			unitFunc.Missilespeed.SetValid(w3uData.Ua1z.String + "," + w3uData.Ua2z.String)
		} else if w3uData.Ua1z.Valid {
			unitFunc.Missilespeed.SetValid(w3uData.Ua1z.String)
		} else if w3uData.Ua2z.Valid {
			unitFunc.Missilespeed.SetValid(w3uData.Ua2z.String)
		}
	}
	if w3uData.Uico.Valid {
		if w3uData.Uico.String == "-" {
			unitFunc.Art.SetValid("")
			unitFunc.Art.Valid = false
		} else {
			unitFunc.Art.SetValid(w3uData.Uico.String)
		}
	}
	if w3uData.Uspa.Valid {
		if w3uData.Uspa.String == "-" {
			unitFunc.Specialart.SetValid("")
			unitFunc.Specialart.Valid = false
		} else {
			unitFunc.Specialart.SetValid(w3uData.Uspa.String)
		}
	}
	// unitFunc.Casterupgradeart.SetValid("?") // TODO: Set the correct value
	if w3uData.Utaa.Valid {
		if w3uData.Utaa.String == "-" {
			unitFunc.Targetart.SetValid("")
			unitFunc.Targetart.Valid = false
		} else {
			unitFunc.Targetart.SetValid(w3uData.Utaa.String)
		}
	}
	if w3uData.Ussi.Valid {
		if w3uData.Ussi.String == "-" {
			unitFunc.Scorescreenicon.SetValid("")
			unitFunc.Scorescreenicon.Valid = false
		} else {
			unitFunc.Scorescreenicon.SetValid(w3uData.Ussi.String)
		}
	}
	if w3uData.Ubpx.Valid {
		if w3uData.Ubpx.String == "-" {
			unitFunc.ButtonposX.SetValid("")
			unitFunc.ButtonposX.Valid = false
		} else {
			unitFunc.ButtonposX.SetValid(w3uData.Ubpx.String)
		}
	}
	if w3uData.Ubpy.Valid {
		if w3uData.Ubpy.String == "-" {
			unitFunc.ButtonposY.SetValid("")
			unitFunc.ButtonposY.Valid = false
		} else {
			unitFunc.ButtonposY.SetValid(w3uData.Ubpy.String)
		}
	}
	if w3uData.Utra.Valid {
		if w3uData.Utra.String == "-" {
			unitFunc.Trains.SetValid("")
			unitFunc.Trains.Valid = false
		} else {
			unitFunc.Trains.SetValid(w3uData.Utra.String)
		}
	}
	if w3uData.Ubui.Valid {
		if w3uData.Ubui.String == "-" {
			unitFunc.Builds.SetValid("")
			unitFunc.Builds.Valid = false
		} else {
			unitFunc.Builds.SetValid(w3uData.Ubui.String)
		}
	}
	if w3uData.Ures.Valid {
		if w3uData.Ures.String == "-" {
			unitFunc.Researches.SetValid("")
			unitFunc.Researches.Valid = false
		} else {
			unitFunc.Researches.SetValid(w3uData.Ures.String)
		}
	}
	if w3uData.Uupt.Valid {
		if w3uData.Uupt.String == "-" {
			unitFunc.Upgrade.SetValid("")
			unitFunc.Upgrade.Valid = false
		} else {
			unitFunc.Upgrade.SetValid(w3uData.Uupt.String)
		}
	}
	if w3uData.Useu.Valid {
		if w3uData.Useu.String == "-" {
			unitFunc.Sellunits.SetValid("")
			unitFunc.Sellunits.Valid = false
		} else {
			unitFunc.Sellunits.SetValid(w3uData.Useu.String)
		}
	}
	if w3uData.Ureq.Valid {
		if w3uData.Ureq.String == "-" {
			unitFunc.Requires.SetValid("")
			unitFunc.Requires.Valid = false
		} else {
			unitFunc.Requires.SetValid(w3uData.Ureq.String)
		}
	}
	if w3uData.Uani.Valid {
		if w3uData.Uani.String == "-" {
			unitFunc.Animprops.SetValid("")
			unitFunc.Animprops.Valid = false
		} else {
			unitFunc.Animprops.SetValid(w3uData.Uani.String)
		}
	}

	// unitFunc.Requires1.SetValid("?") // TODO: Set the correct value
	// unitFunc.Requires2.SetValid("?") // TODO: Set the correct value
	// unitFunc.Requirescount.SetValid("?") // TODO: Set the correct value
	// unitFunc.Buildingsoundlabel.SetValid("?") // TODO: Set the correct value
	// unitFunc.Movementsoundlabel.SetValid("?") // TODO: Set the correct value
	// unitFunc.Loopingsoundfadein.SetValid("?") // TODO: Set the correct value
	// unitFunc.Loopingsoundfadeout.SetValid("?") // TODO: Set the correct value
	// unitFunc.Revive.SetValid("?") // TODO: Set the correct value
	// unitFunc.Attachmentanimprops.SetValid("?") // TODO: Set the correct value
	// unitFunc.Dependencyor.SetValid("?") // TODO: Set the correct value
	if w3uData.Umki.Valid {
		unitFunc.Makeitems.SetValid(w3uData.Umki.String)
	}
	if w3uData.Usei.Valid {
		unitFunc.Sellitems.SetValid(w3uData.Usei.String)
	}
	// unitString.Casterupgradeart.SetValid(?) // TODO: Set the correct value
	// unitString.Casterupgradetip.SetValid(?) // TODO: Set the correct value
	// unitString.Dependencyor.SetValid(?) // TODO: Set the correct value

	if unitFunc.Buttonpos.Valid {
		currentButtonPos := unitFunc.Buttonpos.String
		splitButtonPos := strings.Split(currentButtonPos, ",")
		if len(splitButtonPos) > 1 {
			if unitFunc.ButtonposX.Valid {
				splitButtonPos[0] = unitFunc.ButtonposX.String
			}
			if unitFunc.ButtonposY.Valid {
				splitButtonPos[1] = unitFunc.ButtonposY.String
			}
			unitFunc.Buttonpos.SetValid(splitButtonPos[0] + "," + splitButtonPos[1])
		}
	} else {
		unitFunc.Buttonpos.SetValid(unitFunc.ButtonposX.String + "," + unitFunc.ButtonposY.String)
	}
}

func (w3uData *W3uData) TransformToUnitString(unitString *UnitString) {
	unitString.UnitStringId.SetValid(w3uData.CustomUnitId)

	if w3uData.Unam.Valid {
		unitString.Name.SetValid(w3uData.Unam.String)
	}
	if w3uData.Uhot.Valid {
		unitString.Hotkey.SetValid(w3uData.Uhot.String)
	}
	if w3uData.Ides.Valid {
		unitString.Description.SetValid(w3uData.Ides.String)
	}
	if w3uData.Utip.Valid {
		unitString.Tip.SetValid(w3uData.Utip.String)
	}
	if w3uData.Utub.Valid {
		unitString.Ubertip.SetValid(w3uData.Utub.String)
	}
	if w3uData.Unsf.Valid {
		unitString.Editorsuffix.SetValid(w3uData.Unsf.String)
	}
	if w3uData.Upro.Valid {
		unitString.Propernames.SetValid(w3uData.Upro.String)
	}
	if w3uData.Utpr.Valid {
		unitString.Revivetip.SetValid(w3uData.Utpr.String)
	}
	if w3uData.Uawt.Valid {
		unitString.Awakentip.SetValid(w3uData.Uawt.String)
	}
	// unitString.Casterupgradeart.SetValid(?) // TODO: Set the correct value
	// unitString.Casterupgradetip.SetValid(?) // TODO: Set the correct value
	// unitString.Dependencyor.SetValid(?) // TODO: Set the correct value
}
