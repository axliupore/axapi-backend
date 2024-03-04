package config

import "github.com/redis/go-redis/v9"

// Redis 配置
type Redis struct {
	Addr     string `mapstructure:"addr"  yaml:"addr"`        // 服务器地址:端口
	Password string `mapstructure:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" yaml:"db"`             // redis的哪个数据库
}

func (r *Redis) Options() *redis.Options {
	return &redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	}
}
