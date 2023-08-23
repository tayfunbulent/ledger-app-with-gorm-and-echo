package users

import (
	"net/http"
	"strings"
	"unicode"
	"ledgerApp/src/utils/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"fmt"
)

func CreateUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Failed to decode user data",
			})
		}

		if user.Username == "" || user.Email == "" || user.Password == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Missing required fields",
			})
		}

		if !isEmailValid(user.Email) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid email format",
			})
		}

		if !isPasswordValid(user.Password) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid password format",
			})
		}

		if !isEmailTaken(db, user.Email) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Email already exists",
			})
		}

		if err := createUserWithWallet(db, &user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to create user",
			})
		}

		return c.JSON(http.StatusOK, user)
	}
}

func createUserWithWallet(db *gorm.DB, user *models.User) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		wallet := models.Wallet{UserID: user.ID}
		if err := tx.Create(&wallet).Error; err != nil {
			return fmt.Errorf("failed to create wallet for user ID %d: %w", user.ID, err)
		}

		return nil
	})
}

func isEmailValid(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasNumber := false
	hasLetter := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasNumber = true
		} else if unicode.IsLetter(char) {
			hasLetter = true
		}
	}
	if !hasNumber || !hasLetter {
		return false
	}

	return true
}

func isEmailTaken(db *gorm.DB, email string) bool {
	var count int64
	db.Model(&models.User{}).Where("email = ?", email).Count(&count)

	return count == 0
}
