package timerange

import "time"

// Intersect returns the intersection of given ranges.
// If the intersection is empty, this returns a zero value.
func Intersect(a, b TimeRange) TimeRange {
	return New(
		maxTime(a.start, b.start),
		minTime(a.end, b.end),
	)
}

func minTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
