package middlewares

import (
	"SpotSync/internal/utils/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func ValidateUser(jwtService auth.JWTService, roles []string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized access"})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
			}
			token := parts[1]

			claims, err := jwtService.VerifyToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}

			isAuthorized := contains(roles, claims.Role)
			if len(roles) > 0 && !isAuthorized {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Forbidden"})
			}
			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)
			c.Set("email", claims.Email)
			c.Set("name", claims.Name)
			return next(c)

		}

	}
}

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
