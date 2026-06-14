package models

import "gorm.io/gorm"

// Like Laravel's App\Models\User
type User struct {
    gorm.Model                        // adds ID, CreatedAt, UpdatedAt, DeletedAt
    First_name     string `json:"first_name"  gorm:"not null"`
    Last_name     string `json:"last_name"  gorm:"not null"`
    Phone_number     string `json:"phone_number"  gorm:"not null"`
    Address     string `json:"address"  gorm:"not null"`
    National_id     string `json:"national_id"  gorm:"unique; not null"`
    Email    string `json:"email" gorm:"unique; null"`
    Password string `json:"-"     gorm:"not null"` // json:"-" = hidden like $hidden in Laravel

        // HasMany through UserRole — like $user->roles in Laravel
       // HasMany through UserRole
    UserRoles   []UserRole `json:"user_roles"  gorm:"foreignKey:UserID"`
    

}