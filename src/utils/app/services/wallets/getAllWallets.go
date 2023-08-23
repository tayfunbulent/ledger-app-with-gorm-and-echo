package wallets

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetAllWallets(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var wallets []*models.Wallet

		if err := db.Find(&wallets).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve wallets: "+err.Error())
		}

		return c.JSON(http.StatusOK, wallets)
	}
}
