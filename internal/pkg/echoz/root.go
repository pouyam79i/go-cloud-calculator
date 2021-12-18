package echoz

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func BuildServer(port string) {
	PORT := ":" + port
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(PORT))
}
