package ledstrip

import (
	"log"
	"time"

	"github.com/bmhatfield/led-strip-display/clock"
	"github.com/bmhatfield/sunset/sunset"
)

// DisplayTimer automatically manages the state of a RenderServer
// based upon the time of day and sunset
type DisplayTimer struct {
	sunset *clock.Clock
	server *RenderServer
}

// Run starts the switcher goroutine
func (dt *DisplayTimer) Run() {
	go dt.switcher()
}

func (dt *DisplayTimer) updater() {
	// Set up channel to receive daily updated sunset times
	sunsetTimes := sunset.AutoUpdatingTime()

	// Perform initial recieve of sunset times synchronously
	times := <-sunsetTimes

	// Configure the DisplayTimer with the current SunsetTime
	dt.sunset = clock.From(times.Sunset.Local())

	// Continuously watch for updated sunset times
	for {
		// Blocking receive
		times := <-sunsetTimes

		// Perform updates to state
		dt.sunset = clock.From(times.Sunset.Local())
	}
}

func (dt *DisplayTimer) switcher() {
	for {
		desiredState := dt.shouldBeOn()

		if dt.server.Enabled != desiredState {
			log.Printf("Setting RenderServer to enabled: %t\n", desiredState)

			dt.server.Enabled = desiredState
		}

		// Don't run this in a tight loop
		time.Sleep(30 * time.Second)
	}
}

func (dt *DisplayTimer) shouldBeOn() bool {
	var afterSunset bool

	now := clock.From(time.Now())
	disabled := clock.From(time.Date(0, 0, 0, 1, 0, 0, 0, time.Local))

	if dt.sunset != nil {
		afterSunset = dt.sunset.Before(now)
	}

	return afterSunset || disabled.After(now)
}

// NewDisplayTimer configures and returns a DisplayTimer
func NewDisplayTimer(server *RenderServer) *DisplayTimer {
	timer := &DisplayTimer{}
	timer.server = server

	go timer.updater()

	return timer
}
