package parking_zones

import (
	"SpotSync/internal/config"
	"SpotSync/internal/domain/users"
	"SpotSync/internal/middlewares"
	"SpotSync/internal/utils/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	parkingZRepo := NewRepository(db)
	parkingZService := NewService(parkingZRepo)
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	parkingZHandler := NewHandler(parkingZService)
	api := e.Group("/api/v1/zones")
	api.GET("", parkingZHandler.GetAllParkingZones)
	api.GET("/:id", parkingZHandler.GetParkingZoneByID)
	api.POST("", parkingZHandler.CreateParkingZone, middlewares.ValidateUser(jwtService, []string{users.UserRoleAdmin}))
}
