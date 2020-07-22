package middleware

import (
	"IrisAdminApi/controller"
	"IrisAdminApi/models"
	"github.com/kataras/iris"
	"strings"
)

// 权限验证

func CheckUserPermission(ctx iris.Context) {
	routeName := ctx.GetCurrentRoute().Name()

	if strings.Contains(routeName, ".public") { // 包含public 表示公共权限
		ctx.Next()
		return
	}

	// 验证permission
	id, ok := models.SelectPermission(routeName)
	if !ok {
		ctx.StatusCode(iris.StatusForbidden)
		_, _ = ctx.JSON(controller.ReturnData("error", nil, "无对应权限查看"))
		return
	}

	userId, _ := ctx.Values().GetUint("user_id")
	userHasPermission := models.CheckUserHasPermission(id, userId)
	if userHasPermission {
		ctx.Next()
		return
	} else {
		// 查询用户角色是否有权限
		roles := models.SelectUserRoles(userId)

		if len(roles) == 0 {
			ctx.StatusCode(iris.StatusForbidden)
			_, _ = ctx.JSON(controller.ReturnData("error", nil, "无对应权限查看"))
			return
		}

		// 查询角色是否有权限
		if ok := models.CheckRoleHasPermission(roles, id); !ok {
			ctx.StatusCode(iris.StatusForbidden)
			_, _ = ctx.JSON(controller.ReturnData("error", nil, "无对应权限查看"))
			return
		} else {
			ctx.Next()
		}
	}
}
