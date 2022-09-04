// Package timerange provides simple types to handle a time range.
package timerange

import (
	"fmt"
	"time"
)

// New returns a TimeRange with start time and end time.
// It must be start <= end.
// If start > end, this returns a zero value.
func New(start, end time.Time) TimeRange {
	if start.After(end) {
		return TimeRange{}
	}
	return TimeRange{start: start, end: end}
}

// From returns a TimeRange with start time and duration.
// The duration must be positive.
func From(start time.Time, duration time.Duration) TimeRange {
	return New(start, start.Add(duration))
}

// Until returns a TimeRange with end time and duration.
// The duration must be positive.
func Until(end time.Time, duration time.Duration) TimeRange {
	return New(end.Add(-duration), end)
}

// TimeRange represents an immutable range of time with timezone.
// The range includes start time and end time, i.e., [start, end].
// Start time must be earlier than end time.
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

// IsZero returns true if both start time and end time are zero value.
func (r TimeRange) IsZero() bool {
	return r.start.IsZero() && r.end.IsZero()
}

// Duration returns the duration between start time and end time.
func (r TimeRange) Duration() time.Duration {
	return r.end.Sub(r.start)
}

// Contains returns true if the time is within this range.
func (r TimeRange) Contains(t time.Time) bool {
	return r.start.Equal(t) || r.end.Equal(t) || (r.start.Before(t) && t.Before(r.end))
}

// In returns true if the time is within the range.
// This is a synonym of TimeRange.Contains().
func In(t time.Time, r TimeRange) bool {
	return r.Contains(t)
}

// Before returns true if this range is earlier than the time.
func (r TimeRange) Before(t time.Time) bool {
	return r.end.Before(t)
}

// After returns true if this range is later than the time.
func (r TimeRange) After(t time.Time) bool {
	return r.start.After(t)
}

// Shift returns a TimeRange moved by the duration.
// If the duration is positive, this returns the later range.
// If the duration is negative, this returns the earlier range.
func (r TimeRange) Shift(d time.Duration) TimeRange {
	return New(r.start.Add(d), r.end.Add(d))
}

// ShiftDate returns a TimeRange moved by the duration in days.
// If the duration is positive, this returns the later range.
// If the duration is negative, this returns the earlier range.
func (r TimeRange) ShiftDate(years, months, days int) TimeRange {
	return New(r.start.AddDate(years, months, days), r.end.AddDate(years, months, days))
}

// Extend returns an extended TimeRange for the duration.
// If the duration is positive, this returns the longer range.
// If the duration is negative, this returns the shorter range.
func (r TimeRange) Extend(d time.Duration) TimeRange {
	return New(r.start, r.end.Add(d))
}

// ExtendDate returns an extended TimeRange for the duration in days.
// If the duration is positive, this returns the longer range.
// If the duration is negative, this returns the shorter range.
func (r TimeRange) ExtendDate(years, months, days int) TimeRange {
	return New(r.start, r.end.AddDate(years, months, days))
}
