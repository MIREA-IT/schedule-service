package controller

import (
	"log"
	"net/http"
	"therealmone.github.io/mirea-schedule-bot/schedule"
	"therealmone.github.io/mirea-schedule-bot/utils"
)

var (
	schedules = schedule.Load()
)

func GetSchedules(response http.ResponseWriter, request *http.Request) {
	log.Println("Handling /schedules request")

	utils.Encode(response, schedules)
}
