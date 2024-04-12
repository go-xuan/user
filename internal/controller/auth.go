package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/frame/ginx"
	"github.com/go-xuan/quanx/net/respx"
	"github.com/go-xuan/quanx/os/encryptx"

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
	ip := ginx.GetCorrectIP(ctx)
	var result *model.LoginResult
	result, err = logic.UserLogin(ctx.Request.Context(), param, ip)
	if err != nil {
		respx.BuildError(ctx, err)
		return
	}
	ginx.SetCookie(ctx, result.User.Account)
	respx.BuildSuccess(ctx, result)
}

// 用户登出
func UserLogout(ctx *gin.Context) {
	if user := ginx.GetUser(ctx); user != nil {
		ip := ginx.GetCorrectIP(ctx)
		userId, err := logic.UserLogout(ctx.Request.Context(), user, ip)
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
	if user := ginx.GetUser(ctx); user != nil {
		respx.BuildSuccess(ctx, user)
	} else {
		respx.BuildError(ctx, errors.New(""))
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
