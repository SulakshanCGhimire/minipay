package config

import (
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"minipay/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := GetDSN()
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.User{})

    DB = db
    log.Println("Database connected successfully")
}