package router

import (
	"github.com/axliupore/axapi/axapi-backend/api"
	middleware "github.com/axliupore/axapi/axapi-backend/middelware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(publicGroup *gin.RouterGroup) {
	userRouter := publicGroup.Group("user")
	{
		// 账号注册
		userRouter.POST("register/account", api.UserRegisterAccount)
		// 邮箱注册
		userRouter.POST("register/email")
		// 账号登录
		userRouter.POST("login/account", api.UserLoginAccount)
		// 邮箱登录
		userRouter.POST("login/email")
		// 获取用户的登录态
		userRouter.Use(middleware.JWT())
		userRouter.GET("get", api.GetUser)
	}
}
