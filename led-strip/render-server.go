package ledstrip

import "log"
import "github.com/bmhatfield/led-strip-display/frame"

// RenderServer is an async server for rendering frames
type RenderServer struct {
	strip  Strip
	frames chan frame.HexGRBFrame
}

func (r RenderServer) render() {
	for frame := range r.frames {
		err := r.strip.Render(frame)

		if err != nil {
			log.Println("Unable to render frame:", err)
		}
	}
}

// NewRenderServer returns the channel you can send to a RenderServer on
func NewRenderServer(strip Strip) chan<- frame.HexGRBFrame {
	server := RenderServer{strip: strip}
	server.frames = make(chan frame.HexGRBFrame, 120)

	go server.render()

	return server.frames
}
