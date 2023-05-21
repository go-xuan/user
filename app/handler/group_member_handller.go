package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/auth"
	"github.com/quanxiaoxuan/go-builder/model/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/params"
)

// 群组成员新增
func GroupMemberAdd(context *gin.Context) {
	var err error
	var param params.GroupMemberAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var userIds []int64
	for _, item := range param.GroupMemberList {
		userIds = append(userIds, item.UserId)
	}
	var exist bool
	if exist, err = logic.GroupMemberExist(param.GroupId, userIds); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "新增群组成员已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = auth.GetUserId(context)
	}
	if err = logic.GroupMemberAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}
