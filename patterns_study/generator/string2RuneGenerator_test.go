package generator

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	ch := String2RuneGenerator("qw12аб,")
	msg := make([]rune, 0, 7)
	for val := range ch {
		msg = append(msg, val)
	}
	want := [7]rune{'q', 'w', '1', '2', 'а', 'б', ','}

	for i, v := range msg {
		if v != want[i] {
			t.Fatalf("Expected %v, recieved %v", want, msg)
		}
	}
}
