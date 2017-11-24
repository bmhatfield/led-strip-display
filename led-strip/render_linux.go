package ledstrip

import (
	"github.com/bmhatfield/led-strip-display/frame"
	ws2811 "github.com/jgarff/rpi_ws281x" // Will not build on OSX
)

// LinuxStrip to satisfy the Strip interface
type LinuxStrip struct{}

// Init initializes the ws2811 strip
func (s LinuxStrip) Init(gpio, pixels, brightness int) error {
	return ws2811.Init(gpio, pixels, brightness)
}

// Render publishes 150 pixels to an LED strip
func (s LinuxStrip) Render(strip frame.HexGRBFrame) error {
	for index, color := range strip {
		ws2811.SetLed(index, color)
	}

	return ws2811.Render()
}

// GetStrip returns a LinuxStrip to satisfy the Strip interface
func GetStrip() Strip {
	return LinuxStrip{}
}
