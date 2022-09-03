package timerange

import (
	"testing"
	"time"
)

func TestTimeRange_String(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}
	t.Logf("String() = %s", r)
	got := r.String()
	want := "[2006-01-02T15:04:05Z, 2006-01-02T15:07:05Z]"
	if got != want {
		t.Errorf("String() wants %s but was %s", want, got)
	}
}

func TestTimeRange_Equal(t *testing.T) {
	a := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}
	b := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
	}
	t.Run("a == a", func(t *testing.T) {
		got := a.Equal(a)
		const want = true
		if want != got {
			t.Errorf("Equal() wants %v but was %v", want, got)
		}
	})
	t.Run("a != b", func(t *testing.T) {
		got := a.Equal(b)
		const want = false
		if want != got {
			t.Errorf("Equal() wants %v but was %v", want, got)
		}
	})
}

func TestTimeRange_IsValid(t *testing.T) {
	t.Run("Start < End", func(t *testing.T) {
		r := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		got := r.IsValid()
		const want = true
		if want != got {
			t.Errorf("IsValid() wants %v but was %v", want, got)
		}
	})
	t.Run("Start == End", func(t *testing.T) {
		r := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		}
		got := r.IsValid()
		const want = true
		if want != got {
			t.Errorf("IsValid() wants %v but was %v", want, got)
		}
	})
	t.Run("Start > End", func(t *testing.T) {
		r := TimeRange{
			Start: time.Date(2006, 1, 3, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		got := r.IsValid()
		const want = false
		if want != got {
			t.Errorf("IsValid() wants %v but was %v", want, got)
		}
	})
}

func TestTimeRange_Duration(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}
	got := r.Duration()
	want := 3 * time.Minute
	if got != want {
		t.Errorf("Duration() wants %s but was %s", want, got)
	}
}

func TestTimeRange_Contains(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = false
		if want != got {
			t.Errorf("Contains() wants %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("Contains() wants %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("Contains() wants %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.Contains(point)
		const want = true
		if want != got {
			t.Errorf("Contains() wants %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.Contains(point)
		const want = false
		if want != got {
			t.Errorf("Contains() wants %v but was %v", want, got)
		}
	})
}

func TestTimeRange_Before(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("Before() wants %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("Before() wants %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("Before() wants %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.Before(point)
		const want = false
		if want != got {
			t.Errorf("Before() wants %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.Before(point)
		const want = true
		if want != got {
			t.Errorf("Before() wants %v but was %v", want, got)
		}
	})
}

func TestTimeRange_After(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}

	t.Run("point is before range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 3, 0, 0, time.UTC)
		got := r.After(point)
		const want = true
		if want != got {
			t.Errorf("After() wants %v but was %v", want, got)
		}
	})
	t.Run("point is left edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("After() wants %v but was %v", want, got)
		}
	})
	t.Run("point is in range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("After() wants %v but was %v", want, got)
		}
	})
	t.Run("point is right edge of range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("After() wants %v but was %v", want, got)
		}
	})
	t.Run("point is after range", func(t *testing.T) {
		point := time.Date(2006, 1, 2, 15, 8, 0, 0, time.UTC)
		got := r.After(point)
		const want = false
		if want != got {
			t.Errorf("After() wants %v but was %v", want, got)
		}
	})
}

func TestIntersect(t *testing.T) {
	t.Run("same range", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		got := Intersect(a, a)
		want := a
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
	t.Run("a.Start < b < a.End", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		b := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		}
		got := Intersect(a, b)
		want := b
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
	t.Run("b.Start < a.Start < b.End < a.End", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		b := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		}
		got := Intersect(a, b)
		want := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		}
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
	t.Run("a.Start < b.Start < a.End < b.End", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		}
		b := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		got := Intersect(a, b)
		want := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
		}
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
	t.Run("a < b", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		b := TimeRange{
			Start: time.Date(2006, 1, 2, 16, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 16, 6, 5, 0, time.UTC),
		}
		got := Intersect(a, b)
		want := TimeRange{}
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
	t.Run("a > b", func(t *testing.T) {
		a := TimeRange{
			Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		b := TimeRange{
			Start: time.Date(2006, 1, 2, 14, 5, 5, 0, time.UTC),
			End:   time.Date(2006, 1, 2, 14, 6, 5, 0, time.UTC),
		}
		got := Intersect(a, b)
		want := TimeRange{}
		if !want.Equal(got) {
			t.Errorf("Intersect(): want %v != got %v", want, got)
		}
	})
}
