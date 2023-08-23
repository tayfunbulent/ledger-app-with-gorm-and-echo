package wallets

import (
	"gorm.io/gorm"
	"fmt"
)

func GetWalletBalance(db *gorm.DB, userID int) (float64, error) {
	var wallet struct {
		Balance float64
	}

	if err := db.Select("balance").Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("wallet not found")
		}
		return 0, err
	}

	return wallet.Balance, nil
}
