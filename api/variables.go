package api

import "github.com/bmhatfield/led-strip-display/frame"
import "github.com/bmhatfield/led-strip-display/led-strip"

var (
	// CurrentFrame keeps the live frame in memory
	CurrentFrame *frame.HexGRB

	// Renderer has a FrameQueue where we publish frames
	Renderer *ledstrip.RenderServer
)
