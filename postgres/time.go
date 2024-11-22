package postgres

import "time"

const TimePrecision = time.Microsecond

// Time sets the location to UTC and truncates to the highest precision
// that Postgres supports.
func Time(t time.Time) time.Time {
	return t.Truncate(TimePrecision).UTC()
}
