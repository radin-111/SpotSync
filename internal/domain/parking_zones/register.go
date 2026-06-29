package parking_zones

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	parkingZRepo := NewRepository(db)
	parkingZService := NewService(parkingZRepo)
	parkingZHandler := NewHandler(parkingZService)
	api := e.Group("/api/v1/zones")
	api.POST("", parkingZHandler.CreateParkingZone)
}
