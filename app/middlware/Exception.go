package middlware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

func MiddlewareError(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.Status == http.StatusNotFound {
		r.Response.ClearBuffer()
		r.Response.WriteHeader(404)
		r.Response.WriteJson(g.Map{
			"code":    404,
			"message": "未找到该路由",
		})
	}
}
