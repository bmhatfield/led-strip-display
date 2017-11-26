package clock

import (
	"fmt"
	"time"
)

// Clock represents a 24h wall-clock for operations on time independent of date
type Clock struct {
	baseTime time.Time
}

// Before answers if Clock-Time c occurs before oc
func (c *Clock) Before(oc *Clock) bool {
	h, m, s := c.baseTime.Clock()
	oh, om, os := oc.baseTime.Clock()

	if oh > h {
		return true
	} else if oh >= h && om > m {
		return true
	} else if oh >= h && om >= m && os > s {
		return true
	}

	return false
}

// After answers if Clock-Time c occurs after oc
func (c *Clock) After(oc *Clock) bool {
	h, m, s := c.baseTime.Clock()
	oh, om, os := oc.baseTime.Clock()

	if h > oh {
		return true
	} else if h >= oh && m > om {
		return true
	} else if h >= oh && m >= om && s > os {
		return true
	}

	return false
}

// Diff returns the Absolute Value difference between times
func (c *Clock) Diff(oc *Clock) time.Duration {
	h, m, s := c.baseTime.Clock()
	oh, om, os := oc.baseTime.Clock()

	if c.After(oc) {
		hours := time.Duration(h-oh) * time.Hour
		minutes := time.Duration(m-om) * time.Minute
		seconds := time.Duration(s-os) * time.Second

		return time.Duration(hours + minutes + seconds)
	}

	hours := time.Duration(oh-h) * time.Hour
	minutes := time.Duration(om-m) * time.Minute
	seconds := time.Duration(os-s) * time.Second

	return time.Duration(hours + minutes + seconds)
}

func (c *Clock) String() string {
	h, m, s := c.baseTime.Clock()
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// From returns a Clock based upon a Time
func From(t time.Time) *Clock {
	return &Clock{baseTime: t}
}
