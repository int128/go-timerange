# go-timerange [![go](https://github.com/int128/go-timerange/actions/workflows/go.yaml/badge.svg)](https://github.com/int128/go-timerange/actions/workflows/go.yaml)

This is a Go package to handle a time range.

## Getting Started

To install this package,

```shell
go get github.com/int128/go-timerange
```

Here is an example.

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
	if timerange.In(desiredTime, availableRange) {
		fmt.Printf("The reservation at %s is available.", desiredTime)
	}
}
```
