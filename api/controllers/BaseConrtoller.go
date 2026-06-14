package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	Context *gin.Context
}

func (c *BaseController) Success(data interface{}) {
	c.Context.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": data})
}

func (c *BaseController) Error(code int, message string) {
	c.Context.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}