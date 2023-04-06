package starship

import (
	"reflect"
	"testing"
)

func TestStarshipDefault(t *testing.T) {
	msg := NewStarshipBuilder().Build()
	want := &Starship{
		Name:        "",
		MaxSpeed:    0,
		FuelStorage: 0,
		Mass:        0,
		Capacity:    0,
	}
	if !reflect.DeepEqual(msg, want) {
		t.Fatalf("Mismatch of structures, want %v, recieved: %v", &want, &msg)
	}
}

func TestStarshipCustom(t *testing.T) {
	msg := NewStarshipBuilder().SetName("Nemo").SetMaxSpeed(500).SetFuelStorage(1000).SetMass(50000).SetCapacity(90000).Build()
	want := &Starship{
		Name:        "Nemo",
		MaxSpeed:    500,
		FuelStorage: 1000,
		Mass:        50000,
		Capacity:    90000,
	}
	if !reflect.DeepEqual(msg, want) {
		t.Fatalf("Mismatch of structures, want %v, recieved: %v", want, msg)
	}
}
