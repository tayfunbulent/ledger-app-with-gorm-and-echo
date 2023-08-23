package users

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetAllUsersWithWalletHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var usersWithWallets []*models.UserWithWallet

		if err := db.Table("users").
			Select("users.id, users.username, users.email, users.role, wallets.id as wallet_id, wallets.balance, wallets.created_at").
			Joins("INNER JOIN wallets ON users.id = wallets.user_id").
			Scan(&usersWithWallets).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve users: %v", err))
		}

		return c.JSON(http.StatusOK, usersWithWallets)
	}
}
