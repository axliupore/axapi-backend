package response

import (
	"github.com/axliupore/axapi/axapi-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Success(c *gin.Context) {
	result(c, utils.Success, nil, "success")
}

func SuccessMessage(c *gin.Context, msg string) {
	result(c, utils.Success, nil, msg)
}

func SuccessData(c *gin.Context, data interface{}) {
	result(c, utils.Success, data, "success")
}

func SuccessDetailed(c *gin.Context, msg string, data interface{}) {
	result(c, utils.Success, data, msg)
}

func Error(c *gin.Context, code int) {
	result(c, code, nil, utils.GetMsg(code))
}

func ErrorMessage(c *gin.Context, code int, msg string) {
	result(c, code, nil, msg)
}

func ErrorDetailed(c *gin.Context, code int, data interface{}, msg string) {
	result(c, code, data, msg)
}
