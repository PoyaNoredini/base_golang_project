package seeders

import (
	"gorm.io/gorm"
)

func RunAll(db *gorm.DB) {
    (&RolePermissionSeeder{DB: db}).Run()
	(&UserSeeder{DB: db}).Run()
    
}