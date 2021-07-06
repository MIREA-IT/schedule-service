package dto

import "therealmone.github.io/mirea-schedule-bot/src/enum"

type Lesson struct {
	Number   enum.LessonNumber `json:"number"`
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Teacher  string            `json:"teacher"`
	Audience string            `json:"audience"`
}
