package service

import (
	"BaseProject/config"
	"BaseProject/models"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileManagementService struct{}

// UploadFile — equivalent to your Laravel uploadFile()
func (s *FileManagementService) UploadFile(
	file *multipart.FileHeader,
	basePath string,
	storage string,
) (*models.UploadFiles, error) {

	// 1. open the uploaded file 
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open uploaded file")
	}
	defer src.Close()

	// 2. build the directory path —
	datePath := time.Now().Format("2006-01-02") // Go's time format (equivalent to Y-m-d)
	dirPath := filepath.Join("storage", storage, basePath, datePath)

	// 3. create directory if it doesn't exist 
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, errors.New("failed to create storage directory")
	}

	// 4. generate unique filename 
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return nil, errors.New("file has no extension")
	}
	randomPart, err := generateRandom(5)
	if err != nil {
		return nil, errors.New("failed to generate filename")
	}
	fileName := fmt.Sprintf("%s-%d%s", randomPart, time.Now().Unix(), ext)
	fullPath := filepath.Join(dirPath, fileName)

	// 5. create the file on disk
	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, errors.New("failed to create file on disk")
	}
	defer dst.Close()

	// 6. copy contents
	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(fullPath) // cleanup on failure
		return nil, errors.New("failed to write file")
	}

	// 7. get mime type 
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream" // fallback
	}

	// 8. generate unique token 
	token, err := s.generateFileToken()
	if err != nil {
		os.Remove(fullPath)
		return nil, errors.New("failed to generate file token")
	}

	// 9. save to DB inside a transaction 
	tx := config.DB.Begin()
	if tx.Error != nil {
		os.Remove(fullPath)
		return nil, errors.New("failed to start transaction")
	}

	fileRecord := &models.UploadFiles{
		Title:     fileName,
		Extension: ext[1:], // strip the leading dot — "jpg" not ".jpg"
		Token:     token,
		Path:      dirPath,
		Size:      file.Size,
		Storage:   storage,
		MimeType:  mimeType,
	}

	if err := tx.Create(fileRecord).Error; err != nil {
		tx.Rollback()
		os.Remove(fullPath) // cleanup file if DB fails
		return nil, errors.New("failed to save file record")
	}

	tx.Commit()
	return fileRecord, nil
}

// GetFilePath — builds the full path to the file on disk
func (s *FileManagementService) GetFilePath(fileRecord *models.UploadFiles) string {
	return filepath.Join(fileRecord.Path, fileRecord.Title)
}

// FileExists — checks if the physical file exists
func (s *FileManagementService) FileExists(fileRecord *models.UploadFiles) bool {
	_, err := os.Stat(s.GetFilePath(fileRecord))
	return !os.IsNotExist(err)
}

// generateFileToken —
func (s *FileManagementService) generateFileToken() (string, error) {
	for range 10 { // max 10 attempts — avoids infinite recursion
		token, err := generateRandom(13) // 13 bytes hex = 26 chars
		if err != nil {
			return "", err
		}
		var count int64
		config.DB.Model(&models.UploadFiles{}).Where("token = ?", token).Count(&count)
		if count == 0 {
			return token, nil
		}
	}
	return "", errors.New("failed to generate unique token after 10 attempts")
}

// generateRandom 
func generateRandom(bytes int) (string, error) {
	b := make([]byte, bytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}