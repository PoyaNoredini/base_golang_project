package routes

import (
    "BaseProject/api/controllers/v01"
    "BaseProject/api/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterUserRoutes(api *gin.RouterGroup, authMiddleware ...gin.HandlerFunc) {
    protected := api.Group("/")
    protected.Use(authMiddleware...) // Auth middleware applied to all routes
    {
        userController := v01.UserController{}

        protected.GET("/users",              middleware.RequirePermission("view-user"),   userController.Index)
        protected.POST("/users",             middleware.RequirePermission("create-user"), userController.Create)
        protected.GET("/users/profile",      userController.Profile)
        protected.PUT("/users/update",       middleware.RequirePermission("update-user"), userController.UpdateUser)
        protected.PUT("/users/admin/update/:id", middleware.RequirePermission("update-user"), userController.AdminUpdateUser)
        protected.DELETE("/users/delete/:id",    middleware.RequirePermission("delete-user"), userController.DeleteUser)
    }
}