package repositories

import (
    "minipay/models"

    "gorm.io/gorm"
)

type TransactionRepository struct {
    DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
    return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Create(tx *gorm.DB, transaction *models.Transaction) error {
    return tx.Create(transaction).Error
}

func (r *TransactionRepository) FindByWalletID(walletID uint) ([]models.Transaction, error) {
    var transactions []models.Transaction
    err := r.DB.Where(
        "sender_wallet_id = ? OR receiver_wallet_id = ?",
        walletID, walletID,
    ).Order("created_at desc").Find(&transactions).Error
    return transactions, err
}