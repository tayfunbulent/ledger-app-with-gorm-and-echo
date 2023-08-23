package wallets

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetWalletByID(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		walletID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid wallet ID")
		}

		wallet, err := getWalletByID(db, walletID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "Wallet not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve wallet: "+err.Error())
		}

		return c.JSON(http.StatusOK, wallet)
	}
}

func getWalletByID(db *gorm.DB, walletID int) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := db.Where("id = ?", walletID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
