package middleware

import (
    appErrors "BaseProject/api/error"
    "BaseProject/api/helper"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        token := extractToken(ctx)
        if token == "" {
            respondAndAbort(ctx, appErrors.Unauthorized("Authorization token is required"))
            return
        }

        claims, err := helper.ParseToken(token)
        if err != nil {
            switch {
            case err == helper.ErrTokenExpired:
                respondAndAbort(ctx, appErrors.Unauthorized("Token has expired"))
            case err == helper.ErrTokenInvalid:
                respondAndAbort(ctx, appErrors.Unauthorized("Invalid token"))
            default:
                respondAndAbort(ctx, appErrors.Internal("Authentication error"))
            }
            return
        }

        // Inject into context — accessible in any controller
        ctx.Set("user_id", claims.UserID)
        ctx.Set("user_phone", claims.Phone)

        ctx.Next()
    }
}

// extractToken supports both "Bearer <token>" and raw token headers
func extractToken(ctx *gin.Context) string {
    bearer := ctx.GetHeader("Authorization")
    if len(bearer) > 7 && bearer[:7] == "Bearer " {
        return bearer[7:]
    }
    return ""
}

func respondAndAbort(ctx *gin.Context, err *appErrors.AppError) {
    ctx.AbortWithStatusJSON(err.Status, gin.H{
        "code":    err.Status,
        "message": err.Message,
    })
}