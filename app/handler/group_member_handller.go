package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/quanx/common/authx"
	"github.com/quanxiaoxuan/quanx/common/respx"
	log "github.com/sirupsen/logrus"
	"quan-admin/common"

	"quan-admin/app/logic"
	"quan-admin/model/params"
)

// 群组成员新增
func GroupMemberAdd(context *gin.Context) {
	var err error
	var param params.GroupMemberAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var userIds []int64
	for _, item := range param.GroupMemberList {
		userIds = append(userIds, item.UserId)
	}
	var exist bool
	if exist, err = logic.GroupMemberExist(param.GroupId, userIds); err != nil {
		respx.BuildExceptionResponse(context, respx.UniqueErr, err)
		return
	}
	if exist {
		respx.BuildErrorResponse(context, "新增群组成员已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context, common.SecretKey)
	}
	if err = logic.GroupMemberAdd(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, nil)
	}
}
