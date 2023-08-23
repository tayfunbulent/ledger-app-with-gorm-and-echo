package transactions

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/app/services/users"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

var ErrInsufficientFunds = echo.NewHTTPError(http.StatusBadRequest, "Insufficient funds")

func CreateTransactionHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		senderEmail := c.Request().Header.Get("email")
		senderPassword := c.Request().Header.Get("password")

		senderID, err := users.GetUserIDByCredentials(db, senderEmail, senderPassword)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		var transferData models.TransferData
		if err := c.Bind(&transferData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		if transferData.RecipientEmail == senderEmail {
			return echo.NewHTTPError(http.StatusBadRequest, "You can not transfer funds to yourself")
		}

		if transferData.Amount <= 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Amount must be greater than 0")
		}

		recipientID, err := users.GetUserIDByEmail(db, transferData.RecipientEmail)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Recipient not found")
		}

		if err := createTransaction(db, senderID, recipientID, transferData.Amount); err != nil {
			if err == ErrInsufficientFunds {
				return echo.NewHTTPError(http.StatusBadRequest, "Insufficient funds")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to transfer funds: %v", err))
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Funds transferred successfully",
		})
	}
}

func createTransaction(db *gorm.DB, senderID, recipientID int, amount float64) error {
	var senderWallet, recipientWallet models.Wallet

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", senderID).First(&senderWallet).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", recipientID).First(&recipientWallet).Error; err != nil {
			return err
		}

		if amount > senderWallet.Balance {
			return ErrInsufficientFunds
		}

		if err := tx.Model(&senderWallet).Update("balance", senderWallet.Balance-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&recipientWallet).Update("balance", recipientWallet.Balance+amount).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.Transaction{UserID: uint(senderID), Description: fmt.Sprintf("Transfer to %d", recipientID), Amount: -amount}).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.Transaction{UserID: uint(recipientID), Description: fmt.Sprintf("Transfer from %d", senderID), Amount: amount}).Error; err != nil {
			return err
		}

		return nil
	})
}
