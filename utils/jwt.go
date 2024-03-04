package utils

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT 结构体定义了 JWT 相关的操作
type JWT struct {
	SigningKey []byte // JWT 签名密钥
}

// NewJWT 返回一个新的 JWT 实例，使用全局配置中的签名密钥
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(global.Config.JWT.SigningKey),
	}
}

// CreateToken 根据传入的 claims 构造一个 token
func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateClaims 组装 JWT 所需的声明信息
func (j *JWT) CreateClaims(baseClaims model.BaseClaims) model.CustomClaims {
	// 从全局配置中解析缓冲时间（BufferTime）和过期时间（ExpiresTime）
	bf, _ := ParseDuration(global.Config.JWT.BufferTime)
	ep, _ := ParseDuration(global.Config.JWT.ExpiresTime)
	claims := model.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"gin_template"}, // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间 7天  配置文件
			Issuer:    global.Config.JWT.Issuer,               // 签名的发行者
		},
	}
	return claims
}

// ParseToken 解析 JWT 令牌，返回 CustomClaims 和状态码
func (j *JWT) ParseToken(token string) (*model.CustomClaims, int) {
	// 解析 token 是否为有效的 token
	tokenClaims, err := jwt.ParseWithClaims(token, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	// 如果解析 JWT 令牌时发生了错误
	if err != nil {
		// 声明一个指向 jwt.ValidationError 类型的指针 ve
		var v *jwt.ValidationError
		// 使用 errors 包的 As 函数，尝试将 err 转换为 *jwt.ValidationError 类型
		if errors.As(err, &v) {
			// 根据 ValidationError 中的错误位进行不同的错误处理
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				// 令牌格式错误
				return nil, ErrorTokenMalformed
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				// 令牌已过期
				return nil, ErrorTokenExpired
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				// 令牌尚未生效
				return nil, ErrorTokenNotValidYet
			} else {
				// 无法处理的其他令牌错误
				return nil, ErrorTokenInvalid
			}
		}
	}
	// 如果 tokenClaims 不为 nil，且令牌有效，则返回 CustomClaims 和状态码 Success
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*model.CustomClaims); ok && tokenClaims.Valid {
			return claims, Success
		}
		return nil, ErrorTokenInvalid
	} else {
		return nil, ErrorTokenInvalid
	}
}
