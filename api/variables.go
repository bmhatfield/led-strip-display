package api

import "github.com/bmhatfield/led-strip-display/frame"

var (
	// CurrentFrame keeps the live frame in memory
	CurrentFrame *frame.RGBFrame

	// RenderQueue is where we publish frames
	RenderQueue chan<- frame.HexGRBFrame
)
