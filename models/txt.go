package models

import "gopkg.in/volatiletech/null.v6"

type UnitFunc struct {
	UnitId string

	Art                 null.String
	Specialart          null.String
	Casterupgradeart    null.String
	Targetart           null.String
	Scorescreenicon     null.String
	Buttonpos           null.String
	ButtonposX          null.String
	ButtonposY          null.String
	Missileart          null.String
	Missileart1         null.String
	Missileart2         null.String
	Missilearc          null.String
	Missilearc1         null.String
	Missilearc2         null.String
	Missilespeed        null.String
	Missilespeed1       null.String
	Missilespeed2       null.String
	Missilehoming       null.String
	Missilehoming1      null.String
	Missilehoming2      null.String
	Trains              null.String
	Builds              null.String
	Researches          null.String
	Upgrade             null.String
	Sellunits           null.String
	Requires            null.String
	Requires1           null.String
	Requires2           null.String
	Requirescount       null.String
	Buildingsoundlabel  null.String
	Movementsoundlabel  null.String
	Loopingsoundfadein  null.String
	Loopingsoundfadeout null.String
	Revive              null.String
	Animprops           null.String
	Attachmentanimprops null.String
	Dependencyor        null.String
	Makeitems           null.String
	Attachmentlinkprops null.String
	Boneprops           null.String
	Sellitems           null.String
	Randomsoundlabel    null.String

	Name              null.String
	Hotkey            null.String
	Description       null.String
	Tip               null.String
	Ubertip           null.String
	Editorsuffix      null.String
	Propernames       null.String
	Revivetip         null.String
	Awakentip         null.String
	Casterupgradename null.String
	Casterupgradetip  null.String
}

type UnitFuncs struct {
	CampaignUnitFuncs []*UnitFunc
	HumanUnitFuncs    []*UnitFunc
	NeutralUnitFuncs  []*UnitFunc
	NightElfUnitFuncs []*UnitFunc
	OrcUnitFuncs      []*UnitFunc
	UndeadUnitFuncs   []*UnitFunc
}

type UnitString struct {
	UnitId string

	Name              null.String
	Hotkey            null.String
	Description       null.String
	Tip               null.String
	Ubertip           null.String
	Editorsuffix      null.String
	Propernames       null.String
	Revivetip         null.String
	Awakentip         null.String
	Casterupgradename null.String
	Casterupgradetip  null.String
	Dependencyor      null.String
}

type UnitStrings struct {
	CampaignUnitFuncs   []*UnitString
	HumanUnitStrings    []*UnitString
	NeutralUnitStrings  []*UnitString
	NightElfUnitStrings []*UnitString
	OrcUnitStrings      []*UnitString
	UndeadUnitStrings   []*UnitString
}