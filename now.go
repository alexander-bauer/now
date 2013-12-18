package now

import "time"

func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
	return now.Truncate(time.Hour)
}

func (now *Now) BeginningOfDay() time.Time {
	d := time.Duration(-now.Hour()) * time.Hour
	return now.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if FirstDayMonday {
		if weekday == 0 {
			weekday = 7
		}
		weekday = weekday - 1
	}

	d := time.Duration(-weekday) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfMonth() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.Day())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfYear() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.YearDay())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
	return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// NextOccurrenceOf finds the next occurrence of the given
// year-agnostic date, including the current day, and returns its
// beginning.
//
// For example, if the current date is January 2nd, 2006, and one
// calls this with arguments 1 and 5, it will return January 5th,
// 2006, 00:00. If one calls it with the arguments 1 and 1, it will
// return January 1st, 2007, 00:00.
func (now *Now) NextOccurrenceOf(month time.Month, day int) time.Time {
	var yearsToAdd int

	_, thisMonth, today := now.Time.Date()

	// Check if the given month is before the one we're comparing it
	// to. If so, then we need to use the next year.
	if month < thisMonth {
		yearsToAdd = 1
	} else if month == thisMonth {
		// If the months are equal, check the current day.
		if day < today {
			yearsToAdd = 1
		}
	}

	return now.BeginningOfYear().AddDate(yearsToAdd, int(month)-1, day-1)
}

// LastOccurrenceOf finds the last occurrence of the given
// year-agnostic date, excluding the current day, and returns its
// beginning.
//
// For example, if the current date is January 2nd, 2006, and one
// calls this with arguments 1 and 5, it will return January 5th,
// 2005, 00:00. If one calls it with the arguments 1 and 1, it will
// return January 1st, 2006, 00:00.
func (now *Now) LastOccurrenceOf(month time.Month, day int) time.Time {
	var yearsToSub int

	_, thisMonth, today := now.Time.Date()

	// Check if the given month is after the one we're comparing it
	// to. If so, then we need to use the previous year.
	if month > thisMonth {
		yearsToSub = 1
	} else if month == thisMonth {
		// If the months are equal, check the current day.
		if day >= today {
			yearsToSub = 1
		}
	}

	return now.BeginningOfYear().AddDate(-yearsToSub, int(month)-1, day-1)
}

func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) Sunday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	} else {
		d := time.Duration(7-weekday) * 24 * time.Hour
		return t.Truncate(time.Hour).Add(d)
	}
}

func (now *Now) EndOfSunday() time.Time {
	return now.Sunday().Add(24*time.Hour - time.Nanosecond)
}
