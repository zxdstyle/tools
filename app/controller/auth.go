package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"tools/app/models"
	"tools/app/service/auth"
	"tools/app/support/h"
)

func Login(r *ghttp.Request) {
	if err := auth.ValidateLoginRequest(r); err != "" {
		h.Failed(r, err, 422)
	}

	var user models.User
	r.GetStruct(&user)

	token, err := user.GetToken()
	if err != nil {
		h.Failed(r, err.Error())
	}

	h.Success(r, g.Map{
		"token":    token,
		"username": user.Username,
	})
}
