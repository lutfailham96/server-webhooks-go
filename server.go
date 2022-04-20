package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os/exec"
)

type JsonResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func main() {
	listenPort := flag.Int("port", 7000, "listen port")
	cmdMain := flag.String("cmd", "true", "main command")
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &JsonResponse{
			Status: "OK",
			Data:   "Server running healthy",
		})
	})
	e.GET("/webhooks", func(c echo.Context) error {
		cmd := exec.Command(*cmdMain)
		stdout, err := cmd.Output()
		if err != nil {
			stdout = []byte(err.Error())
		}
		data := &JsonResponse{
			Status: "OK",
			Data:   string(stdout),
		}
		return c.JSON(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *listenPort)))
}
