package users

import (
	"SpotSync/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	usersRepo := NewRepository(db)
	usersService := newService(usersRepo)
	usersHandler := newHandler(usersService)
	api := e.Group("/api/v1/auth")
	api.POST("/register", usersHandler.CreateUser)
}
