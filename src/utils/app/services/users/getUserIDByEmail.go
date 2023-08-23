package users

import (
	"fmt"
	"gorm.io/gorm"
)

func GetUserIDByEmail(db *gorm.DB, email string) (int, error) {
	var user struct {
		ID int
	}

	result := db.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("user not found")
		}
		return 0, result.Error
	}

	return user.ID, nil
}
