package controller

import (
	"IrisAdminApi/models"
	"IrisAdminApi/validates"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

// 权限控制器

// 获取所有的Permission
func GetPermissionData(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	limit := ctx.URLParamIntDefault("limit", 1)

	Permissions, count := models.GetPermissionList(page, limit)
	data := List{}
	data.Data = Permissions
	data.Page = page
	data.Size = limit
	data.Total = count
	_, _ = ctx.JSON(ReturnData("success", data, "获取权限数据成功"))
}

// 插入权限

func CreatePermission(ctx iris.Context) {
	aul := new(validates.CreatePermissionRequest)
	if err = ctx.ReadForm(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ReturnData("error", nil, err.Error()))
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
}

// 删除权限  软删除
