package v01

import (
    errors "BaseProject/api/errors"
    "BaseProject/api/controllers"
    "BaseProject/api/helper"
    "BaseProject/api/validations"
    "BaseProject/config"
    "BaseProject/models"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    controllers.BaseController
}

func (c *AuthController) SendOtpCode(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *errors.AppError) {
        var req validations.SendOtpRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, errors.BadRequest(err.Error())
        }

        otpCode := helper.GenerateOtpCode()


        return c.Success(ctx, gin.H{"message": "OTP code sent successfully", "otp_code": otpCode})
    })
}

func (c *AuthController) LoginWithOtp(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *errors.AppError) {
        var req validations.LoginWithOtpRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, errors.BadRequest(err.Error())
        }

        if !helper.VerifyCode(req.Code, req.Phone) {
            return nil, errors.Unauthorized("Invalid OTP code")
        }

        var user models.User
        if err := config.DB.Where("phone_number = ?", req.Phone).First(&user).Error; err != nil {
            return nil, c.HandleDBError(err, "User not found")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, errors.Internal("Internal server error")
        }

        return c.Success(ctx, gin.H{"message": "Login successful", "token": token, "user": user})
    })
}

func (c *AuthController) LoginWithPassword(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *errors.AppError) {
        var req validations.LoginWithPasswordRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, errors.BadRequest(err.Error())
        }

        var user models.User
        if err := config.DB.Where("phone_number = ?", req.Phone).First(&user).Error; err != nil {
            return nil, c.HandleDBError(err, "User not found")
        }

        if !helper.CheckPassword(req.Password, user.Password) {
            return nil, errors.Unauthorized("Invalid password")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, errors.Internal("Internal server error")
        }

        return c.Success(ctx, gin.H{"message": "Login successful", "token": token, "user": user})
    })
}

func (c *AuthController) Register(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *errors.AppError) {
        var req validations.RegisterRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, errors.BadRequest(err.Error())
        }

        password, err := helper.HashPassword(req.Password)
        if err != nil {
            return nil, errors.Internal("Failed to hash password")
        }

        user := models.User{
            First_name:   req.FirstName,
            Last_name:    req.LastName,
            Phone_number: req.Phone,
            National_id:  req.NationalID,
            Email:        req.Email,
            Password:     password,
        }

        if err := config.DB.Create(&user).Error; err != nil {
            return nil, errors.Internal("Failed to create user")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, errors.Internal("Internal server error")
        }

        return c.Success(ctx, gin.H{"message": "User created successfully", "token": token, "user": user})
    })
}