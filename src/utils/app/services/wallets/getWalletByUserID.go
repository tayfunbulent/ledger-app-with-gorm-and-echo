package wallets

import (
	"ledgerApp/src/utils/models"
	"gorm.io/gorm"
)

func GetWalletByUserID(db *gorm.DB, userID int) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
