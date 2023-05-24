package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/authx"
	"github.com/quanxiaoxuan/go-builder/paramx/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/params"
)

// 角色成员新增
func RoleMemberAdd(context *gin.Context) {
	var err error
	var param params.RoleMemberAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var userIds []int64
	for _, item := range param.RoleMemberList {
		userIds = append(userIds, item.UserId)
	}
	var exist bool
	if exist, err = logic.RoleMemberExist(param.RoleId, userIds); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "新增角色成员已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context)
	}
	if err = logic.RoleMemberAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}
