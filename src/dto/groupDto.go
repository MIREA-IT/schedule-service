package dto

type Group struct {
	GroupCode string `json:"group_code"`
	OddWeek   Week   `json:"odd_week"`
	EvenWeek  Week   `json:"even_week"`
}
