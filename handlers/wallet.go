package handlers

import (
    "minipay/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

type WalletHandler struct {
    WalletService *services.WalletService
}

func NewWalletHandler(walletService *services.WalletService) *WalletHandler {
    return &WalletHandler{WalletService: walletService}
}

func (h *WalletHandler) CreateWallet(c *gin.Context) {
    userID, _ := c.Get("user_id")

    wallet, err := h.WalletService.CreateWallet(userID.(uint))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Wallet created successfully",
        "balance": wallet.Balance,
    })
}

func (h *WalletHandler) GetWallet(c *gin.Context) {
    userID, _ := c.Get("user_id")

    wallet, err := h.WalletService.GetWallet(userID.(uint))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "balance":    wallet.Balance,
        "created_at": wallet.CreatedAt,
    })
}

func (h *WalletHandler) Deposit(c *gin.Context) {
    userID, _ := c.Get("user_id")

    var input struct {
        Amount float64 `json:"amount" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    wallet, err := h.WalletService.Deposit(userID.(uint), input.Amount)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Deposit successful",
        "balance": wallet.Balance,
    })
}