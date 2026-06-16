package v01

import (
	"BaseProject/api/controllers"
	"BaseProject/api/helper"
	"BaseProject/api/validations"
	"BaseProject/config"
	"BaseProject/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	controllers.BaseController
}

func (c *AuthController) SendOtpCode(ctx *gin.Context) {
	var request validations.SendOtpRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	otpCode := helper.GenerateOtpCode()
	if err := helper.InsertCode(otpCode, request.PhoneNumber); err != nil {
		c.Error(ctx, http.StatusInternalServerError, "Failed to save OTP code")
		return
	}

	c.Success(ctx, gin.H{"message": "OTP code sent successfully", "otp_code": otpCode})
}

func (c *AuthController) LoginWithOtp(ctx *gin.Context) {
	var request validations.LoginWithOtpRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !helper.VerifyCode(request.Code, request.Phone) {
		c.Error(ctx, http.StatusUnauthorized, "Invalid OTP code")
		return
	}

	var user models.User
	result := config.DB.Where("phone_number = ?", request.Phone).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Error(ctx, http.StatusNotFound, "User not found")
		return
	}

	if result.Error != nil {
		c.Error(ctx, http.StatusInternalServerError, "Database error")
		return
	}

	jwtToken, err := helper.GenerateToken(user.ID, user.Phone_number)
	if err != nil {
		c.Error(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.Success(ctx, gin.H{"message": "Login successful", "token": jwtToken, "user": user})
}

func (c *AuthController) LoginWithPassword(ctx *gin.Context) {
    var request validations.LoginWithPasswordRequest

    if err := ctx.ShouldBindJSON(&request); err != nil {
        c.Error(ctx, http.StatusBadRequest, err.Error())
        return
    }

    // Fetch user FIRST
    var user models.User
    result := config.DB.Where("phone_number = ?", request.Phone).First(&user)

    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        c.Error(ctx, http.StatusNotFound, "User not found")
        return
    }

    if !helper.CheckPassword(request.Password, user.Password) {
        c.Error(ctx, http.StatusUnauthorized, "Invalid password")
        return
    }

    jwtToken, err := helper.GenerateToken(user.ID, user.Phone_number)
    if err != nil {
        c.Error(ctx, http.StatusInternalServerError, "Internal server error")
        return
    }

    c.Success(ctx, gin.H{"message": "Login successful", "token": jwtToken, "user": user})
}


func (c *AuthController) Register(ctx *gin.Context) {
	var request validations.RegisterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		c.Error(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		First_name: request.FirstName,
		Last_name: request.LastName,
		Phone_number: request.Phone,
		National_id: request.NationalID,
		Email: request.Email,
		Password: password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.Error(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}
	jwtToken, err := helper.GenerateToken(user.ID, user.Phone_number)
    if err != nil {
        c.Error(ctx, http.StatusInternalServerError, "Internal server error")
        return
    }

	c.Success(ctx, gin.H{"message": "User created successfully", "token": jwtToken, "user": user})
}