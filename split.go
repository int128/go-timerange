package timerange

import "time"

// Split returns an array of time points within this range.
// If span is longer than this range, it returns only start time.
func (r TimeRange) Split(span time.Duration) []time.Time {
	var points []time.Time
	for t := r.start; t.Equal(r.end) || t.Before(r.end); t = t.Add(span) {
		points = append(points, t)
	}
	return points
}

// SplitIterator returns an iterator for time points within this range.
// If span is longer than this range, it returns only start time.
//
// Consider this method instead of Split, if the result is too long.
func (r TimeRange) SplitIterator(span time.Duration) SplitIterator {
	return &splitIterator{timeRange: r, span: span}
}

// SplitIterator is an iterator to retrieve time points within a range.
type SplitIterator interface {
	// HasNext returns true if the next time is within the range.
	HasNext() bool

	// Next returns the next time.
	// It the next time is not in the range, this returns a zero value.
	Next() time.Time
}

type splitIterator struct {
	next      time.Time
	timeRange TimeRange
	span      time.Duration
}

func (s *splitIterator) HasNext() bool {
	return s.next.Equal(s.timeRange.end) || s.next.Before(s.timeRange.end)
}

func (s *splitIterator) Next() time.Time {
	if !s.HasNext() {
		return time.Time{}
	}
	current := s.next
	if current.IsZero() {
		current = s.timeRange.start
	}
	s.next = current.Add(s.span)
	return current
}
