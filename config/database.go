package config

import (
    "fmt"
    "log"
    "os"
    "BaseProject/models"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("❌ Error loading .env file")
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }
    err = db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.OtpCode{},
		&models.UserRole{},
		&models.RolePermission{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

    log.Println("✅ Database connected successfully")
    DB = db
}