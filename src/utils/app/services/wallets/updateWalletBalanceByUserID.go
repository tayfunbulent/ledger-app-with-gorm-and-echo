package wallets

import (
	"fmt"
	"net/http"
	"strconv"
	"math"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
)

func UpdateWalletBalanceByUserIDHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateData models.UpdateBalanceData
		if err := c.Bind(&updateData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		userID, err := strconv.Atoi(updateData.UserID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
		}

		err = updateWalletBalanceByUserID(db, userID, updateData.Amount)
		if err != nil {
			if err == ErrInsufficientFunds {
				return echo.NewHTTPError(http.StatusBadRequest, "Insufficient funds")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update wallet balance: %v", err))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Wallet balance updated successfully",
		})
	}
}

func updateWalletBalanceByUserID(db *gorm.DB, userID int, amount float64) error {
	var wallet models.Wallet
	if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return err
	}

	if amount < 0 && math.Abs(amount) > wallet.Balance {
		return ErrInsufficientFunds
	}

	return db.Model(&wallet).Update("balance", gorm.Expr("balance + ?", amount)).Error
}
