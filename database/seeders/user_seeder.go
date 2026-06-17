package seeders

import (
	"BaseProject/models"
	"BaseProject/api/helper"
	"log"
	"gorm.io/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func (s *UserSeeder) Run() {
	var role models.Role
	s.DB.Where(models.Role{Title: "Administrator"}).First(&role)
	password, err := helper.HashPassword("password")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	users := []models.User{
		{
			First_name: "Admin",
			Last_name: "Admin",
			Phone_number: "09372718990",
			Email: "admin@example.com",
			Password: password,
			National_id: "0010000000",
			Address: " Tehran, Iran",
			UserRoles: []models.UserRole{
				{
					RoleID: role.ID,
				},
			},
		},
	}

	for _, user := range users {
		s.DB.Create(&user)
	}
}