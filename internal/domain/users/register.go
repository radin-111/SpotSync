package users

import (
	"SpotSync/internal/config"
	"SpotSync/internal/utils/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	usersRepo := NewRepository(db)
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	usersService := newService(usersRepo, jwtService)
	usersHandler := newHandler(usersService)
	api := e.Group("/api/v1/auth")
	api.POST("/register", usersHandler.CreateUser)
	api.POST("/login", usersHandler.Login)
}
