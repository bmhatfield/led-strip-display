package ledstrip

import (
	"fmt"

	"github.com/bmhatfield/led-strip-display/frame"
)

// DarwinStrip to satisfy the Strip interface
type DarwinStrip struct{}

// Init does nothing on OSX.
func (s DarwinStrip) Init(gpio, pixels, brightness int) error {
	return nil
}

// Render logs pixel values on OSX
func (s DarwinStrip) Render(strip frame.HexGRBFrame) error {
	for index, color := range strip {
		fmt.Printf("%v -> %08X, ", index+1, color)
	}

	return nil
}

// GetStrip returns a DarwinStrip to satisfy the Strip interface
func GetStrip() Strip {
	return DarwinStrip{}
}
