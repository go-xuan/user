package dao

import (
	"strings"

	"github.com/go-xuan/quanx/server/gormx"

	"user/internal/model"
	"user/internal/model/entity"
)

// 群组分页
func GroupPage(in model.GroupPage) (result []*model.Group, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(`select * from t_sys_group`)
	if in.Keyword != "" {
		sql.WriteString(` where code like '%`)
		sql.WriteString(in.Keyword)
		sql.WriteString(`%' or name like '%`)
		sql.WriteString(in.Keyword)
		sql.WriteString(`%'`)
	}
	selectSql := strings.Builder{}
	selectSql.WriteString(`select * from`)
	selectSql.WriteString(` ( `)
	selectSql.WriteString(sql.String())
	selectSql.WriteString(` ) t order by update_time desc`)
	if in.Page != nil && in.Page.PageSize > 0 {
		selectSql.WriteString(in.Page.PgPageSql())
	}
	if err = gormx.This().DB.Raw(selectSql.String()).Scan(&result).Error; err != nil {
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(` ) t`)
	if err = gormx.This().DB.Raw(countSql.String()).Scan(&total).Error; err != nil {
		return
	}
	return
}

// 查询群组编码是否存在
func GroupExist(in *model.GroupSave) (count int64, err error) {
	if err = gormx.This().DB.Model(&entity.Group{}).Where(`code = ? `, in.Code).Count(&count).Error; err != nil {
		return
	}
	return
}

// 群组新增
func GroupCreate(group *entity.Group) error {
	return gormx.This().DB.Create(group).Error
}

// 群组修改
func GroupUpdate(in *model.GroupSave) (err error) {
	var cols = []string{"update_user_id", "update_time"}
	if in.Name != "" {
		cols = append(cols, "name")
	}
	if in.Remark != "" {
		cols = append(cols, "remark")
	}
	return gormx.This().DB.Model(&entity.Group{}).Select(cols).Where("id = ? ", in.Id).Updates(in.Group()).Error
}

// 群组删除
func GroupDelete(id int64) error {
	return gormx.This().DB.Delete(&entity.Group{}, id).Error
}

// 查询群组信息
func GetGroup(id int64) (result *model.Group, err error) {
	result = &model.Group{}
	if err = gormx.This().DB.Model(&entity.Group{}).Where(`id = ?`, id).Scan(&result).Error; err != nil {
		return
	}
	return
}

// 群组成员校验
func GroupUserCount(id int64, userIds []int64) (count int64, err error) {
	if err = gormx.This().DB.Model(&entity.GroupUser{}).Where(`id = ? and user_id in ?`, id, userIds).Count(&count).Error; err != nil {
		return
	}
	return
}

// 群组成员批量新增
func GroupUserCreateBatch(list []*entity.GroupUser) error {
	return gormx.This().DB.Create(&list).Error
}

// 查询群组成员列表
func GroupUserList(id int64) (resultList []*model.GroupUser, err error) {
	if err = gormx.This().DB.Raw(`
select t2.id,
       t2.name,
       t1.is_admin,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from t_sys_group_user t1
  left join t_sys_user t2
    on t1.user_id = t2.id
 where t1.id = ?
 order by t1.is_admin desc, t1.update_time desc`, id).Scan(&resultList).Error; err != nil {
		return
	}
	return
}

// 群组角色校验
func GroupRoleCount(id int64, roleIds []int64) (count int64, err error) {
	if err = gormx.This().DB.Model(&entity.GroupRole{}).Where(`id = ? and role_id in ?`, id, roleIds).Count(&count).Error; err != nil {
		return
	}
	return
}

// 群组角色批量新增
func GroupRoleCreateBatch(list []*entity.GroupRole) error {
	return gormx.This().DB.Create(&list).Error
}

// 查询群组角色列表
func GroupRoleList(id int64) (resultList []*model.GroupRole, err error) {
	if err = gormx.This().DB.Raw(`
select t2.id,
       t2.code,
       t2.name,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from t_sys_group_role t1
  left join t_sys_role t2
    on t1.role_id = t2.id
 where t1.id = ?
 order by t1.update_time desc`, id).Scan(&resultList).Error; err != nil {
		return
	}
	return
}
