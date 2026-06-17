package middleware

import (
    appErrors "BaseProject/api/error"
    "BaseProject/config"
    "BaseProject/models"
    "github.com/gin-gonic/gin"
)

func RequirePermission(permission string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // 1. get user_id from context — Auth middleware already set this
        userID, exists := ctx.Get("user_id")
        if !exists {
            respondAndAbort(ctx, appErrors.Unauthorized("Unauthorized access!"))
            return
        }

        // 2. check if the permission exists at all
        var permissionRecord models.Permission
        if err := config.DB.Where("title = ?", permission).First(&permissionRecord).Error; err != nil {
            // permission title doesn't exist in the system
            respondAndAbort(ctx, appErrors.Forbidden("Access denied!"))
            return
        }

        // 3. check if user has any role that has this permission
        //    users -> user_roles -> roles -> role_permissions -> permissions
        var count int64
        config.DB.Table("users").
            Joins("JOIN user_roles ON user_roles.user_id = users.id").
            Joins("JOIN roles ON roles.id = user_roles.role_id").
            Joins("JOIN role_permissions ON role_permissions.role_id = roles.id").
            Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
            Where("users.id = ? AND permissions.title = ?", userID, permission).
            Count(&count)

        if count == 0 {
            respondAndAbort(ctx, appErrors.Forbidden("Access denied!"))
            return
        }

        ctx.Next()
    }
}