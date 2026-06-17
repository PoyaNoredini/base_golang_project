package routes

import (
    "BaseProject/api/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.Static("/storage", "./storage")

    api := r.Group("/api/v1")

    RegisterAuthRoutes(api)
    RegisterFileRoutes(api, middleware.Auth())
    RegisterUserRoutes(api, middleware.Auth())
}