package api

import "github.com/axliupore/axapi/axapi-backend/service"

var (
	jwtService   = service.JwtService{}
	userService  = service.UserService{}
	emailService = service.EmailService{}
)
