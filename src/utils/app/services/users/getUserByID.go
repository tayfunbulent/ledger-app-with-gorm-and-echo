package users

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"strconv"
	"gorm.io/gorm"
)

func GetUserByIDHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
		}

		user, err := getUserByID(db, userID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "User not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve user: %v", err))
		}

		return c.JSON(http.StatusOK, user)
	}
}

func getUserByID(db *gorm.DB, userID int) (*models.User, error) {
	var user models.User

	err := db.First(&user, userID).Error
	
	if err != nil {
		return nil, err
	}

	return &user, nil
}
