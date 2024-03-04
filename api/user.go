package api

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model"
	"github.com/axliupore/axapi/axapi-backend/model/request"
	"github.com/axliupore/axapi/axapi-backend/model/response"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

// UserRegisterAccount
// @Tags user
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.UserRegisterAccount      true  "账号,密码,确认密码,用户名"
// @Success  200   {object}  response.Response{msg=string}  		"返回注册信息"
// @Router   /api/user/register/account [post]
func UserRegisterAccount(c *gin.Context) {
	var r *request.UserRegisterAccount
	err := c.ShouldBindJSON(&r)
	if err != nil || utils.IsAnyBlank(r.Account, r.Password, r.CheckPassword) {
		response.Error(c, utils.Params)
		return
	}
	if len(r.Account) < 4 {
		response.ErrorMessage(c, utils.Params, "账号不能小于4位")
		return
	}
	if len(r.Password) < 8 || len(r.CheckPassword) < 8 || r.Password != r.CheckPassword {
		response.ErrorMessage(c, utils.Params, "密码不能小于8位或两次密码不一致")
		return
	}
	err = userService.UserRegisterAccount(r)
	if err != nil {
		global.Log.Error("注册失败!", zap.Error(err))
		response.ErrorMessage(c, utils.ErrorRegister, err.Error())
		return
	}
	response.SuccessMessage(c, "注册成功")
}

// UserLoginAccount
// @Tags     user
// @Summary  用户登录
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.UserLoginAccount      				                  true  "账号,密码"
// @Success  200   {object}  response.Response{data=response.UserLoginResponse,msg=string}  "返回用户信息,token,过期时间"
// @Router   /api/user/login/account [post]
func UserLoginAccount(c *gin.Context) {
	var r *request.UserLoginAccount
	err := c.ShouldBindJSON(&r)
	if err != nil || utils.IsAnyBlank(r.Account, r.Password) {
		response.Error(c, utils.Params)
		return
	}
	if len(r.Account) < 4 {
		response.ErrorMessage(c, utils.Params, "账号不能小于4位")
		return
	}
	user, err := userService.UserLoginAccount(r)
	if err != nil {
		global.Log.Error("登录失败!", zap.Error(err))
		response.ErrorMessage(c, utils.ErrorRegister, err.Error())
		return
	}
	if user.Status != 0 {
		global.Log.Error("登陆失败！用户被禁止登录")
		response.ErrorMessage(c, utils.Params, "登陆失败！用户被禁止登录")
		return
	}
	tokenNext(c, *user)
}

// tokenNext 登陆后签发 JWT
func tokenNext(c *gin.Context, user model.User) {
	j := utils.NewJWT() // 唯一签名
	claims := j.CreateClaims(model.BaseClaims{
		Id:       user.Id,
		Account:  user.Account,
		Username: user.Username,
		Role:     user.Role,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Log.Error("获取token失败!", zap.Error(err))
		response.ErrorMessage(c, utils.ErrorTokenInvalid, "获取token失败!")
		return
	}
	// 如果没有使用多点登录，直接在当前登录的设备上设置 token
	if !global.Config.Server.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.SuccessDetailed(c, "登录成功", response.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		})
		return
	}

	// 如果启用了多点登录
	if jwtStr, err := jwtService.GetRedisJWT(user.Account); errors.Is(err, redis.Nil) {
		if err := jwtService.SetRedisJWT(token, user.Account); err != nil {
			global.Log.Error("设置登录状态失败!", zap.Error(err))
			response.ErrorMessage(c, utils.ErrorRedis, "设置登录状态失败")
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.SuccessDetailed(c, "登录成功", response.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		})
	} else if err != nil {
		global.Log.Error("设置登录状态失败!", zap.Error(err))
		response.ErrorMessage(c, utils.ErrorRedis, "设置登录状态失败")
		return
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.ErrorMessage(c, utils.ErrorRedis, "设置登录状态失败")
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.SuccessDetailed(c, "登录成功", response.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		})
	}
}

// GetUser
// @Tags      user
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=model.User,msg=string}  "返回用户信息"
// @Router    /api/user/get [get]
func GetUser(c *gin.Context) {
	id := utils.GetUserId(c)
	loginUser, err := userService.GetUser(id)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.ErrorMessage(c, utils.ErrorNotLogin, "获取失败")
		return
	}
	response.SuccessDetailed(c, "获取成功", loginUser)
}
