// Package timerange provides functionality of date time range.
package timerange

import (
	"fmt"
	"time"
)

// New returns a TimeRange with start and end.
func New(start, end time.Time) TimeRange {
	return TimeRange{start: start, end: end}
}

// NewFrom returns a TimeRange with start and duration.
func NewFrom(start time.Time, duration time.Duration) TimeRange {
	return New(start, start.Add(duration))
}

// TimeRange represents an immutable range of time with timezone.
// The range includes start and end, i.e. [start, end].
// Start must be before end.
type TimeRange struct {
	start time.Time
	end   time.Time
}

// Start returns the start time.
func (r TimeRange) Start() time.Time {
	return r.start
}

// End returns the end time.
func (r TimeRange) End() time.Time {
	return r.end
}

// String returns a string representation of this range in RFC3339.
func (r TimeRange) String() string {
	return fmt.Sprintf("[%s, %s]", r.start.Format(time.RFC3339), r.end.Format(time.RFC3339))
}

// Equal returns true if this range is equivalent to one.
func (r TimeRange) Equal(x TimeRange) bool {
	return r.start.Equal(x.start) && r.end.Equal(x.end)
}

// IsZero returns true if both start and end are zero value.
func (r TimeRange) IsZero() bool {
	return r.start.IsZero() && r.end.IsZero()
}

// IsValid returns true if start <= end.
func (r TimeRange) IsValid() bool {
	return r.start.Equal(r.end) || r.start.Before(r.end)
}

// Duration returns the duration between start and end.
func (r TimeRange) Duration() time.Duration {
	return r.end.Sub(r.start)
}

// Contains returns true if the time is in this range.
func (r TimeRange) Contains(t time.Time) bool {
	return r.start.Equal(t) || r.end.Equal(t) || (r.start.Before(t) && t.Before(r.end))
}

// Before returns true if this range is before the time.
func (r TimeRange) Before(t time.Time) bool {
	return r.end.Before(t)
}

// After returns true if this range is after the time.
func (r TimeRange) After(t time.Time) bool {
	return r.start.After(t)
}

// Intersect returns the intersection of given ranges.
// If the intersection is empty, this returns a zero struct.
func Intersect(a, b TimeRange) TimeRange {
	r := TimeRange{
		start: maxTime(a.start, b.start),
		end:   minTime(a.end, b.end),
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
