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

// 用户分页
func UserPage(context *gin.Context) {
	var err error
	var param params.UserPage
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result *response.PageResponse
	if result, err = logic.UserPage(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 用户列表
func UserList(context *gin.Context) {
	var err error
	var result results.UserSimpleList
	if result, err = logic.UserList(); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 用户新增
func UserAdd(context *gin.Context) {
	var err error
	var param params.UserCreate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var exist bool
	if exist, err = logic.UserPhoneExist(param.Phone); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "此手机号已注册")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = auth.GetUserId(context)
	}
	var userId int64
	if userId, err = logic.UserAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, userId)
	}
}

// 用户修改
func UserUpdate(context *gin.Context) {
	var err error
	var param params.UserUpdate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if param.UpdateUserId == 0 {
		param.UpdateUserId = auth.GetUserId(context)
	}
	if err = logic.UserUpdate(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}

// 用户删除
func UserDelete(context *gin.Context) {
	var err error
	var form struct {
		UserId int64 `form:"userId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if err = logic.UserDelete(form.UserId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, form.UserId)
	}
}

// 用户明细
func UserDetail(context *gin.Context) {
	var err error
	var form struct {
		UserId int64 `form:"userId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result *results.UserDetail
	if result, err = logic.UserDetail(form.UserId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}
