package mapper

import (
	"github.com/quanxiaoxuan/go-builder/database"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
	"quan-admin/model/results"
)

// 群组成员校验
func GroupMemberCount(groupId int64, userIds []int64) (count int64, err error) {
	err = database.GormDB.Model(&entity.SysGroupMember{}).Where(`group_id = ? and member_id in ?`, groupId, userIds).Count(&count).Error
	return
}

// 群组成员批量新增
func GroupMemberAddBatch(memberList entity.SysGroupMemberList) error {
	err := database.GormDB.Create(&memberList).Error
	if err != nil {
		log.Error("群组成员批量新增失败 ： ", err)
		return err
	}
	return nil
}

// 查询群组成员列表
func GroupMemberList(groupId int64) (resultList results.GroupMemberList, err error) {
	err = database.GormDB.Raw(`
select t2.user_id,
       t2.user_name,
       t1.is_admin,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from sys_group_member t1
  left join sys_user t2
    on t1.member_id = t2.user_id
 where t1.group_id = ?
 order by t1.is_admin desc, t1.update_time desc`, groupId).Scan(&resultList).Error
	return
}
