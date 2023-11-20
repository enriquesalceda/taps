package clk

import "time"

type FrozenClock struct {
	frozenTime time.Time
}

func (fc FrozenClock) Now() time.Time {
	return fc.frozenTime
}

func NewFrozenClock(t time.Time) FrozenClock {
	return FrozenClock{frozenTime: t}
}
