package main

import (
    "log"
    "os"

    "BaseProject/config"
    "BaseProject/models"
    "BaseProject/routes"
    "BaseProject/database/seeders" 

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
        // go run main.go --seed
    if len(os.Args) > 1 && os.Args[1] == "--seed" {
            seeders.RunAll(config.DB)
            return
    }


    // Auto migrate — like php artisan migrate
    config.DB.AutoMigrate(&models.User{})

    r := gin.Default()
    routes.RegisterRoutes(r)

    port := os.Getenv("APP_PORT")
    r.Run(":" + port)
}