package users

import (
	"gorm.io/gorm"
)

func GetUserIDByCredentials(db *gorm.DB, email, password string) (int, error) {
	var user struct {
		ID int
	}

	err := db.Table("users").Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
