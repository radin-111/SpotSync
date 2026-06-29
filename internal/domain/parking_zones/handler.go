package parking_zones

import (
	"SpotSync/internal/domain/parking_zones/dto"
	"SpotSync/internal/httpResponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateParkingZone(c *echo.Context) error {

	var req dto.CreateParkingZoneRequest
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

	response, err := h.service.CreateParkingZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)

}
