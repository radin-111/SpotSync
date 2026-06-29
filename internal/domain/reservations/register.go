package reservations

import (
	"SpotSync/internal/config"
	"SpotSync/internal/domain/users"
	"SpotSync/internal/middlewares"
	"SpotSync/internal/utils/auth"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	reservationRepository := NewRepository(db)
	reservationService := NewService(reservationRepository)
	jwtService := auth.NewJWTService(cfg.JWTSecret)
	reservationHandler := NewHandler(reservationService)

	api := e.Group("/api/v1/reservations")

	api.POST("", reservationHandler.CreateReservation, middlewares.ValidateUser(jwtService, []string{users.UserRoleAdmin, users.UserRoleDriver}))

	api.GET("/my-reservations", reservationHandler.GetAllReservationsByUserId, middlewares.ValidateUser(jwtService, []string{users.UserRoleAdmin, users.UserRoleDriver}))

	api.DELETE("/:id", reservationHandler.DeleteReservation, middlewares.ValidateUser(jwtService, []string{users.UserRoleAdmin, users.UserRoleDriver}))
}
