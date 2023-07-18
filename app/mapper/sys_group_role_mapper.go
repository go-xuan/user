package mapper

import (
	"github.com/quanxiaoxuan/quanx/middleware/gormx"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
	"quan-admin/model/results"
)

// 群组角色校验
func GroupRoleCount(groupId int64, roleIds []int64) (count int64, err error) {
	err = gormx.CTL.DB.Model(&entity.SysGroupRole{}).Where(`group_id = ? and role_id in ?`, groupId, roleIds).Count(&count).Error
	return
}

// 群组角色批量新增
func GroupRoleAddBatch(memberList entity.SysGroupRoleList) error {
	err := gormx.CTL.DB.Create(&memberList).Error
	if err != nil {
		log.Error("群组角色批量新增失败 ： ", err)
		return err
	}
	return nil
}

// 查询群组角色列表
func GroupRoleList(groupId int64) (resultList results.GroupRoleList, err error) {
	err = gormx.CTL.DB.Raw(`
select t1.role_id,
       t2.role_code,
       t2.role_name,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from sys_group_role t1
  left join sys_role t2
    on t1.role_id = t2.role_id
 where t1.group_id = ?
 order by t1.update_time desc`, groupId).Scan(&resultList).Error
	return
}
