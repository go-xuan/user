package dao

import (
	"strings"

	"github.com/go-xuan/quanx/importx/gormx"

	"user/internal/model"
	"user/internal/model/entity"
)

// 查询群组编码是否存在
func RoleExist(in *model.RoleSave) (count int64, err error) {
	err = gormx.This().DB.Model(&entity.Role{}).Where(`code = ? `, in.Code).Count(&count).Error
	if err != nil {
		return
	}
	return
}

// 角色新增
func RoleCreate(role *entity.Role) (err error) {
	err = gormx.This().DB.Create(role).Error
	if err != nil {
		return
	}
	return
}

// 角色删除
func RoleDelete(id int64) (err error) {
	err = gormx.This().DB.Delete(&entity.Role{}, id).Error
	if err != nil {
		return
	}
	return
}

// 角色修改
func RoleUpdate(in *model.RoleSave) (err error) {
	var tx = gormx.This().DB.Begin()
	var cols = []string{"update_user_id", "update_time"}
	if in.Name != "" {
		cols = append(cols, "name")
	}
	if in.Remark != "" {
		cols = append(cols, "remark")
	}
	err = tx.Model(&entity.Role{}).Select(cols).Where("role_id", in.Id).Updates(in.Role()).Error
	if len(in.UserList) > 0 {

	}
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// 角色分页查询
func RolePage(in model.RolePage) (result []*model.Role, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(`select * from t_sys_role`)
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
	if in.Page.Page != nil && in.Page.Page.PageSize > 0 {
		selectSql.WriteString(in.Page.Page.PgPageSql())
	}
	err = gormx.This().DB.Raw(selectSql.String()).Scan(&result).Error
	if err != nil {
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(` ) t`)
	err = gormx.This().DB.Raw(countSql.String()).Count(&total).Error
	if err != nil {
		return
	}
	return
}

// 角色信息查询
func RoleInfo(id int64) (result model.Role, err error) {
	err = gormx.This().DB.Model(&entity.Role{}).Where(`id = ?`, id).Scan(&result).Error
	if err != nil {
		return
	}
	return
}

// 用户角色
func RoleList() (resultList []*model.Role, err error) {
	err = gormx.This().DB.Model(&entity.Role{}).Select([]string{"id", "code", "name"}).Order("id").Scan(&resultList).Error
	if err != nil {
		return
	}
	return
}

func RoleUserCount(roleId int64, userIds []int64) (count int64, err error) {
	err = gormx.This().DB.Model(&entity.RoleUser{}).Where(`id = ? and user_id in ?`, roleId, userIds).Count(&count).Error
	if err != nil {
		return
	}
	return
}

// 角色用户批量新增
func RoleUserCreateBatch(list []*entity.RoleUser) (err error) {
	err = gormx.This().DB.Create(&list).Error
	if err != nil {
		return
	}
	return
}

// 角色成员列表
func RoleUserList(roleId int64) (resultList []*model.RoleUser, err error) {
	err = gormx.This().DB.Raw(`
select t2.id,
       t2.name,
       t2.account,
       t1.valid_start,
       t1.valid_end,
       t1.remark
  from t_sys_role_user t1
  left join sys_user t2
    on t1.user_id = t2.id
 where t1.id = ?
   and valid_start <= now()
   and valid_end >= now()
 order by t1.update_time desc`, roleId).Scan(&resultList).Error
	if err != nil {
		return
	}
	return
}
