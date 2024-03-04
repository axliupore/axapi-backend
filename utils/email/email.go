package utils

import (
	"errors"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/utils"
	"gopkg.in/gomail.v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func SendEmail(email string) (string, error) {
	if !verifyEmail(email) {
		return "", errors.New("邮箱格式不正确")
	}
	send := gomail.NewMessage()
	send.SetHeader("From", global.Config.Email.Sender)
	send.SetHeader("To", email)
	send.SetHeader("Subject", global.Config.Email.Title)
	path := filepath.Join("utils", "email", "template.html")
	b, _ := os.ReadFile(path)
	s := string(b)
	code := utils.RandomNumber(6)
	template := strings.ReplaceAll(s, "验证码占位符", code)
	send.SetBody("text/html", template)
	n := gomail.NewDialer(global.Config.Email.Smtp, global.Config.Email.Port, global.Config.Email.User, global.Config.Email.Password)
	if err := n.DialAndSend(send); err != nil {
		return "", errors.New("验证发送失败")
	}
	return code, nil
}

// 验证邮箱是否符合规则
func verifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
