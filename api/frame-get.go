package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/labstack/echo"
)

// RGBFrameGet returns one frame, of RGB values, via GET
func RGBFrameGet(c echo.Context) error {
	if CurrentFrame == nil {
		c.Logger().Warn("Creating empty RGBFrame for currentFrame")
		CurrentFrame = &frame.RGBFrame{}
	}

	return c.JSON(http.StatusOK, *CurrentFrame)
}
