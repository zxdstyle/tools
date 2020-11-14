package controller

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func GetTodayOne(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"one": "你是否知道自己做事的理念和原则是什么？如果不清楚，那么你很有可能随波逐流，接受最时髦的事物，并且不知道你的言行正在受其影响、被其浸染。",
	})
}
