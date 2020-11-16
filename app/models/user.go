package models

import (
	"github.com/gogf/gf/frame/g"
	"tools/app/support/jwt"
)

type User struct {
	Id       uint64
	Username string
	Password string
}

func (user *User) GetToken() (string, error) {
	return jwt.Auth.GenerateToken(g.Map{
		"username": user.Username,
		"id":       user.Id,
	})
}
