package time

import "time"

func StartOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func StartOfCurrentDay() time.Time {
	return StartOfDay(time.Now().Year(), time.Now().Month(), time.Now().Day())
}

// StartOfWeek return the start of the week for the given year and week.
func StartOfWeek(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func StartOfCurrentWeek() time.Time {
	return StartOfWeek(time.Now().ISOWeek())
}

// StartOfMonth return the start of the month for the given year and month.
func StartOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

func StartOfCurrentMonth() time.Time {
	return StartOfMonth(time.Now().Year(), time.Now().Month())
}
