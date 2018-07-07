package ledstrip

import (
	"log"
	"time"

	"github.com/bmhatfield/led-strip-display/frame"
)

// RenderServer is an async server for rendering frames
type RenderServer struct {
	Enabled    bool
	FrameQueue chan frame.HexGRB
	strip      Strip
	ticker     *time.Ticker
	blanked    bool
	lastframe  frame.HexGRB
}

func (r *RenderServer) render() {
	for range r.ticker.C {
		if r.Enabled {
			if r.blanked {
				select {
				case r.FrameQueue <- r.lastframe:
					log.Println("Re-rendering last frame before blanking")
				default:
					log.Println("Unable to re-render last frame before blanking; action would block")
				}

				r.blanked = false
			}

			select {
			case frame := <-r.FrameQueue:
				err := r.strip.Render(frame)

				if err != nil {
					log.Println("Unable to render frame:", err)
				} else {
					r.lastframe = frame
				}
			default:
				continue
			}
		} else {
			if !r.blanked {
				log.Println("Rendering blank frame because strip is disabled")

				// Blank out the strip by rendering an empty frame
				r.strip.Render(frame.HexGRB{})

				// Mark the server as blanked to avoid re-rendering
				// blank frames every tick
				r.blanked = true
			}

			// Slow down render attempts to 0.5fps when strip is disabled.
			time.Sleep(2 * time.Second)
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
	server.FrameQueue = make(chan frame.HexGRB, 120)

	// Create a timing mechanism to ensure consistent FPS
	server.ticker = time.NewTicker(40 * time.Millisecond) // 25 FPS

	// Start our server enabled by default
	server.Enabled = true

	// Run the renderer
	go server.render()

	// Return the RenderServer
	return server
}
