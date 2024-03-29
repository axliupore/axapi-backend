package service

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model"
	"github.com/axliupore/axapi/axapi-backend/model/request"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
}

// UserRegisterAccount 用户账号注册
func (userService *UserService) UserRegisterAccount(r *request.UserRegisterAccount) error {
	if !errors.Is(global.Db.Where("account = ?", r.Account).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("账号已经存在")
	}
	user := model.NewUserAccount(r.Account)
	user.Password = utils.BcryptHash(r.Password)
	user.Username = r.Username
	return global.Db.Create(user).Error
}

// UserLoginAccount 用户账号登录
func (userService *UserService) UserLoginAccount(r *request.UserLoginAccount) (*model.User, error) {
	var u model.User
	if err := global.Db.Where("account = ?", r.Account).First(&u).Error; err != nil {
		return nil, errors.New("账号不存在")
	}
	if ok := utils.BcryptCheck(r.Password, u.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return &u, nil
}

// GetUser 根据 id 获取用户
func (userService *UserService) GetUser(id int64) (model.User, error) {
	var user model.User
	err := global.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (userService *UserService) UserRegisterEmail(r *request.UserRegisterEmail) error {
	if !errors.Is(global.Db.Where("email = ?", r.Email).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已经存在")
	}
	user := model.NewUserEmail(r.Email)
	user.Account = utils.RandomString(12)
	user.Username = r.Username
	return global.Db.Create(user).Error
}

func (userService *UserService) UserLoginEmail(r *request.UserLoginEmail) (*model.User, error) {
	var u model.User
	if err := global.Db.Where("email = ?", r.Email).First(&u).Error; err != nil {
		return nil, errors.New("邮箱不存在")
	}
	return &u, nil
}

// UserUpdate 更新用户信息
func (userService *UserService) UserUpdate(id int64, r *request.UserUpdate) error {
	return global.Db.Model(&model.User{}).
		Select("update_at", "username", "avatar", "profile", "gender").
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"update_time": time.Now(),
			"username":    r.Username,
			"avatar":      r.Avatar,
			"profile":     r.Profile,
			"gender":      r.Gender,
		}).Error
}
