package model

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

// BaseClaims 声明用于在 JWT 中携带用户相关的信息
type BaseClaims struct {
	Id       int64
	Account  string
	Username string
	Role     string
}

type JwtBlacklist struct {
	Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
