package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/auth"
	"github.com/quanxiaoxuan/go-builder/model/response"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/logic"
	"quan-admin/model/params"
)

// 群组角色新增
func GroupRoleAdd(context *gin.Context) {
	var err error
	var param params.GroupRoleAdd
	if err = context.BindJSON(&param); err != nil {
		log.Error("参数错误：", err)
		response.BuildExceptionResponse(context, response.ParamErr, err)
		return
	}
	var roleIds []int64
	for _, item := range param.GroupRoleList {
		roleIds = append(roleIds, item.RoleId)
	}
	var exist bool
	if exist, err = logic.GroupRoleExist(param.GroupId, roleIds); err != nil {
		response.BuildExceptionResponse(context, response.DuplicateErr, err)
		return
	}
	if exist {
		response.BuildErrorResponse(context, "新增群组角色已存在")
		return
	}
	if param.CreateUserId == 0 {
		param.CreateUserId = auth.GetUserId(context)
	}
	if err = logic.GroupRoleAdd(param); err != nil {
		response.BuildErrorResponse(context, err.Error())
	} else {
		response.BuildSuccessResponse(context, nil)
	}
}
