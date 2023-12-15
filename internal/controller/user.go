package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xuan/quanx/authx"
	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/common/respx"
	log "github.com/sirupsen/logrus"

	"user/internal/logic"
	"user/internal/model"
)

// 用户分页
func UserPage(ctx *gin.Context) {
	var err error
	var in model.UserPage
	if err = ctx.BindJSON(&in); err != nil {
		log.Error("参数错误：", err)
		respx.BuildException(ctx, respx.ParamErr, err)
		return
	}
	var result *respx.PageResponse
	result, err = logic.UserPage(in)
	respx.BuildResponse(ctx, result, err)
}

// 用户列表
func UserList(ctx *gin.Context) {
	var err error
	var result []*model.User
	result, err = logic.UserList()
	respx.BuildResponse(ctx, result, err)
}

// 用户新增
func UserSave(ctx *gin.Context) {
	var err error
	var in model.UserSave
	if err = ctx.BindJSON(&in); err != nil {
		log.Error("参数错误：", err)
		respx.BuildException(ctx, respx.ParamErr, err)
		return
	}
	if in.CurrUserId == 0 {
		in.CurrUserId = authx.GetUserId(ctx)
	}
	if in.Id != 0 {
		err = logic.UserUpdate(&in)
		respx.BuildResponse(ctx, nil, err)
	} else {
		var result int64
		result, err = logic.UserCreate(&in)
		respx.BuildResponse(ctx, result, err)
	}
}

// 用户删除
func UserDelete(ctx *gin.Context) {
	var err error
	var in modelx.IdInt64
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.BuildException(ctx, respx.ParamErr, err)
		return
	}
	err = logic.UserDelete(in.Id)
	respx.BuildResponse(ctx, nil, err)
}

// 用户明细
func UserDetail(ctx *gin.Context) {
	var err error
	var in modelx.IdInt64
	if err = ctx.ShouldBindQuery(&in); err != nil {
		log.Error("参数错误：", err)
		respx.BuildException(ctx, respx.ParamErr, err)
		return
	}
	var result *model.User
	result, err = logic.UserDetail(in.Id)
	respx.BuildResponse(ctx, result, err)
}
