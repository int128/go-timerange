// Package timerange provides functionality of date time range.
package timerange

import (
	"fmt"
	"time"
)

// TimeRange represents a range of time with timezone.
// The range includes Start and End, i.e. [Start, End].
// Start must be before End.
type TimeRange struct {
	Start time.Time
	End   time.Time
}

// String returns a string representation of this range in RFC3339.
func (r TimeRange) String() string {
	return fmt.Sprintf("[%s, %s]", r.Start.Format(time.RFC3339), r.End.Format(time.RFC3339))
}

// Equal returns true if this range is equivalent to one.
func (r TimeRange) Equal(x TimeRange) bool {
	return r.Start.Equal(x.Start) && r.End.Equal(x.End)
}

// IsValid returns true if Start <= End.
func (r TimeRange) IsValid() bool {
	return r.Start.Equal(r.End) || r.Start.Before(r.End)
}

// Duration returns the duration between Start and End.
func (r TimeRange) Duration() time.Duration {
	return r.End.Sub(r.Start)
}

// Contains returns true if the time is in this range.
func (r TimeRange) Contains(t time.Time) bool {
	return r.Start.Equal(t) || r.End.Equal(t) || (r.Start.Before(t) && t.Before(r.End))
}

// Before returns true if this range is before the time.
func (r TimeRange) Before(t time.Time) bool {
	return r.End.Before(t)
}

// After returns true if this range is after the time.
func (r TimeRange) After(t time.Time) bool {
	return r.Start.After(t)
}

// Intersect returns an intersection of ranges.
func Intersect(a, b TimeRange) TimeRange {
	r := TimeRange{
		Start: maxTime(a.Start, b.Start),
		End:   minTime(a.End, b.End),
	}
	if !r.IsValid() {
		return TimeRange{}
	}
	return r
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
