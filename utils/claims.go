package utils

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model"
	"github.com/gin-gonic/gin"
	"net"
)

var cookieName = "a-token"

// SetToken 用于在 Gin 上下文中设置名为 "a-token" 的 Cookie
// maxAge 参数表示 Cookie 的最大存活时间，单位是秒
func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加 cookie a-token 向来源的 web 添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	// 如果分离失败，则使用整个 Host 作为主机
	if err != nil {
		host = c.Request.Host
	}
	// 设置 Cookie
	c.SetCookie(cookieName, token, maxAge, "/", host, true, false)
}

// GetToken 用于从 Gin 上下文中获取名为 "a-token" 的 Token
// 优先从 Cookie 中获取，如果 Cookie 中没有，则从 Header 中获取
func GetToken(c *gin.Context) string {
	token, _ := c.Cookie(cookieName)
	if token == "" {
		token = c.Request.Header.Get(cookieName)
	}
	return token
}

// ClearToken 用于清除名为 "a-token" 的 Cookie
func ClearToken(c *gin.Context) {
	// 使用 net.SplitHostPort 函数分离主机和端口
	host, _, err := net.SplitHostPort(c.Request.Host)
	// 如果分离失败，则使用整个 Host 作为主机
	if err != nil {
		host = c.Request.Host
	}
	// 设置名为 "a-token" 的 Cookie 的过期时间为负数，即立即删除
	c.SetCookie(cookieName, "", -1, "/", host, true, false)
}

func GetClaims(c *gin.Context) (*model.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != 0 {
		global.Log.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, errors.New(GetMsg(err))
}

func GetUserId(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.Id
		}
	} else {
		waitUse := claims.(*model.CustomClaims)
		return waitUse.Id
	}
}

func GetUserRole(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Role
		}
	} else {
		waitUse := claims.(*model.CustomClaims)
		return waitUse.Role
	}
}
