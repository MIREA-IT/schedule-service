package enum

type LessonNumber struct {
	Number    int
	StartTime string
	EndTime   string
}

var (
	LessonN1 = LessonNumber{Number: 1, StartTime: "9-00", EndTime: "10-30"}
	LessonN2 = LessonNumber{Number: 2, StartTime: "10-40", EndTime: "12-10"}
	LessonN3 = LessonNumber{Number: 3, StartTime: "12-40", EndTime: "14-10"}
	LessonN4 = LessonNumber{Number: 4, StartTime: "14-20", EndTime: "15-50"}
	LessonN5 = LessonNumber{Number: 5, StartTime: "16-20", EndTime: "17-50"}
	LessonN6 = LessonNumber{Number: 6, StartTime: "18-00", EndTime: "19-30"}

	NumberToEnumMapping = map[int]LessonNumber{
		1: LessonN1,
		2: LessonN2,
		3: LessonN3,
		4: LessonN4,
		5: LessonN5,
		6: LessonN6,
	}
)

func ByNumber(number int) LessonNumber {
	return NumberToEnumMapping[number]
}
