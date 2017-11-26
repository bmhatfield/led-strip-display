package ledstrip

import (
	"log"
	"time"

	"github.com/bmhatfield/led-strip-display/clock"
	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/bmhatfield/sunset/sunset"
)

// ShouldBeOn reports whether the current time is "after dark" or not
func ShouldBeOn(sunsetTime *clock.Clock) bool {
	now := clock.From(time.Now())
	disabled := clock.From(time.Date(0, 0, 0, 1, 0, 0, 0, time.Local))

	return sunsetTime.Before(now) || disabled.After(now)
}

// RenderServer is an async server for rendering frames
type RenderServer struct {
	FrameQueue chan frame.HexGRBFrame
	strip      Strip
	ticker     *time.Ticker
	enabled    bool
	sunset     *clock.Clock
}

func (r *RenderServer) scheduler() {
	// Set up channel to receive daily updated sunset times
	sunsetTimes := sunset.AutoUpdatingTime()

	// Perform initial setup
	times := <-sunsetTimes
	r.sunset = clock.From(times.Sunset.Local())
	go r.renderSwitch()

	// Continuously watch for updated sunset times
	for {
		// Blocking receive
		times := <-sunsetTimes

		// Perform updates to state
		r.sunset = clock.From(times.Sunset.Local())
	}
}

func (r *RenderServer) renderSwitch() {
	for {
		r.enabled = ShouldBeOn(r.sunset)
		time.Sleep(10 * time.Second)
	}
}

func (r *RenderServer) render() {
	currentFrame := frame.HexGRBFrame{}

	for range r.ticker.C {
		if r.enabled {
			select {
			case frame := <-r.FrameQueue:
				log.Println("Rendering received frame...")
				err := r.strip.Render(frame)

				if err != nil {
					log.Println("Unable to render frame:", err)
				}

				currentFrame = frame

			default:
				continue
			}
		} else {
			now := clock.From(time.Now())

			if r.sunset != nil && r.sunset.After(now) {
				log.Println("Strip is disabled, sleeping until", r.sunset)

				r.strip.Render(frame.HexGRBFrame{})
				time.Sleep(r.sunset.Diff(now))

				log.Println("Waking up strip...")
				r.strip.Render(currentFrame)
			}
		}
	}
}

// NewRenderServer returns the channel you can send to a RenderServer on
func NewRenderServer(strip Strip) *RenderServer {
	// Initialize the strip for our purposes
	strip.Init(12, 150, 200)

	// Create the new RenderServer struct
	server := &RenderServer{strip: strip}

	// Create our queue of frames to render
	server.FrameQueue = make(chan frame.HexGRBFrame, 120)

	// Create a timing mechanism to ensure consistent FPS
	server.ticker = time.NewTicker(40 * time.Millisecond) // 25 FPS

	// The scheduler tells the program when to enable or disable the strip
	go server.scheduler()

	// Run the renderer
	go server.render()

	// Return the
	return server
}
