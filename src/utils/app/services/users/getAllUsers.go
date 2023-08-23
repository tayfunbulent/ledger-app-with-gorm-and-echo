package users

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetAllUsersHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []*models.User

		if err := db.Select("id, username, email, password, role").Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve users: %v", err))
		}

		return c.JSON(http.StatusOK, users)
	}
}
