package timerange_test

import (
	"fmt"
	"time"

	"github.com/int128/go-timerange"
)

func ExampleIn() {
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

func ExampleTimeRange_Contains() {
	desiredTime := time.Date(2006, 1, 2, 15, 6, 0, 0, time.UTC)
	availableRange := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	if availableRange.Contains(desiredTime) {
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

func ExampleTimeRange_Shift() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	fmt.Print(r.Shift(15 * time.Minute))
	// output: [2006-01-02T15:19:05Z, 2006-01-02T15:22:05Z]
}

func ExampleTimeRange_Extend() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	fmt.Print(r.Extend(15 * time.Minute))
	// output: [2006-01-02T15:04:05Z, 2006-01-02T15:22:05Z]
}

func ExampleTimeRange_Split() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	for _, t := range r.Split(1 * time.Minute) {
		fmt.Println(t)
	}
	// output:
	// 2006-01-02 15:04:05 +0000 UTC
	// 2006-01-02 15:05:05 +0000 UTC
	// 2006-01-02 15:06:05 +0000 UTC
	// 2006-01-02 15:07:05 +0000 UTC
}

func ExampleTimeRange_SplitIterator() {
	r := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	iter := r.SplitIterator(1 * time.Minute)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
	// output:
	// 2006-01-02 15:04:05 +0000 UTC
	// 2006-01-02 15:05:05 +0000 UTC
	// 2006-01-02 15:06:05 +0000 UTC
	// 2006-01-02 15:07:05 +0000 UTC
}
