package controllers

import (
    errors "BaseProject/api/errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
)

type BaseController struct{}

func (c *BaseController) Success(ctx *gin.Context, data interface{}) {
    ctx.JSON(http.StatusOK, gin.H{
        "code":    0,
        "message": "success",
        "data":    data,
    })
}

func (c *BaseController) Error(ctx *gin.Context, code int, message string) {
    ctx.JSON(code, gin.H{
        "code":    code,
        "message": message,
    })
}


func (c *BaseController) Handle(ctx *gin.Context, fn func() (interface{}, *errors.AppError)) {
    data, err := fn()
    if err != nil {
        c.Error(ctx, err.Status, err.Message)
        return
    }
    c.Success(ctx, data)
}


func (c *BaseController) HandleDBError(err error, notFoundMsg string) *errors.AppError {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errors.NotFound(notFoundMsg)
    }
    return errors.Internal("Database error")
}