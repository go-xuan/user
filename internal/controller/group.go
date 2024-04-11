package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/frame/ginx"
	"github.com/go-xuan/quanx/net/respx"
	log "github.com/sirupsen/logrus"

	"user/internal/logic"
	"user/internal/model"
)

// 群组分页
func GroupPage(ctx *gin.Context) {
	var err error
	var in model.GroupPage
	if err = ctx.BindJSON(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	var result *respx.PageResponse
	result, err = logic.GroupPage(in)
	respx.BuildResponse(ctx, result, err)
}

// 群组明细
func GroupDetail(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	var result model.GroupDetail
	result, err = logic.GroupDetail(in.Id)
	respx.BuildResponse(ctx, result, err)
}

// 群组删除
func GroupDelete(ctx *gin.Context) {
	var err error
	var in modelx.Id[int64]
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.Exception(ctx, respx.ParamErr, err)
		return
	}
	err = logic.GroupDelete(in.Id)
	respx.BuildResponse(ctx, nil, err)
}

// 群组新增
func GroupSave(ctx *gin.Context) {
	var err error
	var in model.GroupSave
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
		result, err = logic.GroupCreate(&in)
	} else {
		result, err = logic.GroupUpdate(&in)
	}
	respx.BuildResponse(ctx, result, err)
}
