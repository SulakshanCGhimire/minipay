package main

import (
    "minipay/config"
    "minipay/handlers"
    "minipay/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    config.ConnectDB()

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
        protected.POST("/wallet/create", handlers.CreateWallet)
        protected.POST("/wallet/deposit", handlers.Deposit)
        protected.GET("/wallet", handlers.GetWallet)
        protected.POST("/transfer", handlers.Transfer)
    }

    r.Run(":8080")
}