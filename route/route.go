package route

import (
	"IrisAdminApi/controller"
	"IrisAdminApi/middleware"
	"github.com/kataras/iris"
)

/**
路由相关启动
*/
func RouteInit(app *iris.Application) {
	app.Get("/", controller.Test)
	app.Post("/login", controller.Login)

	app.Use(middleware.CheckTokenHandler)
	app.Post("/logout", controller.Logout)
}
