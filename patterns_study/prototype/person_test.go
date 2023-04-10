package person

import (
	"reflect"
	"testing"
)

func TestPersonCopy(t *testing.T) {
	p := &Person{name: "Jonny", age: 50, address: "London"}
	msg := p.Copy()
	if p == msg {
		t.Fatalf("Equal addresses: old: %v, new: %v", &p, &msg)
	}
	if !reflect.DeepEqual(p, msg) {
		t.Fatalf("Not equal fields: old: %v, new: %v", p, msg)
	}
}

func TestPersonCreate(t *testing.T) {
	want := &Person{name: "Jonny", age: 50, address: "London"}
	msg := Create("Jonny", 50, "London")
	if want == msg {
		t.Fatalf("Equal addresses: old: %v, new: %v", &want, &msg)
	}
	if !reflect.DeepEqual(want, msg) {
		t.Fatalf("Not equal fields: old: %v, new: %v", want, msg)
	}
}
