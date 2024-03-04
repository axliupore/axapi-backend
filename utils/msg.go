package utils

var msgFlags = map[int]string{
	Success: "ok",
	Error:   "error",
	Params:  "参数错误",

	ErrorSystem:           "系统错误",
	ErrorDatabase:         "数据库错误",
	ErrorTokenAuth:        "token 认证错误",
	ErrorTokenExpired:     "token 已经过期",
	ErrorTokenNotValidYet: "token 尚未生效",
	ErrorTokenMalformed:   "token 不符合规范",
	ErrorTokenInvalid:     "token 无法处理",
	ErrorRedis:            "redis 错误",
	ErrorUpdate:           "修改用户信息失败",
	ErrorSearch:           "搜索失败",
	ErrorAdmin:            "权限不足",
	ErrorDelete:           "删除失败",
	ErrorFile:             "上传文件失败",

	ErrorNotLogin: "未登录",
	ErrorLogin:    "登录失败",
	ErrorRegister: "注册失败",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if !ok {
		return msgFlags[Error]
	}
	return msg
}
