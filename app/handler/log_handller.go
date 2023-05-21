package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/model/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/entity"
)

// 日志删除
func LogAdd(context *gin.Context) {
	var err error
	var sysLog entity.SysLog
	if err = context.BindJSON(&sysLog); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if err = logic.LogAdd(sysLog); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}

// 日志删除
func LogDelete(context *gin.Context) {
	var err error
	var param []string
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if err = logic.LogDelete(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}
