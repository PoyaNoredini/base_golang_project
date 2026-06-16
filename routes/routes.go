package routes

import (
	v01 "BaseProject/api/controllers/v01"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authController := &v01.AuthController{}

	api := r.Group("/api/v1")
	{
		api.Group("/auth") {
			aut.POST("/send-otp-code", authController.SendOtpCode)
			aut.POST("/login-with-otp", authController.LoginWithOtp)
			aut.POST("/login-with-password", authController.LoginWithPassword)
			aut.POST("/register", authController.Register)
		}
	}


}