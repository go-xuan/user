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

// 群组角色新增
func GroupRoleAdd(context *gin.Context) {
	var err error
	var param params.GroupRoleAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		respx.BuildExceptionResponse(context, respx.ParamErr, err)
		return
	}
	var roleIds []int64
	for _, item := range param.GroupRoleList {
		roleIds = append(roleIds, item.RoleId)
	}
	var exist bool
	if exist, err = logic.GroupRoleExist(param.GroupId, roleIds); err != nil {
		respx.BuildExceptionResponse(context, respx.UniqueErr, err)
		return
	}
	if exist {
		respx.BuildErrorResponse(context, "新增群组角色已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = authx.GetUserId(context, common.SecretKey)
	}
	if err = logic.GroupRoleAdd(param); err != nil {
		respx.BuildErrorResponse(context, err.Error())
	} else {
		respx.BuildSuccessResponse(context, nil)
	}
}
