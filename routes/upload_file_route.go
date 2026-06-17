package routes

import (
    v01 "BaseProject/api/controllers/v01"
    "github.com/gin-gonic/gin"
)

func RegisterFileRoutes(api *gin.RouterGroup, middleware ...gin.HandlerFunc) {
    protected := api.Group("/")
    protected.Use(middleware...)
    {
        c := &v01.UploadFileController{}
        protected.POST("/upload-file", c.Upload)
    }
}