package services

import (
    "errors"
    "minipay/models"
    "minipay/repositories"
)

type WalletService struct {
    WalletRepo *repositories.WalletRepository
}

func NewWalletService(walletRepo *repositories.WalletRepository) *WalletService {
    return &WalletService{WalletRepo: walletRepo}
}

func (s *WalletService) CreateWallet(userID uint) (*models.Wallet, error) {
    existing, err := s.WalletRepo.FindByUserID(userID)
    if err == nil && existing.ID != 0 {
        return nil, errors.New("wallet already exists")
    }

    wallet := &models.Wallet{UserID: userID, Balance: 0}
    err = s.WalletRepo.Create(wallet)
    return wallet, err
}

func (s *WalletService) GetWallet(userID uint) (*models.Wallet, error) {
    return s.WalletRepo.FindByUserID(userID)
}

func (s *WalletService) Deposit(userID uint, amount float64) (*models.Wallet, error) {
    if amount <= 0 {
        return nil, errors.New("amount must be greater than zero")
    }

    wallet, err := s.WalletRepo.FindByUserID(userID)
    if err != nil {
        return nil, errors.New("wallet not found")
    }

    wallet.Balance += amount
    err = s.WalletRepo.Save(wallet)
    return wallet, err
}