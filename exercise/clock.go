package exercise

import "fmt"

type Clock struct {
	minutes int
}

const (
	dayInMin = 1440
)

func New(h, m int) Clock {
	return Clock{
		minutes: (( h * 60 + m) % dayInMin + dayInMin) % dayInMin,
	}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.minutes + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes / 60, c.minutes % 60)
}