package squaresSum

import "testing"

func TestSquaresSum(t *testing.T) {
	msg := sum(sqr(gen(1, 2, 3, 4, 5)))
	want := 55
	if msg != want {
		t.Fatalf("Uncorrect answer. Expected: %v, got: %v", want, msg)
	}
}
