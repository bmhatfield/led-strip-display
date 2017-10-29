package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "ui")

	e.Logger.Fatal(e.Start(":5000"))
}
