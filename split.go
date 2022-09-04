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
