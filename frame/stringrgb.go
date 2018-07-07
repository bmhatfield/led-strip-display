package frame

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

// StringRGB represents a JS `rgb(r,g,b)` frame representation
type StringRGB [150]string

// RenderFrame returns a renderable frame version of StringRGB
func (f *StringRGB) RenderFrame() (HexGRB, error) {
	bufferPixel := [3]int{}
	frame := HexGRB{}

	for index, pixel := range f {
		left := strings.Index(pixel, "(")
		right := strings.Index(pixel, ")")

		if left < 0 || right < 0 {
			continue
		}

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

		frame[index] = RGBToHex(bufferPixel)
	}

	return frame, nil
}

// StringRGBFromDisk loads a StringColors from an on-disk JSON encoding
func StringRGBFromDisk(path string) (*StringRGB, error) {
	frame := &StringRGB{}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, frame)

	return frame, err
}
