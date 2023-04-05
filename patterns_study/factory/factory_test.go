package myTime

import (
	"testing"
	"time"
)

func TestMyTimeDMYHMS(t *testing.T) {
	mt, err := CreateMyTime(DMYHMS)
	if err != nil {
		t.Fatalf(err.Error())
	}
	msg := mt.Now()
	if _, err := time.Parse("02/01/2006 15:04:05", msg); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMyTimeDMY(t *testing.T) {
	mt, err := CreateMyTime(DMY)
	if err != nil {
		t.Fatalf(err.Error())
	}
	msg := mt.Now()
	if _, err := time.Parse("02/01/2006", msg); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMyTimeHMS(t *testing.T) {
	mt, err := CreateMyTime(HMS)
	if err != nil {
		t.Fatalf(err.Error())
	}
	msg := mt.Now()
	if _, err := time.Parse("15:04:05", msg); err != nil {
		t.Fatalf(err.Error())
	}
}
