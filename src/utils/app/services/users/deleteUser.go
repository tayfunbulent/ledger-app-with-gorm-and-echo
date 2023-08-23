package users

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
)

func DeleteUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var deleteData models.DeleteUser
		if err := c.Bind(&deleteData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to decode delete data: %v", err))
		}

		userIDUint, err := strconv.ParseUint(deleteData.UserID, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id format")
		}

		userExists, err := doesUserExist(db, uint(userIDUint))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error checking user existence: %v", err))
		}

		if !userExists {
			return echo.NewHTTPError(http.StatusBadRequest, "User with provided user_id does not exist")
		}

		if err := deleteUser(db, uint(userIDUint)); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to delete user: %v", err))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "User deleted successfully",
		})
	}
}

func deleteUser(db *gorm.DB, userID uint) error {
	result := db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
