package api

import (
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model/response"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"github.com/gin-gonic/gin"
)

// UploadFile
// @Tags      file
// @Summary   上传用户头像
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                     true  "上传文件"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "上传成功"
// @Router    /api/file/upload [post]
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
			response.SuccessDetailed(c, "上传成功", response.File{Url: url})
		}
	} else {
		url, err := utils.SaveFileLocal(fileHeader)
		if err != nil {
			response.Error(c, utils.ErrorFile)
		} else {
			response.SuccessDetailed(c, "上传成功", response.File{Url: url})
		}
	}
}
