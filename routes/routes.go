package routes

import (
	v01 "BaseProject/api/controllers/v01"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authController := &v01.AuthController{}

	api := r.Group("/api/v1/auth")
	{
		api.POST("/send-otp-code", authController.SendOtpCode)
		api.POST("/login-with-otp", authController.LoginWithOtp)
		// api.POST("/login-with-password", authController.LoginWithPassword)
	}
}