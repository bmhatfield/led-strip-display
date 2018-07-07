package main

import (
	"github.com/bmhatfield/led-strip-display/api"
	"github.com/bmhatfield/led-strip-display/led-strip"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	// Get our OS-defined LED Strip
	strip := ledstrip.GetStrip()

	// Set up API Render Queue
	api.Renderer = ledstrip.NewRenderServer(strip)

	// Spin up new Echo server object
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// /frame/rgb recieves one frame with values in `rgb(r,g,b)` format
	e.POST("/frame/rgb", api.StringRGBFramePost)

	// /frame/hex/rgb returns one frame with values in `#rrggbb`` format
	e.POST("/frame/hex/rgb", api.HexRGBFramePost)

	// /frame/rgb returns one frame with values in `rgb(r,g,b)` format
	e.GET("/frame/rgb", api.StringRGBFrameGet)

	// /frame/hex/rgb returns one frame with values in `#rrggbb`` format
	e.GET("/frame/hex/rgb", api.HexRGBFrameGet)

	// Serve the UI static files
	e.Static("/", "ui")

	// Listen for connections
	e.Logger.Fatal(e.Start(":5000"))
}
