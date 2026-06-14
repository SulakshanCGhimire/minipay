package config

import "os"

func GetJWTSecret() []byte {
    return []byte(os.Getenv("JWT_SECRET"))
}