package ledstrip

import "github.com/bmhatfield/led-strip-display/frame"

// Strip represents the OS-Specific Strip Interface
type Strip interface {
	Init(gpio, pixels, brightness int) error
	Render(strip frame.HexGRBFrame) error
}
