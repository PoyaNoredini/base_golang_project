package routes

import (
    v01 "BaseProject/api/controllers/v01"
    "github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(api *gin.RouterGroup) {
    auth := api.Group("/auth")
    c := &v01.AuthController{}
    {
        auth.POST("/send-otp-code", c.SendOtpCode)
        auth.POST("/login-with-otp", c.LoginWithOtp)
        auth.POST("/login-with-password", c.LoginWithPassword)
        auth.POST("/register", c.Register)
    }
}