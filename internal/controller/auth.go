package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/quanx/common/respx"
	"github.com/quanxiaoxuan/quanx/engine"
	"github.com/quanxiaoxuan/quanx/public/authx"
	"github.com/quanxiaoxuan/quanx/utils/encryptx"
	log "github.com/sirupsen/logrus"

	"quan-user/internal/logic"
	"quan-user/model"
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
		ip = engine.GetEngine().Config.Server.Host
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
		ip = engine.GetEngine().Config.Server.Host
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
		respx.BuildResponse(ctx, value, nil)
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
	ctx.SetCookie(authx.CookieKey, value, maxAge, "", "", false, true)
	return nil
}

// 加密
func Encrypt(ctx *gin.Context) {
	var text = ctx.Query("text")
	ciphertext, err := encryptx.RSA().Encrypt(text)
	respx.BuildResponse(ctx, ciphertext, err)
}

// 加密
func Decrypt(ctx *gin.Context) {
	var text = ctx.Query("text")
	plaintext, err := encryptx.RSA().Decrypt(text)
	respx.BuildResponse(ctx, plaintext, err)
}
