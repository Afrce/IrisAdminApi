package main

import (
	"IrisAdminApi/config"
	"IrisAdminApi/libs"
	"IrisAdminApi/route"
	"fmt"
	"github.com/kataras/iris"
	"time"
)

func main() {
	f := libs.NewLogApp()
	defer f.Close()

	app := iris.New()
	app.Logger().SetOutput(f)

	route.RouteInit(app)

	if config.Config.HTTPS { // 启动HTTPS服务
		host := fmt.Sprintf("%s:%d", config.Config.Host, 443)
		if err := app.Run(iris.TLS(host, config.Config.Certpath, config.Config.Certkey)); err != nil {
			fmt.Println(err)
		}
	} else { // 启动HTTP服务
		if err := app.Run(
			iris.Addr(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)),
			iris.WithoutServerError(iris.ErrServerClosed),
			iris.WithOptimizations,
			iris.WithTimeFormat(time.RFC3339),
		); err != nil {
			fmt.Println(err)
		}
	}
}
