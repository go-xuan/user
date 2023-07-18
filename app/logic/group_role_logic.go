package logic

import (
	"github.com/quanxiaoxuan/quanx/common/snowflakex"
	"github.com/quanxiaoxuan/quanx/utils/timex"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
)

// 群组角色校验
func GroupRoleExist(groupId int64, roleIds []int64) (exist bool, err error) {
	var count int64
	count, err = mapper.GroupRoleCount(groupId, roleIds)
	if err != nil {
		return
	}
	exist = count > 0
	return
}

// 群组角色新增
func GroupRoleAdd(param params.GroupRoleAdd) error {
	var groupRoleList entity.SysGroupRoleList
	for _, item := range param.GroupRoleList {
		var groupRole = entity.SysGroupRole{
			Id:           snowflakex.Generator.NewId(),
			GroupId:      param.GroupId,
			RoleId:       item.RoleId,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: param.CreateUserId,
			UpdateUserId: param.CreateUserId,
		}
		groupRoleList = append(groupRoleList, &groupRole)
	}
	return mapper.GroupRoleAddBatch(groupRoleList)
}
