# go-timerange [![go](https://github.com/int128/go-timerange/actions/workflows/go.yaml/badge.svg)](https://github.com/int128/go-timerange/actions/workflows/go.yaml)

This is a Go package of date time range.

## Getting Started

```shell
go get github.com/int128/go-timerange
```

```go
package example

import (
	"fmt"
	"time"

	"github.com/int128/go-timerange"
)

func CheckIfAvailable(desiredTime time.Time) {
	availableRange := timerange.New(
		time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		time.Date(2006, 1, 2, 15, 7, 5, 0, time.UTC),
	)
	if availableRange.Contains(desiredTime) {
		fmt.Printf("The reservation at %s is available.", desiredTime)
	}
}
```
