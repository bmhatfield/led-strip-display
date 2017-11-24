package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	ledstrip "github.com/bmhatfield/led-strip-display/led-strip"
	"github.com/labstack/echo"
)

// RGBFramePost recieves one frame, of RGB values, via POST
func RGBFramePost(c echo.Context) error {
	frame := new(frame.RGBFrame)

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
