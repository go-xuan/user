package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx"
	"github.com/go-xuan/quanx/public/authx"
	"github.com/go-xuan/quanx/public/respx"
	"github.com/go-xuan/quanx/utils/encryptx"
	log "github.com/sirupsen/logrus"

	"user/internal/logic"
	"user/internal/model"
)

// 用户登录
func UserLogin(ctx *gin.Context) {
	var err error
	var param model.Login
	if err = ctx.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildException(ctx, respx.ParamErr, err)
		return
	}
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = quanx.GetEngine().Config.Server.Host
	}
	var result *model.LoginResult
	result, err = logic.UserLogin(param, ip)
	if err != nil {
		respx.BuildError(ctx, err)
		return
	}
	if err = SetCookie(ctx, result.User.Account); err != nil {
		respx.BuildError(ctx, err)
		return
	}
	respx.BuildSuccess(ctx, result)
}

// 用户登出
func UserLogout(ctx *gin.Context) {
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = quanx.GetEngine().Config.Server.Host
	}
	if value, ok := ctx.Get(authx.TokenUser); ok {
		userId, err := logic.UserLogout(value.(*authx.User), ip)
		if err != nil {
			respx.BuildError(ctx, err)
			return
		}
		if err = SetCookie(ctx, ""); err != nil {
			respx.BuildError(ctx, err)
			return
		}
		respx.BuildSuccess(ctx, userId)
	}
}

// 验证token
func TokenParse(ctx *gin.Context) {
	if value, ok := ctx.Get(authx.TokenUser); ok {
		respx.BuildNormal(ctx, value, nil)
	}
}

func SetCookie(ctx *gin.Context, value string) error {
	var maxAge = 3600
	if value == "" {
		maxAge = -1
	} else {
		bytes, err := encryptx.RSA().Encrypt(value)
		if err != nil {
			return err
		}
		value = bytes
	}
	ctx.SetCookie(authx.CookieAuth, value, maxAge, "", "", false, true)
	return nil
}

// 加密
func Encrypt(ctx *gin.Context) {
	var text = ctx.Query("text")
	ciphertext, err := encryptx.RSA().Encrypt(text)
	respx.BuildNormal(ctx, ciphertext, err)
}

// 加密
func Decrypt(ctx *gin.Context) {
	var text = ctx.Query("text")
	plaintext, err := encryptx.RSA().Decrypt(text)
	respx.BuildNormal(ctx, plaintext, err)
}
