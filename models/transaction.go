package models

import "gorm.io/gorm"

type Transaction struct {
    gorm.Model
    SenderWalletID   uint    `gorm:"not null"`
    ReceiverWalletID uint    `gorm:"not null"`
    Amount           float64 `gorm:"not null"`
    Status           string  `gorm:"type:varchar(20);default:'success'"`
}