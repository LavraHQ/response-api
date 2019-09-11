package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAuth returns the authentication mode that is enabled for
// Response.
func GetAuth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"auth": echo.Map{
			"enabled": true,
			"mode":    "creds",
		},
	})
}
