package server

import (
	"SpotSync/internal/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, cfg *config.Config) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
