package myTime

import (
	"fmt"
	"time"
)

type MyTime interface {
	Now() string
}

type MyTimeType int

const (
	DMYHMS MyTimeType = iota
	DMY
	HMS
)

type myTimeDMYHMS struct {
	MyTime
}

type myTimeDMY struct {
	MyTime
}

type myTimeHMS struct {
	MyTime
}

func (m *myTimeDMYHMS) Now() string {
	return time.Now().Format("02/01/2006 15:04:05")
}

func (m *myTimeDMY) Now() string {
	return time.Now().Format("02/01/2006")
}

func (m *myTimeHMS) Now() string {
	return time.Now().Format("15:04:05")
}

func CreateMyTime(m MyTimeType) (MyTime, error) {
	switch m {
	case DMYHMS:
		return &myTimeDMYHMS{}, nil
	case DMY:
		return &myTimeDMY{}, nil
	case HMS:
		return &myTimeHMS{}, nil
	default:
		return nil, fmt.Errorf("Wrong type.")
	}
}
