package users

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func UpdateUserRoleHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		updateData := &models.UpdateRoleData{}
		if err := c.Bind(&updateData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to decode update data: %v", err))
		}

		userIDUint, err := strconv.ParseUint(updateData.UserID, 10, 64)
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

		if updateData.Role != "admin" && updateData.Role != "user" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid role provided. Acceptable values are 'admin' or 'user'.")
		}

		if err := updateUserRole(db, uint(userIDUint), updateData.Role); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update user role: %v", err))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "User role updated successfully",
		})
	}
}

func updateUserRole(db *gorm.DB, userID uint, role string) error {
	return db.Model(&models.User{}).Where("id = ?", userID).Update("role", role).Error
}
