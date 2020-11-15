package routes

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"tools/app/controller"
	"tools/app/middlware"
)

func init() {
	app := g.Server()

	app.Use(middlware.MiddlewareError, middlware.MiddlewareCORS)

	app.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Welcome!")
	})

	app.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/json", func(json *ghttp.RouterGroup) {
			json.POST("/format", controller.FormatJson)
			json.POST("/struct", controller.JsonToStruct)
			json.POST("/sql", controller.SqlToStruct)
		})

		group.GET("/today/one", controller.GetTodayOne)
		group.POST("/language", controller.Language)
		group.GET("/language", controller.GetAllLang)

		group.REST("tools", &controller.ToolsController{})
	})
}
