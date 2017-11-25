package ledstrip

import "log"
import "github.com/bmhatfield/led-strip-display/frame"
import "time"

// RenderServer is an async server for rendering frames
type RenderServer struct {
	FrameQueue chan frame.HexGRBFrame
	strip      Strip
	ticker     *time.Ticker
}

func (r RenderServer) render() {
	for _ = range r.ticker.C {
		frame := <-r.FrameQueue

		err := r.strip.Render(frame)

		if err != nil {
			log.Println("Unable to render frame:", err)
		}
	}
}

// NewRenderServer returns the channel you can send to a RenderServer on
func NewRenderServer(strip Strip) *RenderServer {
	// Create the new RenderServer struct
	server := &RenderServer{strip: strip}

	// Create our queue of frames to render
	server.FrameQueue = make(chan frame.HexGRBFrame, 120)

	// Create a timing mechanism to ensure consistent FPS
	server.ticker = time.NewTicker(40 * time.Millisecond) // 25 FPS

	// Run the renderer
	go server.render()

	// Return the
	return server
}
