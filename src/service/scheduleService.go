package service

import (
	"errors"
	"log"
	"mireait.github.io/mirea-schedule-bot/src/api"
	"mireait.github.io/mirea-schedule-bot/src/parser"
)

type ScheduleService struct {
	schedules map[string]*api.GroupSchedule
}

func NewScheduleService(scheduleFilename string) *ScheduleService {
	schedules, err := parser.Parse(scheduleFilename)
	if err != nil {
		log.Fatalln("Schedule parsing error: ", err)
	}

	return &ScheduleService{
		schedules: schedules,
	}
}

func (service *ScheduleService) GetGroupSchedule(groupCode string) (*api.GroupSchedule, error) {
	groupSchedule := service.schedules[groupCode]
	if groupSchedule == nil {
		return nil, errors.New("Schedule not found for group " + groupCode)
	}

	return groupSchedule, nil
}
