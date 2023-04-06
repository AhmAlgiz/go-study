package config

import (
	"testing"
	"time"
)

func TestConfigDefault(t *testing.T) {
	msg := GetConfig()
	want := GetConfig()
	if msg.InitTime.Equal(want.InitTime) != true {
		t.Fatalf("Mismatched config parameter: InitTime. Want: %v, recieved: %v", want.InitTime, msg.InitTime)
	}
}

func TestConfigSleep(t *testing.T) {
	msg := GetConfig()
	time.Sleep(time.Second * 2)
	want := GetConfig()
	if msg.InitTime.Equal(want.InitTime) != true {
		t.Fatalf("Mismatched config parameter: InitTime. Want: %v, recieved: %v", want.InitTime, msg.InitTime)
	}
}
