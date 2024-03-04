package service

import (
	"context"
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	utils "github.com/axliupore/axapi/axapi-backend/utils/email"
	"time"
)

type EmailService struct {
}

func (emailService *EmailService) SendEmail(email string) error {
	// 1.检测 1 分钟之内是否已经发送过了
	if code, _ := emailService.GetCode(email); code != "" {
		return errors.New("请不要重复申请验证码")
	}
	// 2.生成验证码
	code, err := utils.SendEmail(email)
	if err != nil {
		return err
	}
	// 3.存入到 redis 中
	if err := global.Redis.Set(context.Background(), email, code, 1*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (emailService *EmailService) GetCode(email string) (string, error) {
	code, err := global.Redis.Get(context.Background(), email).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

// RemoveEmail 将 email 从 redis 中删除
func (emailService *EmailService) RemoveEmail(email string) {
	global.Redis.Del(context.Background(), email)
}
