package car

import (
	"testing"
)

func TestNewCarDefault(t *testing.T) {
	msg := NewCar()
	want := &Car{
		Model:      "",
		Brand:      "",
		HorsePower: 0,
		Length:     0,
		Width:      0,
	}
	if msg.Model != want.Model || msg.Brand != want.Brand ||
		msg.HorsePower != want.HorsePower || msg.Length != want.Length ||
		msg.Width != want.Width {
		t.Fatalf(`NewCar() = %q, want match for %q`, msg, want)
	}
}

func TestNewCarCustom(t *testing.T) {
	msg := NewCar(Model("XC90"), Brand("Volvo"), HorsePower(255), Length(5000), Width(1950))
	want := &Car{
		Model:      "XC90",
		Brand:      "Volvo",
		HorsePower: 255,
		Length:     5000,
		Width:      1950,
	}
	if msg.Model != want.Model || msg.Brand != want.Brand ||
		msg.HorsePower != want.HorsePower || msg.Length != want.Length ||
		msg.Width != want.Width {
		t.Fatalf(`NewCar() = %q, want match for %q`, msg, want)
	}
}
