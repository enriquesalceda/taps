package clk

import "time"

type Clk interface {
	Now() time.Time
}

type Clock struct{}

func NewClock() *Clock {
	return &Clock{}
}

func (c Clock) Now() time.Time {
	return time.Now()
}
