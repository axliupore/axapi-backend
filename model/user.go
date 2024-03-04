package model

type User struct {
	Id             int64  `db:"id" json:"id"`
	Account        string `db:"account" json:"account"`
	Username       string `db:"username" json:"username"`
	Password       string `db:"password" json:"-"`
	Avatar         string `db:"avatar" json:"avatar"`
	Email          string `db:"email" json:"email"`
	Phone          string `db:"phone" json:"phone"`
	Profile        string `db:"profile" json:"profile"`
	Gender         int8   `db:"gender" json:"gender"`
	Role           string `db:"role" json:"role"`
	Status         int8   `db:"status" json:"-"`
	AccessKey      string `db:"access_key" json:"accessKey"`
	SecretKey      string `db:"secret_key" json:"secretKey"`
	Balance        int64  `db:"balance" json:"balance"`
	InvitationCode string `db:"invitation_code" json:"invitationCode"`
	Model
}

func NewUserAccount(account string) *User {
	user := &User{
		Account: account,
	}
	user.init()
	return user
}

func NewUserEmail(email string) *User {
	user := &User{
		Email: email,
	}
	user.init()
	return user
}

func (u *User) init() {
	u.Gender = 0
	u.Role = "user"
	u.Model.Init()
}
