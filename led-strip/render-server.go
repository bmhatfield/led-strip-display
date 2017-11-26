package ledstrip

import (
	"log"
	"time"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/bmhatfield/sunset/sunset"
)

// ShouldBeOff reports whether the current time is "after dark" or not
func ShouldBeOff(sunsetTime time.Time) bool {
	disableHour, disableMinute := 1, 0
	now := time.Now()

	return sunsetTime.After(now) && now.Hour() >= disableHour && now.Minute() >= disableMinute
}

// RenderServer is an async server for rendering frames
type RenderServer struct {
	FrameQueue  chan frame.HexGRBFrame
	strip       Strip
	ticker      *time.Ticker
	enabled     bool
	nextEnabled time.Time
}

func (r *RenderServer) scheduler() {
	var sunsetTime time.Time
	sunsetTimes := sunset.AutoUpdatingTime()

	for {
		select {
		case times := <-sunsetTimes:
			sunsetTime = times.Sunset.Local()
		default:
		}

		if ShouldBeOff(sunsetTime) {
			r.nextEnabled = sunsetTime
			r.enabled = false
		} else {
			r.enabled = true
		}

		time.Sleep(30 * time.Second)
	}
}

func (r *RenderServer) render() {
	var currentFrame frame.HexGRBFrame

	for range r.ticker.C {
		if r.enabled {
			select {
			case frame := <-r.FrameQueue:
				err := r.strip.Render(frame)

				if err != nil {
					log.Println("Unable to render frame:", err)
				}

				currentFrame = frame

			default:
				continue
			}
		} else {
			if r.nextEnabled.After(time.Now()) {
				r.strip.Render(frame.HexGRBFrame{})
				time.Sleep(r.nextEnabled.Sub(time.Now()))
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
