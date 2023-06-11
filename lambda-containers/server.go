package main

import (
	"os"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	e.GET("/hello", hello)
	e.GET("/users/:id", usersId)
	e.GET("/users/new", usersNew)
	e.GET("/users/1/files/*", usersFile)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	os.WriteFile("routes.json", data, 0644)

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if username == "joe" && password == "secret" {
			return true, nil
		}
	return false, nil
	}))

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Echo!")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func usersId(c echo.Context) error {
	return c.String(http.StatusOK, "/users/:id")
}

func usersNew(c echo.Context) error {
	return c.String(http.StatusOK, "/users/new")
}

func usersFile(c echo.Context) error {
	return c.String(http.StatusOK, "/users/1/files/*")
}