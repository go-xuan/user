package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/os/errorx"
	"github.com/go-xuan/quanx/utils/encryptx"

	"user/internal/logic"
	"user/internal/model"
)

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	var param model.Login
	if err = ctx.ShouldBindJSON(&param); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	ip := ginx.GetCorrectIP(ctx)
	var result *model.LoginResult
	if result, err = logic.UserLogin(ctx, param, ip); err != nil {
		respx.Ctx(ctx).Failed(err)
		return
	}
	ginx.SetAuthCookie(ctx, result.User.Phone)
	respx.Ctx(ctx).Success(result)
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	if user := ginx.GetSessionUser(ctx); user != nil {
		ip := ginx.GetCorrectIP(ctx)
		if err := logic.UserLogout(ctx, user, ip); err != nil {
			respx.Ctx(ctx).Failed(err)
			return
		}
		ginx.RemoveAuthCookie(ctx)
		respx.Ctx(ctx).Success(user.UserId())
	}
}

// CheckLogin 校验登录
func CheckLogin(ctx *gin.Context) {
	if user := ginx.GetSessionUser(ctx); user != nil {
		respx.Ctx(ctx).Success(user)
	} else {
		err := errorx.New("token is invalid")
		respx.Ctx(ctx).Failed(err)
	}
}

// Encrypt 加密
func Encrypt(ctx *gin.Context) {
	respx.Ctx(ctx).Response(encryptx.RSA().Encrypt(ctx.Query("text")))
}

// Decrypt 加密
func Decrypt(ctx *gin.Context) {
	respx.Ctx(ctx).Response(encryptx.RSA().Decrypt(ctx.Query("text")))
}
