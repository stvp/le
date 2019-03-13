# le

`le` is a Go package that provides a basic goroutine-safe `io.Writer` interface
for sending logs to Logentries using basic Token TCP auth (*no TLS*).

## Documentation

[API docs on godoc.org](http://godoc.org/github.com/stvp/le)

## Basic usage

```go
package main

import (
  "fmt"
  "github.com/stvp/le"
)

const (
  LOG_TOKEN = "47329628-ab93-4418-8265-9acdb0333248"
)

func main() {
  logger := le.New(LOG_TOKEN)

  fmt.Fprintf(logger, "%d percent effort", 110)
}
```

