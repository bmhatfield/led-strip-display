package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/labstack/echo"
)

// StringRGBFrameGet returns one frame, of RGB values, via GET
func StringRGBFrameGet(c echo.Context) error {
	if CurrentFrame == nil {
		c.Logger().Warn("Creating empty HexGRB for currentFrame")
		CurrentFrame = &frame.HexGRB{}
	}

	frame, err := CurrentFrame.StringRGB()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, frame)
}

// HexRGBFrameGet returns one frame, of Hex values, via GET
func HexRGBFrameGet(c echo.Context) error {
	if CurrentFrame == nil {
		c.Logger().Warn("Creating empty HexGRB for currentFrame")
		CurrentFrame = &frame.HexGRB{}
	}

	frame, err := CurrentFrame.HexRGB()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, frame)
}
