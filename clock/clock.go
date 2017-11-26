package clock

import (
	"fmt"
	"time"
)

// Clock represents a 24h wall-clock for operations on time independent of date
type Clock struct {
	baseTime  time.Time
	clockTime time.Time
}

// Before answers if Clock-Time c occurs before oc
func (c *Clock) Before(oc *Clock) bool {
	return c.clockTime.Before(oc.clockTime)
}

// After answers if Clock-Time c occurs after oc
func (c *Clock) After(oc *Clock) bool {
	return c.clockTime.After(oc.clockTime)
}

// Diff returns the Absolute Value difference between times
func (c *Clock) Diff(oc *Clock) time.Duration {
	if c.After(oc) {
		return c.clockTime.Sub(oc.clockTime)
	}

	return oc.clockTime.Sub(c.clockTime)
}

func (c *Clock) String() string {
	h, m, s := c.baseTime.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// From returns a Clock based upon a Time
func From(t time.Time) *Clock {
	h, m, s := t.Clock()
	ct := time.Date(0, 0, 0, h, m, s, 0, t.Location())

	return &Clock{
		baseTime:  t,
		clockTime: ct,
	}
}
