package timerange_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/int128/go-timerange"
)

func TestTimeRange_Split(t *testing.T) {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)

	t.Run("inclusive last", func(t *testing.T) {
		got := r.Split(30 * time.Second)
		want := []time.Time{
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 4, 35, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 5, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 5, 35, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 35, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("want != got\n%s", diff)
		}
	})

	t.Run("exclusive last", func(t *testing.T) {
		got := r.Split(2*time.Minute + 30*time.Second)
		want := []time.Time{
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			time.Date(2006, 1, 2, 15, 6, 35, 0, time.UTC),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("want != got\n%s", diff)
		}
	})

	t.Run("span is longer than range", func(t *testing.T) {
		got := r.Split(1 * time.Hour)
		want := []time.Time{
			time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("want != got\n%s", diff)
		}
	})
}
