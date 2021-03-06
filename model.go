package cmscal

import "time"

type BlockDayType int

const (
	BlueDay BlockDayType = iota
	WhiteDay
	NonBlockDay
	ShortWednesday
)

func (t BlockDayType) String() string {
	return []string{"BLUE DAY", "WHITE DAY", "DAILY SCHEDULE", "SHORT WEDNESDAY"}[t]
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
	ScheduleMap map[BlockDayType]DaySchedule
}

type BuildingSchedule struct {
	startDate    time.Time
	numWeekdays  int
	firstDayType BlockDayType
	holidays     map[time.Time]bool

	// "I have altered our deal. Pray I do not alter it again."
	nonBlockChangeDay time.Time

	// Haha nothing matters
	shortWednesdayChangeDay time.Time
}

func nextDayType(bs *BuildingSchedule, yesterday BlockDayType, today time.Time) BlockDayType {
	if today.After(bs.shortWednesdayChangeDay) && today.Weekday() == time.Wednesday {
		return ShortWednesday
	}
	if today.After(bs.nonBlockChangeDay) {
		return NonBlockDay
	}
	if yesterday == WhiteDay {
		return BlueDay
	}
	return WhiteDay
}
