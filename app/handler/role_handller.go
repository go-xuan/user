package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/authx"
	"github.com/quanxiaoxuan/go-builder/paramx/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 用户分页查询
func RolePage(context *gin.Context) {
	var err error
	var param params.RolePage
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result *response.PageResponse
	if result, err = logic.RolePage(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 角色列表
func RoleList(context *gin.Context) {
	var err error
	var result results.RoleSimpleList
	if result, err = logic.RoleList(); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}

// 角色新增
func RoleAdd(context *gin.Context) {
	var err error
	var param params.RoleCreate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var exist bool
	if exist, err = logic.RoleCodeExist(param.RoleCode); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "此角色编码已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context)
	}
	var roleId int64
	if roleId, err = logic.RoleAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, roleId)
	}
}

// 角色修改
func RoleUpdate(context *gin.Context) {
	var err error
	var param params.RoleUpdate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if param.UpdateUserId == 0 {
		param.UpdateUserId = authx.GetUserId(context)
	}
	if err = logic.RoleUpdate(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}

// 角色删除
func RoleDelete(context *gin.Context) {
	var err error
	var form struct {
		RoleId int64 `form:"roleId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	if err = logic.RoleDelete(form.RoleId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, form.RoleId)
	}
}

// 角色详情
func RoleDetail(context *gin.Context) {
	var err error
	var form struct {
		RoleId int64 `form:"roleId" binding:"required"`
	}
	if err = context.ShouldBindQuery(&form); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var result results.RoleDetail
	if result, err = logic.RoleDetail(form.RoleId); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, result)
	}
}
