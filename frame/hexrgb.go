package frame

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// HexRGB represents the CSS format of pixels
type HexRGB [150]uint32

// RenderFrame returns a renderable frame version of HexRGB
func (f *HexRGB) RenderFrame() (HexGRB, error) {
	frame := HexGRB{}

	for index, pixel := range f {
		colors := RGBToColors(pixel)

		grb := [3]int{colors[1], colors[0], colors[2]}

		frame[index] = RGBToHex(grb)
	}

	return frame, nil
}

// UnmarshalJSON converts the string representation into []uint32
func (f HexRGB) UnmarshalJSON(b []byte) error {
	var strings []string
	if err := json.Unmarshal(b, &strings); err != nil {
		return err
	}

	for index, pixel := range strings {
		hexCode := pixel[1:]
		val, err := strconv.ParseInt(hexCode, 16, 64)

		if err != nil {
			return err
		}

		f[index] = uint32(val)
	}

	return nil
}

// MarshalJSON converts the []uint32 into its string representation
func (f HexRGB) MarshalJSON() ([]byte, error) {
	asString := []string{}

	for _, pixel := range f {
		hexString := fmt.Sprintf("#%06x", pixel)
		asString = append(asString, hexString)
	}

	return json.Marshal(asString)
}
