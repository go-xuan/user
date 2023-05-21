package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/auth"
	"github.com/quanxiaoxuan/go-builder/model/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 群组分页
func GroupPage(context *gin.Context) {
	var err error
	var param params.GroupPage
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result *response.PageResponse
	if result, err = logic.GroupPage(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 群组新增
func GroupAdd(context *gin.Context) {
	var err error
	var param params.GroupCreate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var exist bool
	if exist, err = logic.GroupCodeExist(param.GroupCode); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "此群组编码已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = auth.GetUserId(context)
	}
	var groupId int64
	if groupId, err = logic.GroupAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, groupId)
	}
}

// 群组修改
func GroupUpdate(context *gin.Context) {
	var err error
	var param params.GroupUpdate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if param.UpdateUserId == 0 {
		param.UpdateUserId = auth.GetUserId(context)
	}
	if err = logic.GroupUpdate(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}

// 群组删除
func GroupDelete(context *gin.Context) {
	var err error
	var form struct {
		GroupId int64 `form:"groupId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if err = logic.GroupDelete(form.GroupId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, form.GroupId)
	}
}

// 群组明细
func GroupDetail(context *gin.Context) {
	var err error
	var form struct {
		GroupId int64 `form:"groupId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result results.GroupDetail
	if result, err = logic.GroupDetail(form.GroupId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}
