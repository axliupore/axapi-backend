package utils

const (
	Success = 0 // 成功
	Error   = 1 // 失败
	Params  = 2 // 参数错误

	ErrorSystem           = 10001 // 系统错误
	ErrorDatabase         = 10002 // 数据库错误
	ErrorTokenAuth        = 10003 // token 认证错误
	ErrorTokenExpired     = 10004 // token 已经过期
	ErrorTokenNotValidYet = 10005 // token 尚未生效
	ErrorTokenMalformed   = 10006 // token 不符合规范
	ErrorTokenInvalid     = 10007 // token 无法处理
	ErrorRedis            = 10008 // redis 错误
	ErrorUpdate           = 10009 // 修改用户信息失败
	ErrorSearch           = 10010 // 搜索失败
	ErrorAdmin            = 10011 // 权限不足
	ErrorDelete           = 10012 // 删除失败
	ErrorFile             = 10013 // 上传文件失败
	ErrorEmail            = 10014 // 邮件发送失败

	ErrorRegister = 20001 // 注册失败
	ErrorNotLogin = 20002 // 未登录
	ErrorLogin    = 20003 // 登录失败
)
