package timerange_test

import (
	"fmt"
	"time"

	"github.com/int128/go-timerange"
)

func ExampleTimeRange_Contain() {
	desiredTime := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
	availableRange := timerange.TimeRange{
		Start: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		End:   time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	}
	if availableRange.Contain(desiredTime) {
		fmt.Printf("The reservation at %s is available.", desiredTime)
	}
	// output: The reservation at 2006-01-02 15:06:00 +0000 UTC is available.
}
