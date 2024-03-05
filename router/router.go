package router

import (
	_ "github.com/axliupore/axapi/axapi-backend/docs"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/middelware"
	"github.com/axliupore/axapi/axapi-backend/model/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router 初始化总路由
func Router() *gin.Engine {
	if global.Config.Server.Mode == "public" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	if global.Config.Server.Mode != "public" {
		router.Use(gin.Logger())
	}

	if !global.Config.Server.UseOSS {
		router.Static(global.Config.Server.FilePath, "."+global.Config.Server.FilePath)
	}

	router.Use(middleware.Cors())

	if global.Config.Server.Mode == "public" {
		router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	publicGroup := router.Group(global.Config.Server.RouterPrefix)
	{
		// 健康监测
		publicGroup.GET("/health", func(c *gin.Context) {
			response.SuccessMessage(c, "ok")
		})
	}

	{
		InitUserRouter(publicGroup)
		InitEmailRouter(publicGroup)
		InitFileRouter(publicGroup)
	}

	return router
}
