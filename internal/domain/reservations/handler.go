package reservations

import (
	"SpotSync/internal/domain/reservations/dto"
	"SpotSync/internal/httpResponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}
func getCurrentUserID(c *echo.Context) (uint, bool) {
	userId, ok := c.Get("user_id").(uint)
	return userId, ok
}
func (h *handler) CreateReservation(c *echo.Context) error {
	var req dto.CreateReservationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "Invalid request payload",
			Errors:  err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "Validation failed",
			Errors:  err.Error(),
		})
	}
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "User ID not found",
			Errors:  "User ID not found",
		})
	}

	reservation, err := h.service.CreateReservation(&req, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Failed to create reservation",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, reservation)
}

func (h *handler) GetMyReservations(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "User ID not found",
			Errors:  "User ID not found",
		})
	}
	reservations, err := h.service.GetMyReservations(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Failed to get reservations",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, reservations)
}

func (h *handler) DeleteReservation(c *echo.Context) error {
	reservationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "Invalid reservation ID",
			Errors:  err.Error(),
		})
	}
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "User ID not found",
			Errors:  "User ID not found",
		})
	}
	err = h.service.DeleteReservation(uint(reservationId), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Failed to delete reservation",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.DeleteReservationResponse{
		Success: true,
		Message: "Reservation deleted successfully",
	})
}
