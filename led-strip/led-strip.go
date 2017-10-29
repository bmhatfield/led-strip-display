package ledstrip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	ws2811 "github.com/jgarff/rpi_ws281x"
)

// StringColors is a serialized version of an LED strip
type StringColors struct {
	Brightness int     `json:"Brightness"`
	LEDs       [][]int `json:"LEDs"`
}

// LoadStringFromPath loads a StringColors from an on-disk JSON encoding
func LoadStringFromPath(path string) StringColors {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var sc StringColors

	json.Unmarshal(raw, &sc)

	return sc
}

// Render publishes a StringColors to an LED strip
func Render(strip StringColors) {
	fmt.Println("Publishing LED Display Strip...")

	err := ws2811.Init(12, 150, strip.Brightness)

	if err != nil {
		fmt.Println(err)
	}

	for index, color := range strip.LEDs {
		hex := (color[0] * 0x10000) + (color[1] * 0x100) + color[2]
		fmt.Printf("Setting LED %v to %08X\n", index+1, hex)
		ws2811.SetLed(index, uint32(hex))
	}

	err = ws2811.Render()

	if err != nil {
		fmt.Println(err)
	}
}
