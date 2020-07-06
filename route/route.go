package route

import (
	"IrisAdminApi/controller"
	"github.com/kataras/iris"
)

/**
路由相关启动
*/
func RouteInit(app *iris.Application) {
	app.Get("/", controller.Test)
}
