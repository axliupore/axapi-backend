package router

import (
	"github.com/axliupore/axapi/axapi-backend/api"
	middleware "github.com/axliupore/axapi/axapi-backend/middelware"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(publicGroup *gin.RouterGroup) {
	fileRouter := publicGroup.Group("file")
	{
		fileRouter.Use(middleware.JWT())
		fileRouter.POST("upload", api.UploadFile)
	}
}
