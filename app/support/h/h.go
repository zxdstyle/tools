package h

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 定义一些全局辅助方法

func Failed(r *ghttp.Request, message string, code ...interface{}) {
	status := 400
	if len(code) > 0 {
		status = code[0].(int)
	}

	r.Response.WriteHeader(status)
	r.Response.WriteJsonExit(g.Map{
		"code":    status,
		"message": message,
	})
}

func Success(r *ghttp.Request, data ...interface{}) {
	if len(data) > 0 {
		r.Response.WriteJsonExit(data[0])
	}
	r.Response.WriteJsonExit(g.Map{
		"code":    200,
		"message": "success",
	})
}
