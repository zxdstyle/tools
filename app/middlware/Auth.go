package middlware

import (
	"github.com/gogf/gf/net/ghttp"
	"tools/app/support/jwt"
)

func MiddlewareAuth(r *ghttp.Request) {
	jwt.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}
