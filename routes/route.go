package routes

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"tools/app/controller"
	"tools/app/middlware"
	"tools/app/socket"
)

func init() {
	app := g.Server()

	app.BindHandler("/ws", socket.Handler)

	app.Use(middlware.MiddlewareError, middlware.MiddlewareCORS, middlware.MiddlewareThrottle)

	app.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Welcome!")
	})

	app.Group("/api", func(group *ghttp.RouterGroup) {
		group.POST("/login", controller.Login)

		/**************** 角色管理 start ***************/
		// 创建角色
		group.POST("roles", controller.CreateRole)
		// 角色列表
		group.GET("roles", controller.GetRoleList)
		// 角色编辑
		group.PUT("roles/:role_id", controller.UpdateRole)
		// 删除角色
		group.DELETE("roles/:role_id", controller.DeleteRole)
		/**************** 角色管理 end ***************/

		group.Group("/json", func(json *ghttp.RouterGroup) {
			json.POST("/format", controller.FormatJson)
			json.POST("/struct", controller.JsonToStruct)
			json.POST("/sql", controller.SqlToStruct)
		})

		group.GET("/today/one", controller.GetTodayOne)
		group.POST("/language", controller.Language)
		group.GET("/language", controller.GetAllLang)

		group.REST("tools", &controller.ToolsController{})

		group.Middleware(middlware.MiddlewareAuth)

	})
}
