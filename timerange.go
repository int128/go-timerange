// Package timerange provides a simple type of time range.
// It is immutable type.
package timerange

import (
	"fmt"
	"time"
)

// New returns a TimeRange with start and end.
// If start > end, this returns a zero value.
func New(start, end time.Time) TimeRange {
	if start.After(end) {
		return TimeRange{}
	}
	return TimeRange{start: start, end: end}
}

// From returns a TimeRange with start and duration.
// Duration must be positive.
func From(start time.Time, duration time.Duration) TimeRange {
	return New(start, start.Add(duration))
}

// Until returns a TimeRange with end and duration.
// Duration must be positive.
func Until(end time.Time, duration time.Duration) TimeRange {
	return New(end.Add(-duration), end)
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

// Shift returns a TimeRange moved for the duration.
// Duration can be positive or negative.
func (r TimeRange) Shift(d time.Duration) TimeRange {
	return New(r.start.Add(d), r.end.Add(d))
}

// Extend returns an extended TimeRange for the duration.
// Duration can be positive or negative.
func (r TimeRange) Extend(d time.Duration) TimeRange {
	return New(r.start, r.end.Add(d))
}

// In returns true if the time is within the range.
// This is a synonym of TimeRange.Contains().
func In(t time.Time, r TimeRange) bool {
	return r.Contains(t)
}

// Intersect returns the intersection of given ranges.
// If the intersection is empty, this returns a zero struct.
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
