package handlers

import (
    "minipay/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type TransactionHandler struct {
    TransactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
    return &TransactionHandler{TransactionService: transactionService}
}

func (h *TransactionHandler) Transfer(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var input struct {
        ReceiverWalletID uint    `json:"receiver_wallet_id" binding:"required"`
        Amount           float64 `json:"amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.TransactionService.Transfer(userID.(uint), input.ReceiverWalletID, input.Amount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
    userID, _ := c.Get("user_id")

    transactions, err := h.TransactionService.GetTransactions(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}