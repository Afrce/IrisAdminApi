package controller

import "github.com/kataras/iris"

type Data map[interface{}]interface{}

func Test(ctx iris.Context) {
	msg := "test"
	_, _ = ctx.JSON(ReturnData("Success", Data{}, msg))
}
