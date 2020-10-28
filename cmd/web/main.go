package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type healthStatus struct {
	Status string `json:"status" xml:"status"`
}

func health(c echo.Context) error {
	hs := &healthStatus{
		Status: "Up"}
	return c.JSON(http.StatusOK, hs)
}
func info(c echo.Context) error {
	i := map[string]string{
		"version": "1.0.0",
		"desc":    "Sample Go App"}
	return c.JSON(http.StatusOK, i)
}

func main() {
	var host string
	var port string
	flag.StringVar(&host, "host", getEnv("LISTEN_ADDRESS", "0.0.0.0"), "listen host")
	flag.StringVar(&port, "port", getEnv("SERVICE_PORT", "8080"), "listen port")
	flag.Parse()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", health)
	e.GET("/info", info)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
