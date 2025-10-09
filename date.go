package pramanapi

import (
	"fmt"
	"time"
)

func NewDate(t time.Time) *Date {
	year, month, day := t.Date()
	return &Date{
		Year:  uint32(year),
		Month: uint32(month),
		Day:   uint32(day),
	}
}

// DateFromTime creates a Date from a time.Time.
// Deprecated: Use NewDate instead.
func DateFromTime(t time.Time) *Date {
	return NewDate(t)
}

func (d *Date) DaysFrom(t time.Time) int {
	year, month, day := t.Date()

	dtime := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	ttime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	duration := dtime.Sub(ttime)

	return int(duration.Hours() / 24)
}

// Time converts a Date to a time.Time (at midnight UTC).
// Usually, Date is initialized using user input, so this function returns an error if the Date is invalid.
func (d *Date) Time() (time.Time, error) {
	return time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day))
}

func DayOfWeekFromTime(t time.Time) DayOfWeek {
	switch t.Weekday() {
	case time.Sunday:
		return DayOfWeek_SUNDAY
	case time.Monday:
		return DayOfWeek_MONDAY
	case time.Tuesday:
		return DayOfWeek_TUESDAY
	case time.Wednesday:
		return DayOfWeek_WEDNESDAY
	case time.Thursday:
		return DayOfWeek_THURSDAY
	case time.Friday:
		return DayOfWeek_FRIDAY
	case time.Saturday:
		return DayOfWeek_SATURDAY
	default:
		return DayOfWeek_DAY_OF_WEEK_UNSPECIFIED
	}
}
