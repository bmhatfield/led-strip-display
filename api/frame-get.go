package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/labstack/echo"
)

// RGBFrameGet returns one frame, of RGB values, via GET
func RGBFrameGet(c echo.Context) error {
	if currentFrame == nil {
		c.Logger().Warn("Creating empty RGBFrame for currentFrame")
		currentFrame = new(frame.RGBFrame)
	}

	return c.JSON(http.StatusOK, *currentFrame)
}
