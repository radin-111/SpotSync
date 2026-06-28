package users

import (
	"SpotSync/internal/domain/users/dto"
	"SpotSync/internal/httpResponse"

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

	response, err := h.service.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) Login(c *echo.Context) error {
	var req dto.LoginRequest
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
	response, err := h.service.Login(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.ErrorResponse{
			Success: false,
			Message: "Internal server error",
			Errors:  "Invalid credentials",
		})
	}
	return c.JSON(http.StatusOK, response)
}
