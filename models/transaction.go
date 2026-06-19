package models

import (
    "time"
    "gorm.io/gorm"
)

type Transaction struct {
    ID               uint           `gorm:"primarykey" json:"id"`
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
    DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
    SenderWalletID   uint           `gorm:"not null" json:"sender_wallet_id"`
    ReceiverWalletID uint           `gorm:"not null" json:"receiver_wallet_id"`
    Amount           float64        `gorm:"not null" json:"amount"`
    Status           string         `gorm:"type:varchar(20);default:'success'" json:"status"`
}
