package pramanapi

import "time"

func (d *Date) DaysFrom(t time.Time) int {
	year, month, day := t.Date()

	dtime := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.UTC)
	ttime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	duration := dtime.Sub(ttime)

	return int(duration.Hours() / 24)
}
