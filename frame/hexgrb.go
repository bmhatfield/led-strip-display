package frame

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// HexGRB represents the strip-side format of pixels
type HexGRB [150]uint32

// RenderFrame returns a renderable frame version of HexGRB
func (f *HexGRB) RenderFrame() (HexGRB, error) {
	return *f, nil
}

// HexRGB returns a HexRGB version of the frame
func (f *HexGRB) HexRGB() (HexRGB, error) {
	frame := HexRGB{}

	for index, pixel := range f {
		colors := RGBToColors(pixel)

		rgb := [3]int{colors[1], colors[0], colors[2]}

		frame[index] = RGBToHex(rgb)
	}

	return frame, nil
}

// UnmarshalJSON converts the string representation into []uint32
func (f HexGRB) UnmarshalJSON(b []byte) error {
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
func (f HexGRB) MarshalJSON() ([]byte, error) {
	asString := []string{}

	for _, pixel := range f {
		hexString := fmt.Sprintf("#%06x", pixel)
		asString = append(asString, hexString)
	}

	return json.Marshal(asString)
}
