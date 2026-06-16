package helper

import (
	"fmt"
	"math/rand"
	"time"

	"BaseProject/config"
	"BaseProject/models"
)

func GenerateOtpCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func VerifyCode(code string, mobile string) bool {
    var otp models.OtpCode

    result := config.DB.
        Where("code = ? AND phone_number = ?", code, mobile).
        Where("created_at >= ?", time.Now().Add(-2*time.Minute)). // expires after 2 min
        Where("deleted_at IS NULL").
        First(&otp)

    if result.Error != nil {
        return false
    }

    // Soft delete after use — so it can't be reused (like invalidating token)
    config.DB.Delete(&otp)
	//

    return true
}

func InsertCode(code string, mobile string) error {
	otp := models.OtpCode{
		Code:         code,
		PhoneNumber: mobile,
	}
	return config.DB.Create(&otp).Error
}