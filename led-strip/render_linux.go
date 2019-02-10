package ledstrip

import (
	"errors"

	"github.com/bmhatfield/led-strip-display/frame"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

// LinuxStrip to satisfy the Strip interface
type LinuxStrip struct {
	device *ws2811.WS2811
}

// Init initializes the ws2811 strip
func (s LinuxStrip) Init(gpio, pixels, brightness int) error {
	options := ws2811.DefaultOptions

	options.Channels[0].GpioPin = gpio
	options.Channels[0].LedCount = pixels
	options.Channels[0].Brightness = brightness

	device, err := ws2811.MakeWS2811(&options)
	if err != nil {
		return err
	}

	if device == nil {
		return errors.New("No error returned, but strip device is nil")
	}

	s.device = device
	return s.device.Init()
}

// Render publishes 150 pixels to an LED strip
func (s LinuxStrip) Render(strip frame.HexGRB) error {
	err := s.device.Wait()
	if err != nil {
		return err
	}

	leds := s.device.Leds(0)
	for index, color := range strip {
		leds[index] = color
	}

	return s.device.Render()
}

// GetStrip returns a LinuxStrip to satisfy the Strip interface
func GetStrip() Strip {
	return LinuxStrip{}
}
