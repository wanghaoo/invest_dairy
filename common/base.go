package common

import (
	 "net/http"
	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
		if c.Request().Method == "OPTIONS" {
			c.Response().WriteHeader(http.StatusNoContent)
			return nil
	}
    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    return next(c)
  }
}