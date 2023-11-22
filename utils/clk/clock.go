package clk

import "time"

type Clock interface {
	Now() time.Time
}

type Clk struct{}

func NewClock() *Clk {
	return &Clk{}
}

func (c Clk) Now() time.Time {
	return time.Now()
}
