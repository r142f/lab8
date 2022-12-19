package clock

import "time"

type Clock interface {
	Now() time.Time
}

type NormalClock struct{}

func (nc *NormalClock) Now() time.Time {
	return time.Now()
}

type SetableClock struct {
	now time.Time
}

func (sc *SetableClock) SetNow(now time.Time) {
	sc.now = now
}

func (sc *SetableClock) Now() time.Time {
	return sc.now
}
