package repositories

import (
    "minipay/models"

    "gorm.io/gorm"
)

type WalletRepository struct {
    DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
    return &WalletRepository{DB: db}
}

func (r *WalletRepository) FindByUserID(userID uint) (*models.Wallet, error) {
    var wallet models.Wallet
    err := r.DB.Where("user_id = ?", userID).First(&wallet).Error
    return &wallet, err
}

func (r *WalletRepository) Create(wallet *models.Wallet) error {
    return r.DB.Create(wallet).Error
}

func (r *WalletRepository) Save(wallet *models.Wallet) error {
    return r.DB.Save(wallet).Error
}

func (r *WalletRepository) FindByID(id uint) (*models.Wallet, error) {
    var wallet models.Wallet
    err := r.DB.First(&wallet, id).Error
    return &wallet, err
}