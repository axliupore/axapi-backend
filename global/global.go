package global

import (
	"github.com/axliupore/axapi/axapi-backend/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db         *gorm.DB          // 数据库
	Config     config.Config     // 配置信息
	Log        *zap.Logger       // 日志
	Redis      *redis.Client     // redis 客户端
	BlackCache local_cache.Cache // 本地缓存，用于缓存 JWT 的黑名单
)
