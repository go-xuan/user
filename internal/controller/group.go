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
	var err error
	var in model.GroupPage
	if err = ctx.ShouldBindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.GroupPage(in))
}

// GroupDetail 群组明细
func GroupDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(logic.GroupDetail(in.Id))
}

// GroupDelete 群组删除
func GroupDelete(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	respx.Ctx(ctx).Response(nil, logic.GroupDelete(in.Id))
}

// GroupSave 群组新增
func GroupSave(ctx *gin.Context) {
	var err error
	var in model.GroupSave
	if err = ctx.ShouldBindJSON(&in); err != nil {
		respx.Ctx(ctx).ParamError(err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = ginx.GetSessionUser(ctx).UserId()
	}
	var result int64
	if in.Id == 0 {
		result, err = logic.GroupCreate(&in)
	} else {
		result, err = logic.GroupUpdate(&in)
	}
	respx.Ctx(ctx).Response(result, err)
}
