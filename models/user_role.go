package models

import "gorm.io/gorm"

type UserRole struct {
    gorm.Model
    UserID uint   `json:"user_id" gorm:"not null;index"`
    RoleID uint   `json:"role_id" gorm:"not null;index"`

    // Relationships — like $with in Laravel
    User   User   `json:"user"   gorm:"foreignKey:UserID"`
    Role   Role   `json:"role"   gorm:"foreignKey:RoleID"`
}