package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username string `gorm:"size:255" json:"username"`
	Email    string `gorm:"uniqueIndex;size:255" json:"email"`
	Password string `gorm:"size:255" json:"password"`
	Role     string `gorm:"size:255;default:'user'" json:"role"`

	Wallet Wallet `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Transaction struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"wallet_id"`
	UserID      uint      `gorm:"index" json:"user_id"`
	Description string    `gorm:"size:255" json:"description"`
	Amount      float64   `json:"amount"`
	Balance     float64   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
}

type Wallet struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"wallet_id"`
	UserID    uint      `gorm:"index;unique" json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type UserWithWallet struct {
	UserID    uint      `gorm:"column:id" json:"user_id"`
	Username  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Role      string    `gorm:"column:role" json:"role"`
	WalletID  uint      `gorm:"column:wallet_id" json:"wallet_id"`
	Balance   float64   `gorm:"column:balance" json:"balance"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type UpdateRoleData struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type UpdateData struct {
	UserID string      `json:"user_id"`
	User   User `json:"user"`
}

type GetBalanceData struct {
	UserID string `json:"user_id"`
	Time   string `json:"time"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type DeleteUser struct {
	UserID string `json:"user_id"`
}

type WalletWithoutCreatedAt struct {
	ID      int     `gorm:"primaryKey"`
	UserID  int     `gorm:"unique"`
	Balance float64 `gorm:"column:balance"`
}

type UpdateAmount struct {
	Amount float64 `json:"amount"`
}

type UpdateBalanceData struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type TransferData struct {
	RecipientEmail string  `json:"recipient_email"`
	Amount         float64 `json:"amount"`
}
