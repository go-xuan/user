package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/utils/encryptx"

	"user/internal/logic"
	"user/internal/model"
)

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	var param model.Login
	if err = ctx.ShouldBindJSON(&param); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	ip := ginx.GetCorrectIP(ctx)
	var result *model.LoginResult
	if result, err = logic.UserLogin(ctx, param, ip); err != nil {
		respx.Error(ctx, err.Error())
		return
	}
	ginx.SetAuthCookie(ctx, result.User.Phone)
	respx.Success(ctx, result)
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	if user := ginx.GetSessionUser(ctx); user != nil {
		ip := ginx.GetCorrectIP(ctx)
		if err := logic.UserLogout(ctx, user, ip); err != nil {
			respx.Error(ctx, err.Error())
			return
		}
		ginx.RemoveAuthCookie(ctx)
		respx.Success(ctx, user.UserId())
	}
}

// CheckLogin 校验登录
func CheckLogin(ctx *gin.Context) {
	if user := ginx.GetSessionUser(ctx); user != nil {
		respx.Success(ctx, user)
	} else {
		respx.Error(ctx, "token is invalid")
	}
}

// Encrypt 加密
func Encrypt(ctx *gin.Context) {
	res, err := encryptx.RSA().Encrypt(ctx.Query("text"))
	respx.Response(ctx, res, err)
}

// Decrypt 加密
func Decrypt(ctx *gin.Context) {
	res, err := encryptx.RSA().Decrypt(ctx.Query("text"))
	respx.Response(ctx, res, err)
}
