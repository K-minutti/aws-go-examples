package main

import (
	"os"
	"context"
	"net/http"
	
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

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
	e.GET("/nodes/main/:id", getNodes)

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

func getNodes(c echo.Context) error {
	id := c.Param("id")
	path := "nodes/parents/" +  id + "/main.json"
	
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		fmt.Println("S3_BUCKET environment variable not set")
	}
	aws_region := os.Getenv("AWS_REGION")
	if bucket == "" {
		fmt.Println("AWS_REGION environment variable not set")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(aws_region))

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key: &path,
	}

	output, err := client.GetObject(context.TODO(), input)
	if err != nil {
		fmt.Println("Error reading s3 object:", err)
		os.Exit(1)
	}

	data := make([]byte, *output.ContentLength)
	_, err = output.Body.Read(data)
	if err != nil {
		fmt.Println("Error reading s3 object data:", err)
		os.Exit(1)
	}

	content := string(data)

	fmt.Println("S3 object content:", content)
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