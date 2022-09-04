package timerange_test

import (
	"testing"
	"time"

	"github.com/int128/go-timerange"
)

func TestIntersect(t *testing.T) {
	t.Run("same range", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, a)
		want := a
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("a.start < b < a.end", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		b := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, b)
		want := b
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("b.start < a.start < b.end < a.end", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		b := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, b)
		want := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("a.start < b.start < a.end < b.end", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		)
		b := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, b)
		want := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("a < b", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		b := timerange.New(
			time.Date(2006, 1, 2, 16, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 16, 6, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, b)
		var want timerange.TimeRange
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("a > b", func(t *testing.T) {
		a := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		b := timerange.New(
			time.Date(2006, 1, 2, 14, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 14, 6, 5, 0, time.UTC),
		)
		got := timerange.Intersect(a, b)
		var want timerange.TimeRange
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
}
