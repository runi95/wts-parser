package models

type UnitBalanceTemplate struct {
	RowCount    int
	UnitBalance []*UnitBalance
}

type UnitUITemplate struct {
	RowCount int
	UnitUI   []*UnitUI
}

type UnitAbilitiesTemplate struct {
	RowCount      int
	UnitAbilities []*UnitAbilities
}

type UnitDataTemplate struct {
	RowCount int
	UnitData []*UnitData
}

type UnitWeaponsTemplate struct {
	RowCount    int
	UnitWeapons []*UnitWeapons
}
