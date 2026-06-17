package models

import "gorm.io/gorm"

type UploadFiles struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Extension string `json:"extension" gorm:"not null"`
	Token string `json:"token" gorm:"not null"`
	Path string `json:"path" gorm:"not null"`
	Size int64 `json:"size" gorm:"not null"`
	Storage string `json:"storage" gorm:"not null"`
	MimeType string `json:"mime_type" gorm:"not null"`
}