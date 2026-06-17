package controllers

import (
    appErrors "BaseProject/api/error"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "errors"
)

type BaseController struct{}

func (c *BaseController) Success(ctx *gin.Context, data interface{}) {
    ctx.JSON(http.StatusOK, gin.H{
        "status":  true,
        "message": "success",
        "data":    data,
    })
}

func (c *BaseController) Error(ctx *gin.Context, code int, message string) {
    ctx.JSON(code, gin.H{
        "status":  false,
        "message": message,
    })
}


func (c *BaseController) Handle(ctx *gin.Context, fn func() (interface{}, *appErrors.AppError)) {
    data, err := fn()
    if err != nil {
        c.Error(ctx, err.Status, err.Message)
        return
    }
    c.Success(ctx, data)
}


func (c *BaseController) HandleDBError(err error, notFoundMsg string) *appErrors.AppError {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return appErrors.NotFound(notFoundMsg)
    }
    return appErrors.Internal("Database error")
}