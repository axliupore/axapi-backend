package api

import (
	"github.com/axliupore/axapi/axapi-backend/model/request"
	"github.com/axliupore/axapi/axapi-backend/model/response"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"github.com/gin-gonic/gin"
)

// SendEmail
// @Tags email
// @Summary  获取验证码
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.Email        			true  "邮箱"
// @Success  200   {object}  response.Response{msg=string}  	  "返回信息"
// @Router   /api/email/send [post]
func SendEmail(c *gin.Context) {
	var r *request.Email
	err := c.ShouldBindJSON(&r)
	if err != nil || utils.IsAnyBlank(r.Email) {
		response.Error(c, utils.Params)
		return
	}
	err = emailService.SendEmail(r.Email)
	if err != nil {
		response.ErrorMessage(c, utils.ErrorEmail, err.Error())
		return
	}
	response.SuccessMessage(c, "发送成功")
}
