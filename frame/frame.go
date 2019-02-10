package frame

import (
	"fmt"
)

// RenderableFrame represents a frame that can be rendered by a displaystrip,
// which accepts a HexGRB frame
type RenderableFrame interface {
	RenderFrame() (HexGRB, error)
}

// RGBToHex takes an array of 3 color values (R, G, B) and returns a hex uint32
func RGBToHex(colors [3]int) uint32 {
	r := colors[0] * 0x10000
	g := colors[1] * 0x100
	b := colors[2]

	return uint32(r + g + b)
}

// RGBToString takes an array of 3 color values (R, G, B) and returns a string
func RGBToString(colors [3]int) string {
	return fmt.Sprintf("rgb(%d, %d, %d)", colors[0], colors[1], colors[2])
}

// RGBToColors takes a hex value and returns an array of color values
func RGBToColors(pixel uint32) [3]int {
	r := int(pixel / 0x10000)
	g := int((pixel / 0x100) - uint32(r)*0x100)
	b := int(pixel - ((pixel / 0x100) * 0x100))

	return [3]int{r, g, b}
}

// RGBSubtract manipulates RGB channels of a pixel
func RGBSubtract(pixel [3]int, r, g, b int) [3]int {
	if pixel[0] > r {
		pixel[0] -= r
	}

	if pixel[1] > g {
		pixel[1] -= g
	}

	if pixel[2] > b {
		pixel[2] -= b
	}

	return pixel
}
