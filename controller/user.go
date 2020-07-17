package controller

import (
	"IrisAdminApi/models"
	"IrisAdminApi/validates"
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

type Data map[interface{}]interface{}

func Test(ctx iris.Context) {
	msg := "test"
	_, _ = ctx.JSON(ReturnData("Success", Data{}, msg))
}

var (
	err error
)

func Login(ctx iris.Context) {
	aul := new(validates.LoginRequest)
	if err = ctx.ReadForm(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		fmt.Println(err)
		_, _ = ctx.JSON(ReturnData("success", nil, err.Error()))
		return
	}

	// 进行数据验证
	if err = validates.Validate.Struct(*aul); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validates.Trans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusUnprocessableEntity)
				_, _ = ctx.JSON(ReturnData("error", nil, e))
				return
			}
		}
	}

	// 查询用户信息
	user := new(models.User)
	user.Name = aul.Username
	user.GetUserByUsername()
	// 检测密码正确与否
	// 生成对应的token返回
}
