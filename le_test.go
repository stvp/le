package le

import (
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	writer := NewWriter(os.Getenv("TOKEN"))
	_, err := writer.Write([]byte("Testing\nA\nB\nC"))
	if err != nil {
		t.Error(err)
	}
	_, err = writer.Write([]byte("Another line\n\nBlank lines are ignored"))
	if err != nil {
		t.Error(err)
	}
}
