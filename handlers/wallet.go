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