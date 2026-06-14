package main

import (
    "log"
    "os"

    "BaseProject/config"
    "BaseProject/models"
    "BaseProject/routes"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env — like Laravel's bootstrap process
    if err := godotenv.Load(); err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    // Connect DB
    config.ConnectDB()

    // Auto migrate — like php artisan migrate
    config.DB.AutoMigrate(&models.User{})

    r := gin.Default()
    routes.RegisterRoutes(r)

    port := os.Getenv("APP_PORT")
    r.Run(":" + port)
}