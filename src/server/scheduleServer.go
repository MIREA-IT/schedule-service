package server

import (
	"context"
	"log"
	"mireait.github.io/mirea-schedule-bot/src/api"
	"mireait.github.io/mirea-schedule-bot/src/service"
)

type ScheduleServer struct {
	api.UnimplementedScheduleServiceServer

	ScheduleService *service.ScheduleService
}

func (server *ScheduleServer) GetGroupSchedule(ctx context.Context, request *api.GroupScheduleRequest) (*api.GroupSchedule, error) {
	log.Printf("Received GetGroupSchedule request. Group code: %s, initiator user: %s",
		request.GetGroupCode(), request.GetInitiatorUser())

	schedule, err := server.ScheduleService.GetGroupSchedule(request.GroupCode)
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
