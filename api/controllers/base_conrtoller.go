package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type BaseController struct{}

// ctx passed directly — no more storing it in struct
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