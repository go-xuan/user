package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/net/respx"

	"user/internal/logic"
	"user/internal/model"
)

// GroupPage 群组分页
func GroupPage(ctx *gin.Context) {
	var in model.GroupPage
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}

	res, err := logic.GroupPage(in)
	respx.Response(ctx, res, err)
}

// GroupDetail 群组明细
func GroupDetail(ctx *gin.Context) {
	var in modelx.Id[int64]
	if err := ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	res, err := logic.GroupDetail(in.Id)
	respx.Response(ctx, res, err)
}

// GroupDelete 群组删除
func GroupDelete(ctx *gin.Context) {
	var in modelx.Id[int64]
	if err := ctx.ShouldBindQuery(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	err := logic.GroupDelete(in.Id)
	respx.Response(ctx, nil, err)
}

// GroupSave 群组新增
func GroupSave(ctx *gin.Context) {
	var in model.GroupSave
	if err := ctx.ShouldBindJSON(&in); err != nil {
		respx.ParamError(ctx, err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	if in.Id == 0 {
		result, err := logic.GroupCreate(&in)
		respx.Response(ctx, result, err)
	} else {
		result, err := logic.GroupUpdate(&in)
		respx.Response(ctx, result, err)
	}
}
