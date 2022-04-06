package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()  // echo を利用する
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello test")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
