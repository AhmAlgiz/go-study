package concurrency

import (
	"strconv"
	"testing"
)

func TestQuery(t *testing.T) {
	conns := make([]Connection, 0, 3)
	for i := 0; i < 3; i++ {
		conns = append(conns, Connection{name: strconv.Itoa(i)})
	}
	msg := Query(conns, "test")
	if msg == "" {
		t.Fatalf("Empty or uncorrect response. Got: %v", msg)
	}
}
