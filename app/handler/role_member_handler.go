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

// 角色成员新增
func RoleMemberAdd(context *gin.Context) {
	var err error
	var param params.RoleMemberAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var userIds []int64
	for _, item := range param.RoleMemberList {
		userIds = append(userIds, item.UserId)
	}
	var exist bool
	if exist, err = logic.RoleMemberExist(param.RoleId, userIds); err != nil {
		respx.BuildExceptionResponse(context, respx.UniqueErr, err)
		return
	}
	if exist {
		respx.BuildErrorResponse(context, "新增角色成员已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context, common.SecretKey)
	}
	if err = logic.RoleMemberAdd(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, nil)
	}
}
