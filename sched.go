package cmscal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

type Schedule map[DayScheduleType]DaySchedule

var ScheduleSixth = Schedule{
	BlueDay: DaySchedule{
		{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
		{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 2"},
		{StartHour: 10, StartMinute: 59, Duration: time.Minute * 45, Description: "Period 3"},
		{StartHour: 11, StartMinute: 49, Duration: time.Minute * 30, Description: "Lunch"},
		{StartHour: 12, StartMinute: 19, Duration: time.Minute * 45, Description: "Period 3"},
		{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 4"},
		{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
		{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
	},
	WhiteDay: DaySchedule{
		{StartHour: 8, StartMinute: 47, Duration: time.Minute * 32, Description: "Advisory/Period 1"},
		{StartHour: 9, StartMinute: 24, Duration: time.Minute * 90, Description: "Period 5"},
		{StartHour: 10, StartMinute: 59, Duration: time.Minute * 45, Description: "Period 6"},
		{StartHour: 11, StartMinute: 49, Duration: time.Minute * 30, Description: "Lunch"},
		{StartHour: 12, StartMinute: 19, Duration: time.Minute * 45, Description: "Period 6"},
		{StartHour: 13, StartMinute: 9, Duration: time.Minute * 90, Description: "Period 7"},
		{StartHour: 14, StartMinute: 44, Duration: time.Minute * 45, Description: "Period 8"},
		{StartHour: 15, StartMinute: 32, Duration: time.Minute * 8, Description: "Period 1"},
	},
}

func MakeHolidayMap(loc *time.Location) map[time.Time]bool {
	m := make(map[time.Time]bool)
	holidays := []time.Time{
		time.Date(2020, time.November, 2, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 3, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 25, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 26, 0, 0, 0, 0, loc),
		time.Date(2020, time.November, 27, 0, 0, 0, 0, loc),
	}
	for _, h := range holidays {
		m[h] = true
	}
	return m
}

func ICalForSchedule(loc *time.Location, s Schedule) string {
	now := time.Now()

	cal := ics.NewCalendar()
	cal.SetName("CMS Sixth Grade 2020-2021")
	cal.SetDescription("CMS Sixth Grade 2020-2021")
	cal.SetXWRCalName("CMS Sixth Grade 2020-2021")
	cal.SetXWRCalDesc("CMS Sixth Grade 2020-2021")

	holidayMap := MakeHolidayMap(loc)
	startDate := time.Date(2020, time.August, 17, 0, 0, 0, 0, loc)
	daysFromStart := 0
	currentDateType := BlueDay
	const totalWeekdays = 90
	weekdaysDone := 0
	for weekdaysDone < totalWeekdays {
		date := startDate.AddDate(0, 0, daysFromStart)
		daysFromStart += 1
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}
		weekdaysDone += 1
		if _, present := holidayMap[date]; present {
			continue
		}

		for _, period := range s[currentDateType] {
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
		currentDateType = (currentDateType + 1) % 2
	}

	return cal.Serialize()
}
