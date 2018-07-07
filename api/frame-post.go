package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/labstack/echo"
)

// StringRGBFramePost recieves one frame, of RGB values, via POST
func StringRGBFramePost(c echo.Context) error {
	f := &frame.StringRGB{}

	err := c.Bind(f)

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	queuableFrame, err := f.RenderFrame()

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	Renderer.FrameQueue <- queuableFrame
	CurrentFrame = &queuableFrame

	return c.JSON(http.StatusOK, f)
}

// HexRGBFramePost recieves one frame, of HexRGB values, via POST
func HexRGBFramePost(c echo.Context) error {
	f := &frame.HexRGB{}

	err := c.Bind(f)

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	queuableFrame, err := f.RenderFrame()

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	Renderer.FrameQueue <- queuableFrame
	CurrentFrame = &queuableFrame

	return c.JSON(http.StatusOK, f)
}
