package users

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
	"strconv"
)

func UpdateUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateData models.UpdateData
		if err := c.Bind(&updateData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to decode update data: %v", err))
		}

		if updateData.User.Username == "" || updateData.User.Email == "" || updateData.User.Password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing required fields")
		}

		userID, err := strconv.ParseUint(updateData.UserID, 10, 32)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id format")
		}

		userExists, err := doesUserExist(db, uint(userID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error checking user existence: %v", err))
		}

		if !userExists {
			return echo.NewHTTPError(http.StatusBadRequest, "User with provided user_id does not exist")
		}

		existingUserID, err := getUserIDByEmail(db, updateData.User.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve user: %v", err))
		}

		if existingUserID != 0 && existingUserID != uint(userID) {
			return echo.NewHTTPError(http.StatusBadRequest, "Email is already in use")
		}

		updateData.User.ID = uint(userID)
		if err := updateUser(db, &updateData.User); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update user: %v", err))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "User updated successfully",
		})
	}
}

func doesUserExist(db *gorm.DB, userID uint) (bool, error) {
    var count int64
    if err := db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

func getUserIDByEmail(db *gorm.DB, email string) (uint, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return user.ID, nil
}

func updateUser(db *gorm.DB, user *models.User) error {
	return db.Model(user).Where("id = ?", user.ID).Updates(models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}
