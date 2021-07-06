package controller

import (
	"log"
	"net/http"
	"therealmone.github.io/mirea-schedule-bot/src/dto"
	"therealmone.github.io/mirea-schedule-bot/src/parser"
	"therealmone.github.io/mirea-schedule-bot/src/utils"
)

var (
	schedules = initSchedules()
)

func initSchedules() *[]dto.Schedule {
	schedule, err := parser.Parse("../../resources/ИИТ_3к_20-21_весна.xlsx")
	if err != nil {
		log.Fatal("error: ", err)
	}

	return &[]dto.Schedule{
		*schedule,
	}
}

func GetSchedules(response http.ResponseWriter, request *http.Request) {
	log.Println("Handling /schedules request")

	utils.Encode(response, schedules)
}
