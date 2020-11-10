package routes

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"tools/app/controller"
)

func init() {
	app := g.Server()

	app.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Welcome!")
	})

	app.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/json", func(json *ghttp.RouterGroup) {
			json.POST("/format", controller.FormatJson)
			json.POST("/struct", controller.JsonToStruct)
			json.POST("/sql", controller.SqlToStruct)
		})
	})
}
