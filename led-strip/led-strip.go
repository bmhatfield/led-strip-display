package ledstrip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// StringColors is a serialized version of an LED strip
type StringColors struct {
	Brightness int      `json:"Brightness"`
	LEDs       [][3]int `json:"LEDs"`
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
