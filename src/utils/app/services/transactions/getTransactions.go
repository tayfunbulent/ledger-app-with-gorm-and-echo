package transactions

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/app/services/users"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetTransactionsHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Request().Header.Get("email")
		password := c.Request().Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		transactions, err := getTransactions(db, userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve user transactions: %v", err))
		}

		return c.JSON(http.StatusOK, transactions)
	}
}

func getTransactions(db *gorm.DB, userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	err := db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
