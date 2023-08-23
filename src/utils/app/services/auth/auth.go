package auth

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
	"errors"
)

func AuthenticateUserMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			email := c.Request().Header.Get("email")
			password := c.Request().Header.Get("password")

			if email == "" || password == "" {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"message": "Email and password are required",
				})
			}

			authenticated, err := authenticateUser(db, email, password)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Authentication failed: %v", err))
			}

			if !authenticated {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed: invalid credentials")
			}

			return next(c)
		}
	}
}

func authenticateUser(db *gorm.DB, email, password string) (bool, error) {
	var user models.User

	err := db.Where("email = ? AND password = ?", email, password).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("invalid credentials")
		}
		return false, err
	}

	if user.Role == "admin" {
		return true, nil
	}

	return false, nil
}
