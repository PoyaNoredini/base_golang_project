package models 



import "gorm.io/gorm"

type RolePermission struct {
    gorm.Model
    RoleID uint   `json:"role_id" gorm:"not null;index"`
    PermissionID uint   `json:"permission_id" gorm:"not null;index"`

    // Relationships — like $with in Laravel
    Role   Role   `json:"role"   gorm:"foreignKey:RoleID"`
    Permission   Permission   `json:"permission"   gorm:"foreignKey:PermissionID"`
}