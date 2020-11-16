package middlware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"tools/app/support/casbin"
	"tools/app/support/h"
	"tools/app/support/jwt"
)

func MiddlewareAuth(r *ghttp.Request) {

	jwt.Auth.MiddlewareFunc()(r)

	ok, err := casbin.Enforcer.Enforce(1, r.RequestURI, r.Method)
	if err != nil {
		h.Failed(r, err.Error(), 500)
	}

	if !ok {
		h.Failed(r, "您没有权限访问", http.StatusForbidden)
	}

	r.Middleware.Next()
}
