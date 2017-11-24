package frame

import (
	"strconv"
	"strings"

	ledstrip "github.com/bmhatfield/led-strip-display/led-strip"
)

// RGBFrame represents a single frame in RGB format
type RGBFrame struct {
	Brightness int      `json:"brightness"`
	Pixels     []string `json:"pixels"`
}

// StringColors returns a GRB slice of pixels
func (f *RGBFrame) StringColors() (*ledstrip.StringColors, error) {
	strip := new(ledstrip.StringColors)

	for _, pixel := range f.Pixels {
		left := strings.Index(pixel, "(")
		right := strings.Index(pixel, ")")
		rgb := strings.Split(pixel[left+1:right], ", ")

		red, err := strconv.Atoi(rgb[0])
		if err != nil {
			return nil, err
		}

		green, err := strconv.Atoi(rgb[1])
		if err != nil {
			return nil, err
		}

		blue, err := strconv.Atoi(rgb[2])
		if err != nil {
			return nil, err
		}

		strip.LEDs = append(strip.LEDs, [3]int{green, red, blue})
	}

	strip.Brightness = f.Brightness

	return strip, nil
}
