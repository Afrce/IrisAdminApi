package controller

import (
	"IrisAdminApi/config"
	"IrisAdminApi/models"
	"IrisAdminApi/validates"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type Data map[interface{}]interface{}

type jwtCustomClaims struct {
	jwt.StandardClaims
	Uid uint
	Exp int64
}

func Test(ctx iris.Context) {
	msg := "test"
	_, _ = ctx.JSON(ReturnData("Success", Data{}, msg))
}

var (
	err error
)

func Login(ctx iris.Context) {

	var (
		pwd     string
		expTime time.Time
	)

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
	pwd = user.Password

	if err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(aul.Password)); err != nil {
		_, _ = ctx.JSON(ReturnData("error", nil, "账号密码错误，请查证后再试"))
		return
	}
	// 生成对应的token返回

	if aul.Remember == 1 {
		expTime = time.Now().Local().Add(time.Hour * 24 * 7)
	} else {
		expTime = time.Now().Local().Add(time.Hour * 12)
	}

	stdClaims := jwt.StandardClaims{
		ExpiresAt: expTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        string(user.ID),
		Issuer:    "GO",
	}
	claims := &jwtCustomClaims{
		StandardClaims: stdClaims,
		Exp:            int64(expTime.Unix()),
		Uid:            user.ID,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(config.Config.Jwt))
	if err != nil {
		ctx.Application().Logger().Error("生成token失败：" + err.Error())
	}

	// 将token写入 数据库

	accessToken := new(models.AccessToken)
	accessToken.UserID = user.ID
	accessToken.Token = token
	accessToken.ExpiresAt = expTime

	models.CreateToken(*accessToken)

	data := make(map[string]string)

	data["token"] = token
	_, _ = ctx.JSON(ReturnData("success", data, "登录成功"))
	return
}

func Logout(ctx iris.Context) {
	// 删除token
	token := ctx.Values().GetString("token")
	Uid, _ := ctx.Values().GetUint("user_id")
	fmt.Println(Uid)
	models.DeleteToken(token, Uid)
	_, _ = ctx.JSON(ReturnData("success", nil, "退出登录成功"))
	return
}
