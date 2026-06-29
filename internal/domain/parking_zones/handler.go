package parking_zones

import (
	"SpotSync/internal/domain/parking_zones/dto"
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

func (h *handler) GetAllParkingZones(c *echo.Context) error {
	response, err := h.service.GetAllParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetParkingZoneByID(c *echo.Context) error {
	zoneId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Success: false,
			Message: "Invalid parking zone ID",
			Errors:  err.Error(),
		})
	}
	zone, err := h.service.GetParkingZoneByID(uint(zoneId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Failed to get parking zone",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, zone)
}
