package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/app/ginx"
	"github.com/go-xuan/quanx/app/modelx"
	"github.com/go-xuan/quanx/net/respx"

	"user/internal/logic"
	"user/internal/model"
)

// UserPage 用户分页
func UserPage(ctx *gin.Context) {
	var err error
	var in model.UserPage
	if err = ctx.BindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.UserPage(in))
}

// UserList 用户列表
func UserList(ctx *gin.Context) {
	respx.Ctx(ctx).Response(logic.UserList())
}

// UserSave 用户新增
func UserSave(ctx *gin.Context) {
	var err error
	var in model.UserSave
	if err = ctx.BindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	if in.Id != 0 {
		respx.Ctx(ctx).Response(nil, logic.UserUpdate(&in))
	} else {
		respx.Ctx(ctx).Response(logic.UserCreate(&in))
	}
}

// UserDelete 用户删除
func UserDelete(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(nil, logic.UserDelete(in.Id))
}

// UserDetail 用户明细
func UserDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.UserDetail(in.Id))
}
