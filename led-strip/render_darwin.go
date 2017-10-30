package ledstrip

import "fmt"

// Render publishes a StringColors to an LED strip
func Render(strip StringColors) {
	fmt.Println("Publishing LED Display Strip...")

	for index, color := range strip.LEDs {
		hex := (color[0] * 0x10000) + (color[1] * 0x100) + color[2]
		fmt.Printf("%v -> %08X, ", index+1, hex)
	}

	fmt.Println("Rendering Complete!")
	fmt.Println()
}
