package wallets

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"ledgerApp/src/utils/app/services/users"
	"gorm.io/gorm"
)

func GetWalletHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Request().Header.Get("email")
		password := c.Request().Header.Get("password")

		userID, err := users.GetUserIDByCredentials(db, email, password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
		}

		wallet, err := GetWalletByUserID(db, userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve wallet: %v", err))
		}

		return c.JSON(http.StatusOK, wallet)
	}
}
