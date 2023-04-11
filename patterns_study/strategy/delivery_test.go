package delivery

import "testing"

func TestDeliveryCar(t *testing.T) {
	msg := (DeliveryType{Car{}}).Calculator.Calculate(10, 10, 10)
	want := 500
	if msg != want {
		t.Fatalf("Uncorrect result: expected %v, got %v", want, msg)
	}
}

func TestDeliveryBicycle(t *testing.T) {
	msg := (DeliveryType{Bicycle{}}).Calculator.Calculate(10, 10, 10)
	want := 200
	if msg != want {
		t.Fatalf("Uncorrect result: expected %v, got %v", want, msg)
	}
}

func TestDeliveryWalker(t *testing.T) {
	msg := (DeliveryType{Walker{}}).Calculator.Calculate(10, 10, 10)
	want := 100
	if msg != want {
		t.Fatalf("Uncorrect result: expected %v, got %v", want, msg)
	}
}
