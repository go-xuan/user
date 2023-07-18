package mapper

import (
	"github.com/quanxiaoxuan/quanx/middleware/gormx"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
	"quan-admin/model/results"
)

func RoleMemberCount(roleId int64, userIds []int64) (count int64, err error) {
	err = gormx.CTL.DB.Model(&entity.SysRoleMember{}).Where(`role_id = ? and member_id in ?`, roleId, userIds).Count(&count).Error
	return
}

// 角色用户批量新增
func RoleMemberAddBatch(memberList entity.SysRoleMemberList) error {
	err := gormx.CTL.DB.Create(&memberList).Error
	if err != nil {
		log.Error("角色用户批量新增失败 ： ", err)
		return err
	}
	return nil
}

// 角色成员列表
func RoleMemberList(roleId int64) (resultList results.RoleMemberList, err error) {
	err = gormx.CTL.DB.Raw(`
select t2.user_id,
       t2.user_name,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from sys_role_member t1
  left join sys_user t2
    on t1.member_id = t2.user_id
 where t1.role_id = ?
   and valid_start <= now()
   and valid_end >= now()
 order by t1.update_time desc`, roleId).Scan(&resultList).Error
	return
}
