package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	startServer()
}

func startServer() {
	e := echo.New()
	err := e.Start(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server started")
}
