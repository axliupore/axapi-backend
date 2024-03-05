package api

import (
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model/response"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.Error(c, utils.Params)
		return
	}
	if global.Config.Server.UseOSS {
		url, err := utils.UploadFile(fileHeader)
		if err != nil {
			response.Error(c, utils.ErrorFile)
		} else {
			response.SuccessDetailed(c, "上传成功", url)
		}
	} else {
		url, err := utils.SaveFileLocal(fileHeader)
		if err != nil {
			response.Error(c, utils.ErrorFile)
		} else {
			response.SuccessDetailed(c, "上传成功", url)
		}
	}
}
