package api

import (
	"net/http"

	"github.com/bmhatfield/led-strip-display/frame"
	"github.com/labstack/echo"
)

// RGBFramePost recieves one frame, of RGB values, via POST
func RGBFramePost(c echo.Context) error {
	f := &frame.RGBFrame{}

	err := c.Bind(f)

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	queuableFrame, err := f.ToHexGRBFrame()

	if err != nil {
		c.Logger().Warnf("Unable to process request with RGBFrame: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	Renderer.FrameQueue <- queuableFrame
	CurrentFrame = f

	return c.JSON(http.StatusOK, f)
}
