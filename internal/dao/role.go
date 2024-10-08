package dao

import (
	"fmt"

	"github.com/go-xuan/quanx/core/gormx"

	"user/internal/model"
	"user/internal/model/entity"
)

// RoleExist 查询群组编码是否存在
func RoleExist(in *model.RoleSave) (count int64, err error) {
	if err = gormx.DB().Model(&entity.Role{}).Where(`code = ? `, in.Code).Count(&count).Error; err != nil {
		return
	}
	return
}

// RoleCreate 角色新增
func RoleCreate(role *entity.Role) error {
	return gormx.DB().Create(role).Error
}

// RoleDelete 角色删除
func RoleDelete(id int64) error {
	return gormx.DB().Delete(&entity.Role{}, id).Error
}

// RoleUpdate 角色修改
func RoleUpdate(in *model.RoleSave) (err error) {
	var tx = gormx.DB().Begin()
	var cols = []string{"update_user_id", "update_time"}
	if in.Name != "" {
		cols = append(cols, "name")
	}
	if in.Remark != "" {
		cols = append(cols, "remark")
	}
	if err = tx.Model(&entity.Role{}).Select(cols).Where("role_id", in.Id).Updates(in.Role()).Error; err != nil {
		tx.Rollback()
		return
	}
	if len(in.UserList) > 0 {

	}
	tx.Commit()
	return
}

// RolePage 角色分页查询
func RolePage(in model.RolePage) (result []*model.Role, total int64, err error) {
	db := gormx.DB().Model(&entity.Role{})
	if in.Keyword != "" {
		db = db.Where(fmt.Sprintf("code like '%%%s%%' or name like '%%%s%%'", in.Keyword, in.Keyword))
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	if in.Page != nil && in.Page.PageSize > 0 {
		db.Limit(in.Page.PageSize).Offset(in.Page.Offset()).Order("update_time desc")
	}
	if err = db.Scan(&result).Error; err != nil {
		return
	}
	return
}

// RoleInfo 角色信息查询
func RoleInfo(id int64) (result *model.Role, err error) {
	result = &model.Role{}
	if err = gormx.DB().Model(&entity.Role{}).Where(`id = ?`, id).Scan(&result).Error; err != nil {
		return
	}
	return
}

// RoleList 用户角色
func RoleList() (resultList []*model.Role, err error) {
	if err = gormx.DB().Model(&entity.Role{}).Select([]string{"id", "code", "name"}).Order("id").Scan(&resultList).Error; err != nil {
		return
	}
	return
}

func RoleUserCount(roleId int64, userIds []int64) (count int64, err error) {
	if err = gormx.DB().Model(&entity.RoleUser{}).Where(`id = ? and user_id in ?`, roleId, userIds).Count(&count).Error; err != nil {
		return
	}
	return
}

// RoleUserCreateBatch 角色用户批量新增
func RoleUserCreateBatch(list []*entity.RoleUser) error {
	return gormx.DB().Create(&list).Error
}

// RoleUserList 角色成员列表
func RoleUserList(roleId int64) (resultList []*model.RoleUser, err error) {
	if err = gormx.DB().Raw(`
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
 order by t1.update_time desc`, roleId).Scan(&resultList).Error; err != nil {
		return
	}
	return
}
