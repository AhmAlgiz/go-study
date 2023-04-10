package streamPool

import "testing"

func TestStreamPoolData(t *testing.T) {
	data := NewData()
	data.SendPackage("12345")

	pool := NewPool(5, data)

	msg := (<-*pool).GetLastPackage()
	want := "12345"

	if msg != want {
		t.Fatalf("Uncorrect response, expected: %v, got: %v", want, msg)
	}
}
