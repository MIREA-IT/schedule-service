package parser

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"therealmone.github.io/mirea-schedule-bot/src/dto"
	"therealmone.github.io/mirea-schedule-bot/src/utils"
)

func TestScheduleLoad(t *testing.T) {
	schedule, err := Parse("../../resources/ИИТ_3к_20-21_весна.xlsx")

	if err != nil {
		t.Fatal("error: ", err)
	}

	if len(schedule.Groups) != 33 {
		t.Fatal("invalid groups count")
	}

	assertGroupSchedule(&schedule.Groups[0], "ИАБО-01-18-schedule.json", t)
	assertGroupSchedule(&schedule.Groups[16], "ИКБО-13-18-schedule.json", t)
	assertGroupSchedule(&schedule.Groups[32], "ИНБО-04-18-schedule.json", t)
}

func assertGroupSchedule(group *dto.Group, expectedScheduleJson string, t *testing.T) {
	expectedGroup := new(dto.Group)

	bytes, err := ioutil.ReadFile("../../resources-test/" + expectedScheduleJson)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, expectedGroup)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(&expectedGroup, &group) {
		t.Fatalf("Group schedule not equals to expected.\nExpected: %s\nActual: %s",
			utils.Marshal(expectedGroup), utils.Marshal(group))
	}
}
