package users

import (
	"SpotSync/internal/domain/users/dto"
	httpResponse "SpotSync/internal/httpresponse"

	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateUser(c *echo.Context) error {
	var req dto.RegistrationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Details: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response)
}
