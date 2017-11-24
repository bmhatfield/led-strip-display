package main

import (
	"github.com/bmhatfield/led-strip-display/api"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	// Spin up new Echo server object
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// /frame/rgb recieves one frame with values in `rgb(r,g,b)` format
	e.POST("/frame/rgb", api.RGBFramePost)

	// /frame/rgb returns one frame with values in `rgb(r,g,b)` format
	e.GET("/frame/rgb", api.RGBFrameGet)

	// Serve the UI static files
	e.Static("/", "ui")

	// Listen for connections
	e.Logger.Fatal(e.Start(":5000"))
}
