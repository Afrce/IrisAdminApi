package middleware

import (
	"IrisAdminApi/config"
	"IrisAdminApi/controller"
	"IrisAdminApi/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"time"
)

// 检测用户是否登录的中间件

func CheckTokenHandler(ctx iris.Context) {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(controller.ReturnData("error", nil, "请先登录"))
		return
	}

	token = token[7:]
	// 解析JWT 的token
	type jwtCustomClaims struct {
		jwt.StandardClaims
		Uid uint
		Exp int64
	}

	uClaims := jwtCustomClaims{}

	_, err := jwt.ParseWithClaims(token, &uClaims, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		}
		return []byte(config.Config.Jwt), nil
	})
	if err != nil {
		ctx.Application().Logger().Error("解析token错误", err)
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(controller.ReturnData("error", nil, "登录已失效，请重新登录"))
		return
	}

	exp := uClaims.Exp
	uid := uClaims.Uid

	// 比较有效时间
	nowTime := time.Now().Unix()

	if nowTime > exp {
		// 过期
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(controller.ReturnData("error", nil, "登录已失效，请重新登录"))
		return
	}

	// 验证数据库内是否有token
	if ok := models.SelectToken(token, uid); !ok {
		// 无效
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(controller.ReturnData("error", nil, "登录已失效，请重新登录"))
		return
	}

	ctx.Values().Set("user_id", uid)
	ctx.Values().Set("token", token)
	ctx.Next()

}
