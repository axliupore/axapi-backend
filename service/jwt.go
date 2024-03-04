package service

import (
	"context"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/utils"
)

type JwtService struct{}

// SetRedisJWT JWT 存入 redis 并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, username string) error {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.Config.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.Redis.Set(context.Background(), username, jwt, timer).Err()
	return err
}

// GetRedisJWT 从 redis 中获取 JWT
func (jwtService *JwtService) GetRedisJWT(username string) (string, error) {
	return global.Redis.Get(context.Background(), username).Result()
}
