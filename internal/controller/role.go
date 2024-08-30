package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/app/ginx"
	"github.com/go-xuan/quanx/app/modelx"
	"github.com/go-xuan/quanx/net/respx"
	"user/internal/logic"
	"user/internal/model"
)

// RolePage 用户分页查询
func RolePage(ctx *gin.Context) {
	var err error
	var in model.RolePage
	if err = ctx.BindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.RolePage(in))

}

// RoleList 角色列表
func RoleList(ctx *gin.Context) {
	respx.Ctx(ctx).Response(logic.RoleList())
}

// RoleSave 角色新增
func RoleSave(ctx *gin.Context) {
	var err error
	var in model.RoleSave
	if err = ctx.BindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	if in.Id == 0 {
		respx.Ctx(ctx).Response(logic.RoleCreate(&in))
	} else {
		err = logic.RoleUpdate(&in)
		respx.Ctx(ctx).Response(nil, logic.RoleUpdate(&in))
	}
}

// RoleDelete 角色删除
func RoleDelete(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(nil, logic.RoleDelete(in.Id))
}

// RoleDetail 角色详情
func RoleDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.RoleDetail(in.Id))
}
