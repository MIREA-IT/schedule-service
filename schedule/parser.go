package schedule

import "therealmone.github.io/mirea-schedule-bot/dto"

func Load() []dto.Schedule {
	var schedules []dto.Schedule

	schedules = append(schedules, dto.Schedule{Name: "3 курс ИТ"}, dto.Schedule{Name: "4 курс ИТ"})

	return schedules
}
