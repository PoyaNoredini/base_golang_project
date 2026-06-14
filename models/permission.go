package models


import "gorm.io/gorm"

// Like Laravel's App\Models\User
type Permission struct {
    gorm.Model                        // adds ID, CreatedAt, UpdatedAt, DeletedAt
    Title     string `json:"title"  gorm:"not null"`
    Discription    string `json:"discription"  gorm:"not null"`

	
	RolePermissions   []RolePermission `json:"role_permissions"  gorm:"foreignKey:PermissionID"`
}