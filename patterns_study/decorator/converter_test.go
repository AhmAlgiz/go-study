package converter

import "testing"

func TestConverterRublesToDollar(t *testing.T) {
	base := 10
	days := 10
	rate := 80
	fn := ConvertRubblesToDollars(CalculateCostRubles, rate)
	msg := fn(base, days)
	want := base * days * rate
	if msg != want {
		t.Fatalf("Uncorrect response, expected: %v, got: %v", want, msg)
	}

}
