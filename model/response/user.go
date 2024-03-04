package response

import "github.com/axliupore/axapi/axapi-backend/model"

type UserLoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
