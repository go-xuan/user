package logic

import (
	"errors"
	
	"github.com/quanxiaoxuan/quanx/common/snowflakex"
	"github.com/quanxiaoxuan/quanx/utils/timex"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
)

func RoleMemberAdd(param params.RoleMemberAdd) error {
	var userIds []int64
	for _, item := range param.RoleMemberList {
		userIds = append(userIds, item.UserId)
	}
	count, err := mapper.RoleMemberCount(param.RoleId, userIds)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("新增成员已存在")
	}
	var memberList entity.SysRoleMemberList
	for _, item := range param.RoleMemberList {
		var roleUser = entity.SysRoleMember{
			Id:           snowflakex.Generator.NewId(),
			RoleId:       param.RoleId,
			MemberId:     item.UserId,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: param.CreateUserId,
			UpdateUserId: param.CreateUserId,
		}
		memberList = append(memberList, &roleUser)
	}
	return mapper.RoleMemberAddBatch(memberList)
}
