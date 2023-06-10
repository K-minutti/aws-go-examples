package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// func main() {
// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/", hello)

// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func handler(c echo.Context) error {
// 	return c.String(http.StatusOK, "Echo!")
// }