package parser

import (
	"encoding/json"
	"io/ioutil"
	"mireait.github.io/mirea-schedule-bot/src/api"
	"reflect"
	"testing"
)

func TestScheduleLoad(t *testing.T) {
	schedule, err := Parse("../../resources/ИИТ_3к_20-21_весна.xlsx")

	if err != nil {
		t.Fatal("error: ", err)
	}

	if len(schedule) != 33 {
		t.Fatal("invalid groups count")
	}

	assertGroupSchedule(schedule["ИАБО-01-18"], "ИАБО-01-18-schedule.json", t)
	assertGroupSchedule(schedule["ИКБО-13-18"], "ИКБО-13-18-schedule.json", t)
	assertGroupSchedule(schedule["ИНБО-04-18"], "ИНБО-04-18-schedule.json", t)
}

func assertGroupSchedule(group *api.GroupSchedule, expectedScheduleJson string, t *testing.T) {
	expectedGroup := new(api.GroupSchedule)

	bytes, err := ioutil.ReadFile("../../resources-test/" + expectedScheduleJson)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, expectedGroup)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(&expectedGroup, &group) {
		t.Fatal("Group schedule not equals to expected")
	}
}
