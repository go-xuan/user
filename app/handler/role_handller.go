package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/quanx/common/authx"
	"github.com/quanxiaoxuan/quanx/common/respx"
	log "github.com/sirupsen/logrus"
	"quan-admin/common"

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
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var result *respx.PageResponse
	if result, err = logic.RolePage(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, result)
	}
}

// 角色列表
func RoleList(context *gin.Context) {
	var err error
	var result results.RoleSimpleList
	if result, err = logic.RoleList(); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, result)
	}
}

// 角色新增
func RoleAdd(context *gin.Context) {
	var err error
	var param params.RoleCreate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var exist bool
	if exist, err = logic.RoleCodeExist(param.RoleCode); err != nil {
		respx.BuildExceptionResponse(context, respx.UniqueErr, err)
		return
	}
	if exist {
		respx.BuildErrorResponse(context, "此角色编码已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context, common.SecretKey)
	}
	var roleId int64
	if roleId, err = logic.RoleAdd(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, roleId)
	}
}

// 角色修改
func RoleUpdate(context *gin.Context) {
	var err error
	var param params.RoleUpdate
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	if param.UpdateUserId == 0 {
		param.UpdateUserId = authx.GetUserId(context, common.SecretKey)
	}
	if err = logic.RoleUpdate(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, nil)
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
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	if err = logic.RoleDelete(form.RoleId); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, form.RoleId)
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
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var result results.RoleDetail
	if result, err = logic.RoleDetail(form.RoleId); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, result)
	}
}
