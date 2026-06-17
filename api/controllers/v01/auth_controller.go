package v01

import (
    appErrors "BaseProject/api/error"
    "BaseProject/api/controllers"
    "BaseProject/api/helper"
    "BaseProject/api/validations"
    "BaseProject/config"
    "BaseProject/models"
    "BaseProject/api/resource"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    controllers.BaseController
}

func (c *AuthController) SendOtpCode(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
        var req validations.SendOtpRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, appErrors.BadRequest(err.Error())
        }

        otpCode := helper.GenerateOtpCode()


        return gin.H{"message": "OTP code sent successfully", "otp_code": otpCode}, nil
    })
}

func (c *AuthController) LoginWithOtp(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
        var req validations.LoginWithOtpRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, appErrors.BadRequest(err.Error())
        }

        if !helper.VerifyCode(req.Code, req.Phone) {
            return nil, appErrors.Unauthorized("Invalid OTP code")
        }

        var user models.User
        if err := config.DB.Where("phone_number = ?", req.Phone).First(&user).Error; err != nil {
            return nil, c.HandleDBError(err, "User not found")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, appErrors.Internal("Internal server error")
        }

        return gin.H{"message": "Login successful", "token": token, "user": resource.NewUserResource(user)}, nil
    })
}

func (c *AuthController) LoginWithPassword(ctx *gin.Context) {
    c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
        var req validations.LoginWithPasswordRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, appErrors.BadRequest(err.Error())
        }

        var user models.User
        if err := config.DB.Preload("UserRoles.Role").Where("phone_number = ?", req.Phone).First(&user).Error; err != nil {
            return nil, c.HandleDBError(err, "User not found")
        }

        if !helper.CheckPassword(req.Password, user.Password) {
            return nil, appErrors.Unauthorized("Invalid password")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, appErrors.Internal("Internal server error")
        }

        return gin.H{"message": "Login successful", "token": token, "user": resource.NewUserResource(user)}, nil
    })
}

func (c *AuthController) Register(ctx *gin.Context) {
        c.Handle(ctx, func() (interface{}, *appErrors.AppError) {
        var req validations.RegisterRequest
        if err := ctx.ShouldBindJSON(&req); err != nil {
            return nil, appErrors.BadRequest(err.Error())
        }

        password, err := helper.HashPassword(req.Password)
        if err != nil {
            return nil, appErrors.Internal("Failed to hash password")
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
            return nil, appErrors.Internal("Failed to create user")
        }

        token, err := helper.GenerateToken(user.ID, user.Phone_number)
        if err != nil {
            return nil, appErrors.Internal("Internal server error")
        }

        return gin.H{"message": "User created successfully", "token": token, "user": resource.NewUserResource(user)}, nil
    })
}