package ledstrip

import (
	"fmt"

	ws2811 "github.com/jgarff/rpi_ws281x"
)

// Render publishes a StringColors to an LED strip
func Render(strip StringColors) {
	fmt.Println("Publishing LED Display Strip...")

	err := ws2811.Init(12, 150, strip.Brightness)

	if err != nil {
		fmt.Println(err)
	}

	for index, color := range strip.LEDs {
		hex := (color[0] * 0x10000) + (color[1] * 0x100) + color[2]
		fmt.Printf("%v -> %08X, ", index+1, hex)
		ws2811.SetLed(index, uint32(hex))
	}

	err = ws2811.Render()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Rendering Complete!")
	fmt.Println()
}
