package wallets

import (
	"fmt"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
)

func GetWalletBalanceAtTimeHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestData models.GetBalanceData
		if err := c.Bind(&requestData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
		}

		t, err := time.Parse(time.RFC3339, requestData.Time)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid time format")
		}

		balance, err := getBalanceAtTime(db, requestData.UserID, t)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve balance: %v", err))
		}

		responseData := models.BalanceResponse{
			Balance: balance,
		}

		return c.JSON(http.StatusOK, responseData)
	}
}

func getBalanceAtTime(db *gorm.DB, userID string, t time.Time) (float64, error) {
	var transaction struct {
		Balance float64
	}

	if err := db.Table("transactions").
		Select("balance").
		Where("user_id = ? AND created_at <= ?", userID, t).
		Order("created_at DESC").
		Limit(1).
		First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}

	return transaction.Balance, nil
}
