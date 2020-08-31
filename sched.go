package cmscal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

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

var ScheduleSixth = GradeSchedule{
	Name:        "CMS Sixth Grade 2020-2021",
	Description: "Block schedule for CMS Sixth Grade 2020-2021",
	ScheduleMap: map[DayScheduleType]DaySchedule{
		BlueDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 2"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 45, Description: "Period 3"},
			{StartHour: 11, StartMinute: 49, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 12, StartMinute: 19, Duration: time.Minute * 45, Description: "Period 3"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 4"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
		WhiteDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 5"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 45, Description: "Period 6"},
			{StartHour: 11, StartMinute: 49, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 12, StartMinute: 19, Duration: time.Minute * 45, Description: "Period 6"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 7"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
	},
}

var ScheduleSeventh = GradeSchedule{
	Name:        "CMS Seventh Grade 2020-2021",
	Description: "Block schedule for CMS Seventh Grade 2020-2021",
	ScheduleMap: map[DayScheduleType]DaySchedule{
		BlueDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 2"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 11, StartMinute: 34, Duration: time.Minute * 90, Description: "Period 3"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 4"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
		WhiteDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 5"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 11, StartMinute: 34, Duration: time.Minute * 90, Description: "Period 6"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 7"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
	},
}

var ScheduleEighth = GradeSchedule{
	Name:        "CMS Eighth Grade 2020-2021",
	Description: "Block schedule for CMS Eighth Grade 2020-2021",
	ScheduleMap: map[DayScheduleType]DaySchedule{
		BlueDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 2"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 90, Description: "Period 3"},
			{StartHour: 12, StartMinute: 34, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 4"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
		WhiteDay: {
			{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
			{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 5"},
			{StartHour: 10, StartMinute: 59, Duration: time.Minute * 90, Description: "Period 6"},
			{StartHour: 12, StartMinute: 34, Duration: time.Minute * 30, Description: "Lunch"},
			{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 7"},
			{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
			{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
		},
	},
}

func makeHolidayMap(loc *time.Location) map[time.Time]bool {
	m := make(map[time.Time]bool)
	for _, h := range []time.Time{
		time.Date(2020, time.September, 7, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 2, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 3, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 25, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 26, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 27, 0, 0, 0, 0, loc),
	} {
		m[h] = true
	}
	return m
}

func MakeBuildingSchedule(loc *time.Location) *BuildingSchedule {
	return &BuildingSchedule{
		startDate:    time.Date(2020, time.August, 17, 0, 0, 0, 0, loc),
		numWeekdays:  90,
		firstDayType: BlueDay,
		holidays:     makeHolidayMap(loc),
	}
}

func ICalForSchedule(bs *BuildingSchedule, gs *GradeSchedule) string {
	now := time.Now()

	cal := ics.NewCalendar()
	cal.SetXWRCalName(gs.Name)
	cal.SetXWRCalDesc(fmt.Sprintf("%s. Source code at <https://github.com/lucasbergman/cmscal>.", gs.Description))
	cal.SetName(gs.Name)
	cal.SetDescription(fmt.Sprintf("%s. Source code at <https://github.com/lucasbergman/cmscal>.", gs.Description))

	startDate := bs.startDate
	daysFromStart := 0
	currentDayType := bs.firstDayType
	weekdaysDone := 0
	for weekdaysDone < bs.numWeekdays {
		date := startDate.AddDate(0, 0, daysFromStart)
		daysFromStart += 1
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}
		weekdaysDone += 1
		if _, present := bs.holidays[date]; present {
			continue
		}

		for _, period := range gs.ScheduleMap[currentDayType] {
			start := date.Add(time.Duration(period.StartHour)*time.Hour + time.Duration(period.StartMinute)*time.Minute)
			end := start.Add(period.Duration)

			hasher := sha1.New()
			hasher.Write([]byte(start.String()))
			hasher.Write([]byte(end.String()))
			hasher.Write([]byte(period.Description))
			hash := hex.EncodeToString(hasher.Sum(nil))

			event := cal.AddEvent(fmt.Sprintf("%s@cmscal.bergmans.us", hash))
			event.SetDtStampTime(now)
			event.SetStartAt(start)
			event.SetEndAt(end)
			event.AddProperty(ics.ComponentProperty("TRANSP"), "TRANSPARENT")
			event.SetSummary(period.Description)
			event.SetDescription(period.Description)
		}
		currentDayType = (currentDayType + 1) % 2
	}

	return cal.Serialize()
}
