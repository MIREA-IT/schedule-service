package parser

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"regexp"
	"strconv"
	"therealmone.github.io/mirea-schedule-bot/src/dto"
	"therealmone.github.io/mirea-schedule-bot/src/enum"
)

const (
	GroupNamePattern = "[a-zA-Zа-яА-Я]{4}-\\d{2}-\\d{2}"
)

var (
	WeekDayMapping = map[int]func(week *dto.Week, day *dto.WeekDay){
		0: func(week *dto.Week, day *dto.WeekDay) { week.Monday = *day },
		1: func(week *dto.Week, day *dto.WeekDay) { week.Tuesday = *day },
		2: func(week *dto.Week, day *dto.WeekDay) { week.Wednesday = *day },
		3: func(week *dto.Week, day *dto.WeekDay) { week.Thursday = *day },
		4: func(week *dto.Week, day *dto.WeekDay) { week.Friday = *day },
		5: func(week *dto.Week, day *dto.WeekDay) { week.Saturday = *day },
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

func Parse(scheduleFileName string) (*dto.Schedule, error) {
	var schedule dto.Schedule

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

			schedule.Groups = append(schedule.Groups, collectGroupSchedule(groupCode, &context))
		}
	}

	return &schedule, nil
}

func collectGroupSchedule(groupCode string, context *ScheduleContext) dto.Group {
	oddWeek := new(dto.Week)
	evenWeek := new(dto.Week)

	for i := 3; i < 75; i += 12 {
		oddWeekDay := new(dto.WeekDay)
		evenWeekDay := new(dto.WeekDay)

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

	return dto.Group{
		GroupCode: groupCode,
		OddWeek:   *oddWeek,
		EvenWeek:  *evenWeek,
	}
}

func collectLesson(rowIndex int, context *ScheduleContext) dto.Lesson {
	return dto.Lesson{
		Number:   enum.ByNumber(resolveNumber(context.lessonNumbers, rowIndex)),
		Name:     (*context.lessons)[rowIndex],
		Type:     (*context.typesOfLesson)[rowIndex],
		Teacher:  (*context.teacherNames)[rowIndex],
		Audience: (*context.audiences)[rowIndex],
	}
}

func resolveNumber(numbers *[]string, i int) int {
	number := (*numbers)[i]

	if len(number) == 0 {
		return toInt((*numbers)[i-1])
	} else {
		return toInt(number)
	}
}

func toInt(s string) int {
	value, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatal("Invalid int: ", err)
	}

	return int(value)
}
