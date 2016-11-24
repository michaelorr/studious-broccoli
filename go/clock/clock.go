package clock

import "fmt"

const testVersion = 4
const dayLen = 24
const hourLen = 60

type Clock struct {
	Hour, Minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}.normalizeClock()
}

func (c Clock) normalizeClock() Clock {
	// order is important here
	return c.normalizeMinutes().normalizeHours()
}

func (c Clock) normalizeMinutes() Clock {
	for c.Minute < 0 {
		c.Minute += hourLen
		c.Hour -= 1
	}
	for c.Minute >= hourLen {
		c.Minute -= hourLen
		c.Hour += 1
	}
	return c
}

func (c Clock) normalizeHours() Clock {
	c.Hour = c.Hour % dayLen
	if c.Hour < 0 {
		c.Hour += dayLen
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes
	return c.normalizeClock()
}
