package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/authx"
	"github.com/quanxiaoxuan/go-builder/paramx/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/conf"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 用户登录
func UserLogin(context *gin.Context) {
	var err error
	var param params.UserLogin
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	ip := context.ClientIP()
	if ip == "::1" {
		ip = conf.Config.System.Host
	}
	var result *results.LoginResult
	if result, err = logic.UserLogin(param, ip); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 验证token
func TokenParse(context *gin.Context) {
	var err error
	var result *authx.Param
	token := context.Request.Header.Get("Authorization")
	if result, err = logic.TokenParse(token); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 用户登出
func UserLogout(context *gin.Context) {
	ip := context.ClientIP()
	if ip == "::1" {
		ip = conf.Config.System.Host
	}
	var err error
	var userId int64
	token := context.Request.Header.Get("Authorization")
	if userId, err = logic.UserLogout(token, ip); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, userId)
	}
}
