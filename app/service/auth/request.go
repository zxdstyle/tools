package auth

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type LoginRequest struct {
	Username string `v:"username@required"`
	Password string `v:"password@required|password2"`
}

func ValidateLoginRequest(r *ghttp.Request) string {
	var login LoginRequest
	if err := r.GetStruct(&login); err != nil {
		return err.Error()
	}

	if err := gvalid.CheckStruct(login, nil); err != nil {
		return err.FirstString()
	}

	return ""
}
