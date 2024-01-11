package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s *server) runHttpServer() error {
	s.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"version":       s.cfg.AppVersion,
			"development":   s.cfg.Server.Development,
			"read_timeout":  s.cfg.Server.ReadTimeout,
			"write_timeout": s.cfg.Server.WriteTimeout,
		})
	})

	s.echo.Server.ReadTimeout = time.Second * s.cfg.Server.ReadTimeout
	s.echo.Server.WriteTimeout = time.Second * s.cfg.Server.WriteTimeout

	return s.echo.Start(s.cfg.Server.Port)
}
