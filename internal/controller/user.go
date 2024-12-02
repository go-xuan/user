package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"

	"user/internal/logic"
	"user/internal/model"
)

// UserPage 用户分页
func UserPage(ctx *gin.Context) {
	var in model.UserPage
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	res, err := logic.UserPage(in)
	respx.Response(ctx, res, err)
}

// UserList 用户列表
func UserList(ctx *gin.Context) {
	res, err := logic.UserList()
	respx.Response(ctx, res, err)
}

// UserSave 用户新增
func UserSave(ctx *gin.Context) {
	var in model.UserSave
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	if in.Id != 0 {
		err := logic.UserUpdate(&in)
		respx.Response(ctx, nil, err)
	} else {
		res, err := logic.UserCreate(&in)
		respx.Response(ctx, res, err)
	}
}

// UserDelete 用户删除
func UserDelete(ctx *gin.Context) {
	var in modelx.Id[int64]
	if err := ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	if err := logic.UserDelete(in.Id); err != nil {
		respx.Error(ctx, err.Error())
	} else {
		respx.Success(ctx, nil)
	}
}

// UserDetail 用户明细
func UserDetail(ctx *gin.Context) {
	var in modelx.Id[int64]
	if err := ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	res, err := logic.UserDetail(in.Id)
	respx.Response(ctx, res, err)
}
