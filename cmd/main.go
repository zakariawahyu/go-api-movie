package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-movie/config"
	"github.com/zakariawahyu/go-api-movie/internal/server"
)

func main() {
	e := echo.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		e.Logger.Fatalf("err config : %v", err)
	}

	s := server.NewHttpServer(cfg)
	e.Logger.Fatal(s.Run())
}
