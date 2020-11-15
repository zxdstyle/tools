package middlware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
	"tools/app/support/h"
)

// 全局中间件

// 错误捕捉中间件
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

// 跨域中间件
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 限流中间件
func MiddlewareThrottle(r *ghttp.Request) {
	bucket := ratelimit.NewBucket(1*time.Second, 100)

	r.Response.Header().Set("X-Ratelimit-Limit", "100")
	r.Response.Header().Set("X-Ratelimit-Remaining", gconv.String(bucket.Rate()))

	if bucket.TakeAvailable(1) < 1 {
		h.Failed(r, "请求频繁，请稍后再试", 429)
	}
	r.Middleware.Next()
}
