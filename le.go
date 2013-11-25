package le

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	LOGENTRIES_ADDRESS = "data.logentries.com:10000"
)

type Writer struct {
	Token     string
	conn      net.Conn
	channel   chan []byte
	waitGroup sync.WaitGroup
}

// New returns a new Writer with a given Logentries token and buffer size. If
// the token is invalid, Logentries will ignore all submitted logs. The buffer
// size is the maximum number of writes that will be queued for sending before
// the Writer begins rejecting new writes.
func New(token string, buffer int) (w *Writer, err error) {
	w = &Writer{
		Token:   token,
		channel: make(chan []byte, buffer),
	}

	if err = w.connect(); err != nil {
		return nil, err
	}

	w.start()

	return w, nil
}

// Write queues lines of text for sending to Logentries. It returns an error
// only if the write buffer is full.
func (w *Writer) Write(p []byte) (n int, err error) {
	if len(w.channel) >= cap(w.channel) {
		return 0, fmt.Errorf("Buffer is full")
	}

	w.channel <- p
	w.waitGroup.Add(1)

	return len(p), nil
}

// Wait will block until all queued writes have been sent to Logentries.
func (w *Writer) Wait() {
	w.waitGroup.Wait()
}

func (w *Writer) start() {
	go func() {
		for {
			w.write(<-w.channel)
			w.waitGroup.Done()
		}
	}()
}

func (w *Writer) connect() (err error) {
	if w.conn != nil {
		w.conn.Close()
		w.conn = nil
	}

	w.conn, err = net.DialTimeout("tcp", LOGENTRIES_ADDRESS, time.Second)
	return err
}

func (w *Writer) write(lines []byte) {
	if w.conn == nil {
		if err := w.connect(); err != nil {
			connectFailed(err)
			return
		}
	}

	for _, line := range bytes.Split(lines, []byte{'\n'}) {
		// Logentries ignores blank lines
		if len(line) == 0 {
			continue
		}

		_, err := fmt.Fprintf(w.conn, "%s %s\n", w.Token, line)
		if err != nil {
			writeFailed(err)
			if err = w.connect(); err != nil {
				connectFailed(err)
				continue
			}
			fmt.Fprintf(w.conn, "%s %s\n", w.Token, line)
		}
	}
}

func connectFailed(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: Couldn't connect to %s: %s", LOGENTRIES_ADDRESS, err.Error())
}

func writeFailed(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: Couldn't write to %s: %s", LOGENTRIES_ADDRESS, err.Error())
}
