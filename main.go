package main

import (
	"log"
	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("gello world!")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Static("/static", "../_static")

	e.Start(":8080")
}