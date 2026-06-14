package main

import (
    "minipay/config"
    "minipay/handlers"

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

    r.Run(":8080")
}