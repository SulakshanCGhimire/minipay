package handlers

import (
    "minipay/config"
    "minipay/models"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func Transfer(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var input struct {
        ReceiverWalletID uint    `json:"receiver_wallet_id" binding:"required"`
        Amount           float64 `json:"amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Amount <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be greater than zero"})
        return
    }

    // Get sender wallet
    var senderWallet models.Wallet
    if err := config.DB.Where("user_id = ?", userID).First(&senderWallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Sender wallet not found"})
        return
    }

    // Check balance
    if senderWallet.Balance < input.Amount {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
        return
    }

    // Get receiver wallet
    var receiverWallet models.Wallet
    if err := config.DB.First(&receiverWallet, input.ReceiverWalletID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Receiver wallet not found"})
        return
    }

    // Prevent self-transfer
    if senderWallet.ID == receiverWallet.ID {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot transfer to your own wallet"})
        return
    }

    // DB transaction — atomic
    err := config.DB.Transaction(func(tx *gorm.DB) error {
        senderWallet.Balance -= input.Amount
        if err := tx.Save(&senderWallet).Error; err != nil {
            return err
        }

        receiverWallet.Balance += input.Amount
        if err := tx.Save(&receiverWallet).Error; err != nil {
            return err
        }

        transaction := models.Transaction{
            SenderWalletID:   senderWallet.ID,
            ReceiverWalletID: receiverWallet.ID,
            Amount:           input.Amount,
        }
        return tx.Create(&transaction).Error
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Transfer failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Transfer successful",
        "balance": senderWallet.Balance,
    })
}

func GetTransactions(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var wallet models.Wallet
    if err := config.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    var transactions []models.Transaction
    config.DB.Where(
        "sender_wallet_id = ? OR receiver_wallet_id = ?",
        wallet.ID, wallet.ID,
    ).Order("created_at desc").Find(&transactions)

    c.JSON(http.StatusOK, gin.H{
        "transactions": transactions,
    })
}