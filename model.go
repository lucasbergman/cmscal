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
