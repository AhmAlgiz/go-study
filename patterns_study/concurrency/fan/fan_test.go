package fan

import "testing"

func TestFan(t *testing.T) {
	g := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	want := 110
	ch1, ch2 := compute(g), compute(g)
	var sum int
	for n := range merge(ch1, ch2) {
		sum += n
	}
	if sum != want {
		t.Fatalf("Uncorrect answer. Want: %v, got: %v", want, sum)
	}
}
