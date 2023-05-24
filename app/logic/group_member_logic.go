package logic

import (
	"github.com/quanxiaoxuan/go-builder/snowflakex"
	"github.com/quanxiaoxuan/go-utils/timex"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
	"quan-admin/model/params"
)

// 群组成员新增
func GroupMemberAdd(param params.GroupMemberAdd) error {
	var memberList entity.SysGroupMemberList
	for _, item := range param.GroupMemberList {
		var member = entity.SysGroupMember{
			Id:           snowflakex.NewSnow.NewId(),
			GroupId:      param.GroupId,
			MemberId:     item.UserId,
			IsAdmin:      item.IsAdmin,
			ValidStart:   timex.ToTime(item.ValidStart),
			ValidEnd:     timex.ToTime(item.ValidEnd),
			Remark:       item.Remark,
			CreateUserId: param.CreateUserId,
			UpdateUserId: param.CreateUserId,
		}
		memberList = append(memberList, &member)
	}
	return mapper.GroupMemberAddBatch(memberList)
}
