# le

le is a [logentries](https://logentries.com) writer for Go that sends logs
asynchronously. All connection errors are logged to stderr.

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
  logger, err := le.New(LOG_TOKEN, 500)
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(logger, "%d percent effort", 110)

  logger.Wait() // Wait for queued logs to be sent
}
```

## Development

    TOKEN=<testing_token_here> go test

