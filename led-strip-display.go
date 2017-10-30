package main

import (
	"net/http"
	"strconv"
	"strings"

	ledstrip "github.com/bmhatfield/led-strip-display/led-strip"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// RGBFrame represents a single frame in RGB format
type RGBFrame struct {
	Brightness int      `json:"brightness"`
	Pixels     []string `json:"pixels"`
}

// StringColors returns a GRB slice of pixels
func (f RGBFrame) StringColors() (*ledstrip.StringColors, error) {
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

		strip.LEDs = append(strip.LEDs, []int{green, red, blue})
	}

	strip.Brightness = f.Brightness

	return strip, nil
}

// RGBFramePost recieves one frame, of RGB values, via POST
func RGBFramePost(c echo.Context) error {
	frame := new(RGBFrame)

	err := c.Bind(frame)

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	colors, err := frame.StringColors()

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	ledstrip.Render(*colors)

	return c.JSON(http.StatusOK, frame)
}

func main() {
	// Spin up new Echo server object
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// /frame/rgb recieves one frame with values in `rgb(r,g,b)` format
	e.POST("/frame/rgb", RGBFramePost)

	// Serve the UI static files
	e.Static("/", "ui")

	// Listen for connections
	e.Logger.Fatal(e.Start(":5000"))
}
