package models 

import "gorm.io/gorm"

type OtpCode struct{
	gorm.Model    
	Code     string `json:"code"  gorm:"not null"`
	PhoneNumber     string `json:"phone_number"  gorm:"not null"`
}