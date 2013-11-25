# le

le is a simple Logentries writer for Go that sends logs asynchronously in the
background and handles automatically re-connecting, as needed.

Connection errors are logged to stderr.

**Security note:** Currently, `le` should only be used on trusted networks. It
does not yet use TLS. (Pull requests welcome.)

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

