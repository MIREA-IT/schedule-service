package schedule

import "testing"

func TestScheduleLoad(t *testing.T) {
	schedules := Load()

	if len(schedules) == 0 {
		t.Fatal("schedules are empty")
	}
}
