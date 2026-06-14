package models


import "gorm.io/gorm"

// Like Laravel's App\Models\User
type Role struct {
    gorm.Model                        // adds ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string `json:"title"  gorm:"not null"`
    Discription    string `json:"discription"  gorm:"not null"`

	    // HasMany through UserRole
    UserRoles   []UserRole `json:"user_roles"  gorm:"foreignKey:RoleID"`
	RolePermissions   []RolePermission `json:"role_permissions"  gorm:"foreignKey:RoleID"`
}