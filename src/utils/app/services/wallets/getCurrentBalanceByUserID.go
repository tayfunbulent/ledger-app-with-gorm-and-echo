package wallets

import (
	"gorm.io/gorm"
	"ledgerApp/src/utils/models"
)

func GetCurrentBalanceByUserID(db *gorm.DB, userID int) (float64, error) {
	var wallet models.Wallet

	if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return 0, err
	}

	return wallet.Balance, nil
}
