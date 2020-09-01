package cmscal

import "time"

type DayScheduleType int

const (
	BlueDay DayScheduleType = iota
	WhiteDay
)

func (s DayScheduleType) String() string {
	return []string{"BLUE DAY", "WHITE DAY"}[s]
}

type ClassPeriod struct {
	StartHour   int
	StartMinute int
	Duration    time.Duration
	Description string
}

type DaySchedule []ClassPeriod

type GradeSchedule struct {
	Name        string
	Description string
	ScheduleMap map[DayScheduleType]DaySchedule
}

type BuildingSchedule struct {
	startDate    time.Time
	numWeekdays  int
	firstDayType DayScheduleType
	holidays     map[time.Time]bool
}
