package parser

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"mireait.github.io/mirea-schedule-bot/src/api"
	"regexp"
	"strconv"
)

const (
	GroupNamePattern = "[a-zA-Zа-яА-Я]{4}-\\d{2}-\\d{2}"
)

var (
	WeekDayMapping = map[int]func(week *api.Week, day *api.WeekDay){
		0: func(week *api.Week, day *api.WeekDay) { week.Monday = day },
		1: func(week *api.Week, day *api.WeekDay) { week.Tuesday = day },
		2: func(week *api.Week, day *api.WeekDay) { week.Wednesday = day },
		3: func(week *api.Week, day *api.WeekDay) { week.Thursday = day },
		4: func(week *api.Week, day *api.WeekDay) { week.Friday = day },
		5: func(week *api.Week, day *api.WeekDay) { week.Saturday = day },
	}
)

type ScheduleContext struct {
	lessonNumbers *[]string
	weakParity    *[]string
	lessons       *[]string
	typesOfLesson *[]string
	teacherNames  *[]string
	audiences     *[]string
}

func Parse(scheduleFileName string) (map[string]*api.GroupSchedule, error) {
	schedule := make(map[string]*api.GroupSchedule)

	unparsedSchedule, err := excelize.OpenFile(scheduleFileName)
	if err != nil {
		return nil, err
	}

	cols, err := unparsedSchedule.GetCols("Лист1")
	if err != nil {
		return nil, err
	}

	lessonNumbers := cols[1]
	weakParity := cols[4]

	for index, column := range cols {
		groupCode := column[1]
		isGroupColumn, _ := regexp.Match(GroupNamePattern, []byte(groupCode))

		if isGroupColumn {
			context := ScheduleContext{
				lessonNumbers: &lessonNumbers,
				weakParity:    &weakParity,
				lessons:       &column,
				typesOfLesson: &cols[index+1],
				teacherNames:  &cols[index+2],
				audiences:     &cols[index+3],
			}

			schedule[groupCode] = collectGroupSchedule(groupCode, &context)
		}
	}

	return schedule, nil
}

func collectGroupSchedule(groupCode string, context *ScheduleContext) *api.GroupSchedule {
	oddWeek := new(api.Week)
	evenWeek := new(api.Week)

	for i := 3; i < 75; i += 12 {
		oddWeekDay := new(api.WeekDay)
		evenWeekDay := new(api.WeekDay)

		for j := i; j < i+12; j++ {
			lesson := collectLesson(j, context)

			if j%2 == 0 {
				evenWeekDay.Lessons = append(evenWeekDay.Lessons, lesson)
			} else {
				oddWeekDay.Lessons = append(oddWeekDay.Lessons, lesson)
			}
		}

		WeekDayMapping[i/12](oddWeek, oddWeekDay)
		WeekDayMapping[i/12](evenWeek, evenWeekDay)
	}

	return &api.GroupSchedule{
		GroupCode: groupCode,
		OddWeek:   oddWeek,
		EvenWeek:  evenWeek,
	}
}

func collectLesson(rowIndex int, context *ScheduleContext) *api.Lesson {
	return &api.Lesson{
		Number:   resolveNumber(context.lessonNumbers, rowIndex),
		Name:     (*context.lessons)[rowIndex],
		Type:     (*context.typesOfLesson)[rowIndex],
		Teacher:  (*context.teacherNames)[rowIndex],
		Audience: (*context.audiences)[rowIndex],
	}
}

func resolveNumber(numbers *[]string, i int) uint32 {
	number := (*numbers)[i]

	if len(number) == 0 {
		return toInt((*numbers)[i-1])
	} else {
		return toInt(number)
	}
}

func toInt(s string) uint32 {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatal("Invalid int: ", err)
	}

	return uint32(value)
}
