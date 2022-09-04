package timerange_test

import (
	"testing"
	"time"

	"github.com/int128/go-timerange"
)

func TestNew(t *testing.T) {
	t.Run("start < end", func(t *testing.T) {
		r := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		got := r.String()
		want := "[2006-01-02T15:04:05Z, 2006-01-02T15:07:05Z]"
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("start == end", func(t *testing.T) {
		r := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		)
		got := r.String()
		want := "[2006-01-02T15:04:05Z, 2006-01-02T15:04:05Z]"
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("start > end", func(t *testing.T) {
		r := timerange.New(
			time.Date(2006, 1, 3, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		got := r.IsZero()
		const want = true
		if want != got {
			t.Errorf("want %v but was %v (r=%s)", want, got, r)
		}
	})
}

func TestFrom(t *testing.T) {
	r := timerange.From(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		15*time.Minute,
	)
	got := r.String()
	want := "[2006-01-02T15:04:05Z, 2006-01-02T15:19:05Z]"
	if want != got {
		t.Errorf("want %v but was %v", want, got)
	}
}

func TestUntil(t *testing.T) {
	r := timerange.Until(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		15*time.Minute,
	)
	got := r.String()
	want := "[2006-01-02T14:49:05Z, 2006-01-02T15:04:05Z]"
	if want != got {
		t.Errorf("want %v but was %v", want, got)
	}
}

func TestTimeRange_Start(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	got := r.Start()
	want := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	if got != want {
		t.Errorf("want %s but was %s", want, got)
	}
}

func TestTimeRange_End(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	got := r.End()
	want := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
	if got != want {
		t.Errorf("want %s but was %s", want, got)
	}
}

func TestTimeRange_String(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	t.Logf("String() = %s", r)
	got := r.String()
	want := "[2006-01-02T15:04:05Z, 2006-01-02T15:07:05Z]"
	if got != want {
		t.Errorf("want %s but was %s", want, got)
	}
}

func TestTimeRange_Equal(t *testing.T) {
	a := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	b := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
	)
	t.Run("a == a", func(t *testing.T) {
		got := a.Equal(a)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("a != b", func(t *testing.T) {
		got := a.Equal(b)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
}

func TestTimeRange_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var r timerange.TimeRange
		got := r.IsZero()
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("non-zero", func(t *testing.T) {
		r := timerange.New(
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		)
		got := r.IsZero()
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
}

func TestTimeRange_Duration(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	got := r.Duration()
	want := 3 * time.Minute
	if got != want {
		t.Errorf("want %s but was %s", want, got)
	}
}

func TestTimeRange_Contains(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
}

func TestTimeRange_Before(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.Before(point)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
}

func TestTimeRange_After(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.After(point)
		const want = true
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("want %v but was %v", want, got)
		}
	})
}

func TestTimeRange_Shift(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	t.Run("past", func(t *testing.T) {
		got := r.Shift(-15 * time.Minute)
		want := timerange.New(
			time.Date(2006, 1, 2, 14, 50, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 14, 52, 5, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("future", func(t *testing.T) {
		got := r.Shift(15 * time.Minute)
		want := timerange.New(
			time.Date(2006, 1, 2, 15, 20, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 22, 5, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
}

func TestTimeRange_Extend(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	t.Run("shorter", func(t *testing.T) {
		got := r.Extend(-15 * time.Second)
		want := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 50, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
	t.Run("longer", func(t *testing.T) {
		got := r.Extend(15 * time.Second)
		want := timerange.New(
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 20, 0, time.UTC),
		)
		if !want.Equal(got) {
			t.Errorf("want %v != got %v", want, got)
		}
	})
}

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
