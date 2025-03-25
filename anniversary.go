package today

import (
	"time"
)

// IsInternationalDay indicates if the provided time.Time is the requested
// international day (which is the 50 hours spanning the world -- from UTC+14 to
// UTC-12 -- for which it could be valid).
func IsInternationalDay(day int, month time.Month, zero time.Time) bool {
	if zero.IsZero() {
		return false
	}
	zero = zero.Round(time.Second).UTC()
	// the furthest west starting time
	// note: -1 sec to avoid equality comparison
	s := time.Date(zero.Year(), month, day, 0, 0, -1, 0, time.FixedZone("UTC+14", 14*60*60))
	// the further east ending time
	// note: day+1 = end of day
	e := time.Date(zero.Year(), month, day+1, 0, 0, 0, 0, time.FixedZone("UTC-12", -12*60*60))
	// supplied time within those bounds?
	return zero.Before(e) && zero.After(s)
}

// IsDayMonth indicates that the provided time.Time is the requested day and month.
func IsDayMonth(day int, month time.Month, t time.Time) bool {
	return t.Month() == month && t.Day() == day
}
