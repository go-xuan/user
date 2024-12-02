package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"
	"user/internal/logic"
	"user/internal/model"
)

// RolePage 用户分页查询
func RolePage(ctx *gin.Context) {
	var in model.RolePage
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	res, err := logic.RolePage(in)
	respx.Response(ctx, res, err)

}

// RoleList 角色列表
func RoleList(ctx *gin.Context) {
	res, err := logic.RoleList()
	respx.Response(ctx, res, err)
}

// RoleSave 角色新增
func RoleSave(ctx *gin.Context) {
	var in model.RoleSave
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	if in.Id == 0 {
		res, err := logic.RoleCreate(&in)
		respx.Response(ctx, res, err)
	} else {
		err := logic.RoleUpdate(&in)
		respx.Response(ctx, nil, err)
	}
}

// RoleDelete 角色删除
func RoleDelete(ctx *gin.Context) {
	var in modelx.Id[int64]
	if err := ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	err := logic.RoleDelete(in.Id)
	respx.Response(ctx, nil, err)
}

// RoleDetail 角色详情
func RoleDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	res, err := logic.RoleDetail(in.Id)
	respx.Response(ctx, res, err)
}
