package wallets

import (
	"net/http"
	"math"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/app/services/users"
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
)

var ErrInsufficientFunds = echo.NewHTTPError(http.StatusBadRequest, "Insufficient funds")

func UpdateWalletBalanceHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Request().Header.Get("email")
		password := c.Request().Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		var updateData models.UpdateAmount
		if err := c.Bind(&updateData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if err := updateWalletBalance(db, uint(userID), updateData.Amount); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update wallet balance").SetInternal(err)
		}

		if err := createTransaction(db, uint(userID), updateData.Amount); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create transaction").SetInternal(err)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Wallet balance updated successfully",
		})
	}
}

func updateWalletBalance(db *gorm.DB, userID uint, amount float64) error {
	var wallet models.Wallet
	if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return err
	}

	if amount < 0 && math.Abs(amount) > wallet.Balance {
		return ErrInsufficientFunds
	}

	return db.Model(&wallet).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func createTransaction(db *gorm.DB, userID uint, amount float64) error {
	transaction := models.Transaction{
		UserID:     userID,
		Description: "Wallet update",
		Amount:      amount,
	}

	return db.Create(&transaction).Error
}
