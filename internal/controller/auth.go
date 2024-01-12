package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx"
	"github.com/go-xuan/quanx/commonx/respx"
	"github.com/go-xuan/quanx/importx/ginx"
	"github.com/go-xuan/quanx/utilx/encryptx"

	"user/internal/logic"
	"user/internal/model"
)

// 用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	var param model.Login
	if err = ctx.BindJSON(&param); err != nil {
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = quanx.GetServer().Host
	}
	var result *model.LoginResult
	result, err = logic.UserLogin(param, ip)
	if err != nil {
		respx.BuildError(ctx, err)
		return
	}
	ginx.SetCookie(ctx, result.User.Account)
	respx.BuildSuccess(ctx, result)
}

// 用户登出
func UserLogout(ctx *gin.Context) {
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = quanx.GetServer().Host
	}
	if value, ok := ctx.Get("user"); ok {
		userId, err := logic.UserLogout(value.(*ginx.User), ip)
		if err != nil {
			respx.BuildError(ctx, err)
			return
		}
		ginx.RemoveCookie(ctx)
		respx.BuildSuccess(ctx, userId)
	}
}

// 验证token
func TokenParse(ctx *gin.Context) {
	if value, ok := ctx.Get("user"); ok {
		respx.BuildResponse(ctx, value, nil)
	}
}

// 加密
func Encrypt(ctx *gin.Context) {
	ciphertext, err := encryptx.RSA().Encrypt(ctx.Query("text"))
	respx.BuildResponse(ctx, ciphertext, err)
}

// 加密
func Decrypt(ctx *gin.Context) {
	plaintext, err := encryptx.RSA().Decrypt(ctx.Query("text"))
	respx.BuildResponse(ctx, plaintext, err)
}
