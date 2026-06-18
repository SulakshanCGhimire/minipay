package handlers

import (
    "minipay/config"
    "minipay/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var existing models.Wallet
    if err := config.DB.Where("user_id = ?", userID).First(&existing).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet already exists"})
        return
    }

    wallet := models.Wallet{
        UserID:  userID.(uint),
        Balance: 0,
    }

    if err := config.DB.Create(&wallet).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Wallet created successfully",
        "balance": wallet.Balance,
    })
}

func Deposit(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var input struct {
        Amount float64 `json:"amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if input.Amount <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be greater than zero"})
        return
    }

    var wallet models.Wallet
    if err := config.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    wallet.Balance += input.Amount
    config.DB.Save(&wallet)

    c.JSON(http.StatusOK, gin.H{
        "message": "Deposit successful",
        "balance": wallet.Balance,
    })
}

func GetWallet(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var wallet models.Wallet
    if err := config.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "balance":    wallet.Balance,
        "created_at": wallet.CreatedAt,
    })
}