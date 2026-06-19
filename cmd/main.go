package main

import (
    "minipay/config"
    "minipay/handlers"
    "minipay/middleware"
    "minipay/repositories"
    "minipay/services"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    config.ConnectDB()

    // Repositories
    walletRepo := repositories.NewWalletRepository(config.DB)
    transactionRepo := repositories.NewTransactionRepository(config.DB)

    // Services
    walletService := services.NewWalletService(walletRepo)
    transactionService := services.NewTransactionService(walletRepo, transactionRepo)

    // Handlers
    walletHandler := handlers.NewWalletHandler(walletService)
    transactionHandler := handlers.NewTransactionHandler(transactionService)

    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    protected := r.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/me", func(c *gin.Context) {
            userID, _ := c.Get("user_id")
            c.JSON(200, gin.H{"user_id": userID})
        })
        protected.POST("/wallet/create", walletHandler.CreateWallet)
        protected.GET("/wallet", walletHandler.GetWallet)
        protected.POST("/wallet/deposit", walletHandler.Deposit)
        protected.POST("/transfer", transactionHandler.Transfer)
        protected.GET("/transactions", transactionHandler.GetTransactions)
    }

    r.Run(":8080")
}