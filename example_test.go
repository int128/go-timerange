package timerange_test

import (
	"fmt"
	"time"

	"github.com/int128/go-timerange"
)

func ExampleTimeRange_Contains() {
	desiredTime := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
	availableRange := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	if timerange.In(desiredTime, availableRange) {
		fmt.Printf("The reservation at %s is available.", desiredTime)
	}
	// output: The reservation at 2006-01-02 15:06:00 +0000 UTC is available.
}

func ExampleNew() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	fmt.Print(r)
	// output: [2006-01-02T15:04:05Z, 2006-01-02T15:07:05Z]
}

func ExampleFrom() {
	start := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	r := timerange.From(start, 15*time.Minute)
	fmt.Print(r)
	// output: [2006-01-02T15:04:05Z, 2006-01-02T15:19:05Z]
}

func ExampleUntil() {
	end := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	r := timerange.Until(end, 15*time.Minute)
	fmt.Print(r)
	// output: [2006-01-02T14:49:05Z, 2006-01-02T15:04:05Z]
}

func ExampleTimeRange_String() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	fmt.Print(r)
	// output: [2006-01-02T15:04:05Z, 2006-01-02T15:07:05Z]
}
