package model

import "github.com/flipped-aurora/gin-vue-admin/server/model/user"

type Request struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type Response struct {
	Username string `json:"username"` // 用户名
}
type LoginResponse struct {
	User      user.FaUser `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
