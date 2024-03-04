package service

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/model"
	"github.com/axliupore/axapi/axapi-backend/model/request"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"gorm.io/gorm"
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
	err := global.Db.Create(user).Error
	return err
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
