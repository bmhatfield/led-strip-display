package frame

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

// HexGRBFrame represents the strip-side format of pixels
type HexGRBFrame [150]uint32

// ConvertToHex takes an array of 3 color values (G, R, B) and returns a hex uint32
func ConvertToHex(colors [3]int) uint32 {
	return uint32((colors[0] * 0x10000) + (colors[1] * 0x100) + colors[2])
}

// RGBFrame represents a single frame in RGB format. Intended to be used in API
// communication between frontend/backend services. Inefficient.
type RGBFrame struct {
	Brightness int      `json:"brightness"`
	Pixels     []string `json:"pixels"`
}

// ToHexGRBFrame returns a HexGRBFrame of pixels
func (f *RGBFrame) ToHexGRBFrame() (HexGRBFrame, error) {
	bufferPixel := [3]int{}
	frame := HexGRBFrame{}

	for index, pixel := range f.Pixels {
		left := strings.Index(pixel, "(")
		right := strings.Index(pixel, ")")
		rgb := strings.Split(pixel[left+1:right], ", ")

		red, err := strconv.Atoi(rgb[0])
		if err != nil {
			return frame, err
		}

		green, err := strconv.Atoi(rgb[1])
		if err != nil {
			return frame, err
		}

		blue, err := strconv.Atoi(rgb[2])
		if err != nil {
			return frame, err
		}

		bufferPixel[0] = green
		bufferPixel[1] = red
		bufferPixel[2] = blue

		frame[index] = ConvertToHex(bufferPixel)
	}

	return frame, nil
}

// RGBFrameFromDisk loads a StringColors from an on-disk JSON encoding
func RGBFrameFromDisk(path string) (*RGBFrame, error) {
	frame := &RGBFrame{}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, frame)

	return frame, err
}
