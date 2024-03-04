package router

import (
	"github.com/axliupore/axapi/axapi-backend/api"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(publicGroup *gin.RouterGroup) {
	emailRouter := publicGroup.Group("email")
	{
		emailRouter.POST("send", api.SendEmail)
	}
}
