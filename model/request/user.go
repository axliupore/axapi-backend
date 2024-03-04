package request

// UserRegisterAccount 用户账号注册请求体
type UserRegisterAccount struct {
	Account       string `json:"account"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	CheckPassword string `json:"checkPassword"`
}

// UserLoginAccount 用户账号登录请求体
type UserLoginAccount struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
