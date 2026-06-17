package v01

import (
	appErrors "BaseProject/api/error"
	"BaseProject/api/controllers"
	"BaseProject/api/service"
	"BaseProject/models"
	"fmt"
	"net/http"
	"BaseProject/config"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadFileController struct {
	controllers.BaseController
	fileService service.FileManagementService
}

// Upload — POST /api/v1/files/upload
func (c *UploadFileController) Upload(ctx *gin.Context) {
	c.Handle(ctx, func() (interface{}, *appErrors.AppError) {

		file, err := ctx.FormFile("file")
		if err != nil {
			return nil, appErrors.BadRequest("No valid file provided for upload")
		}

		fileRecord, err := c.fileService.UploadFile(file, "uploads", "public")
		if err != nil {
			return nil, appErrors.Internal(err.Error())
		}

		fileURL := fmt.Sprintf("%s/%s/%s",
			os.Getenv("APP_URL"),
			fileRecord.Path,
			fileRecord.Title,
		)
	
		return gin.H{
			"file_token": fileRecord.Token,
			"file_url":   fileURL,
		}, nil
	})
}

