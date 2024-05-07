package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"
	log "github.com/sirupsen/logrus"

	"user/internal/logic"
	"user/internal/model"
)

// 用户分页查询
func RolePage(ctx *gin.Context) {
	var err error
	var in model.RolePage
	if err = ctx.BindJSON(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	var result *respx.PageResponse
	result, err = logic.RolePage(in)
	respx.BuildResponse(ctx, result, err)

}

// 角色列表
func RoleList(ctx *gin.Context) {
	result, err := logic.RoleList()
	respx.BuildResponse(ctx, result, err)
}

// 角色新增
func RoleSave(ctx *gin.Context) {
	var err error
	var in model.RoleSave
	if err = ctx.BindJSON(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetUserId(ctx)
	}
	var result int64
	if in.Id == 0 {
		result, err = logic.RoleCreate(&in)
		respx.BuildResponse(ctx, result, err)
	} else {
		err = logic.RoleUpdate(&in)
		respx.BuildResponse(ctx, nil, err)
	}
}

// 角色删除
func RoleDelete(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	err = logic.RoleDelete(in.Id)
	respx.BuildResponse(ctx, nil, err)
}

// 角色详情
func RoleDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	var result *model.RoleDetail
	result, err = logic.RoleDetail(in.Id)
	respx.BuildResponse(ctx, result, err)
}
