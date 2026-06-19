package services

import (
    "errors"
    "minipay/config"
    "minipay/models"
    "minipay/repositories"

    "gorm.io/gorm"
)

type TransactionService struct {
    WalletRepo      *repositories.WalletRepository
    TransactionRepo *repositories.TransactionRepository
}

func NewTransactionService(
    walletRepo *repositories.WalletRepository,
    transactionRepo *repositories.TransactionRepository,
) *TransactionService {
    return &TransactionService{
        WalletRepo:      walletRepo,
        TransactionRepo: transactionRepo,
    }
}

func (s *TransactionService) Transfer(senderUserID uint, receiverWalletID uint, amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be greater than zero")
    }

    senderWallet, err := s.WalletRepo.FindByUserID(senderUserID)
    if err != nil {
        return errors.New("sender wallet not found")
    }

    if senderWallet.Balance < amount {
        return errors.New("insufficient balance")
    }

    receiverWallet, err := s.WalletRepo.FindByID(receiverWalletID)
    if err != nil {
        return errors.New("receiver wallet not found")
    }

    if senderWallet.ID == receiverWallet.ID {
        return errors.New("cannot transfer to your own wallet")
    }

    return config.DB.Transaction(func(tx *gorm.DB) error {
        senderWallet.Balance -= amount
        if err := tx.Save(&senderWallet).Error; err != nil {
            return err
        }

        receiverWallet.Balance += amount
        if err := tx.Save(&receiverWallet).Error; err != nil {
            return err
        }

        transaction := &models.Transaction{
            SenderWalletID:   senderWallet.ID,
            ReceiverWalletID: receiverWallet.ID,
            Amount:           amount,
        }
        return s.TransactionRepo.Create(tx, transaction)
    })
}

func (s *TransactionService) GetTransactions(userID uint) ([]models.Transaction, error) {
    wallet, err := s.WalletRepo.FindByUserID(userID)
    if err != nil {
        return nil, errors.New("wallet not found")
    }
    return s.TransactionRepo.FindByWalletID(wallet.ID)
}