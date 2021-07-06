package dto

type Week struct {
	Monday    WeekDay `json:"monday"`
	Tuesday   WeekDay `json:"tuesday"`
	Wednesday WeekDay `json:"wednesday"`
	Thursday  WeekDay `json:"thursday"`
	Friday    WeekDay `json:"friday"`
	Saturday  WeekDay `json:"saturday"`
}
