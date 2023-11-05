package main

import (
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping/:target", func(c echo.Context) error {
		target := c.Param("target")

		result, err := ping(target)

		if err != nil {
			return err
		}

		return c.String(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func ping(target string) (string, error) {
	cmd := exec.Command("sh", "-c", "ping -c2 "+target)

	data, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(data), nil
}
