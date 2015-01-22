package le

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"
)

var (
	Address  = "data.logentries.com:10000"
	newlines = []byte{'\n'}
)

// Writer implements the io.WriteCloser interface for sending token-based logs
// to Logentries.
type Writer struct {
	Token string
	conn  net.Conn
}

// NewWriter returns a new Writer with a given Logentries token. If the token
// is invalid, Logentries will ignore all submitted logs.
func NewWriter(token string) *Writer {
	return &Writer{Token: token}
}

// Write sends lines of text to Logentries with our associated token. If the
// write fails at any point, we return 0 regardless of how many bytes were
// actually written.
func (w *Writer) Write(p []byte) (n int, err error) {
	if w.conn == nil {
		err = w.connect()
		if err != nil {
			return 0, err
		}
	}

	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		_, err = fmt.Fprintf(w.conn, "%s %s\n", w.Token, scanner.Bytes())
		if err != nil {
			w.Close()
			return 0, err
		}
	}

	return len(p), scanner.Err()
}

func (w *Writer) Close() (err error) {
	if w.conn != nil {
		err = w.conn.Close()
		w.conn = nil
	}
	return err
}

func (w *Writer) connect() (err error) {
	w.conn, err = net.DialTimeout("tcp", Address, time.Second)
	return err
}
