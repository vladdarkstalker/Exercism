package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    hour int
    min int
}

func New(h, m int) Clock {
    allTime := (h*60 + m) % (24 * 60)
    if allTime < 0 {
        allTime += 24 * 60
    }
	return Clock{hour: allTime/60, min: allTime%60}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.min+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.min-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}
