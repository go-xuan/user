package mapper

import (
	"strings"

	"github.com/quanxiaoxuan/go-builder/gormx"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
	"quan-admin/model/params"
	"quan-admin/model/results"
)

// 查询群组编码是否存在
func RoleCodeExist(roleCode string) (count int64, err error) {
	err = gormx.GormDB.Model(&entity.SysRole{}).Where(`role_code = ? `, roleCode).Count(&count).Error
	return
}

// 角色新增
func RoleAdd(role entity.SysRole) error {
	err := gormx.GormDB.Create(&role).Error
	if err != nil {
		log.Error("角色新增失败 ： ", err)
		return err
	}
	return nil
}

// 角色删除
func RoleDelete(roleId int64) error {
	err := gormx.GormDB.Delete(&entity.SysRole{}, roleId).Error
	if err != nil {
		log.Error("角色删除失败 ： ", err)
		return err
	}
	return nil
}

// 角色修改
func RoleUpdate(param params.RoleUpdate) error {
	tx := gormx.GormDB.Begin()
	sql := strings.Builder{}
	sql.WriteString(`update sys_role set update_time = now(), update_user_id = ? `)
	if param.RoleName != "" {
		sql.WriteString(`, role_name = '` + param.RoleName + `'`)
	}
	if param.Remark != "" {
		sql.WriteString(`, remark = '` + param.Remark + `'`)
	}
	sql.WriteString(` where role_id = ?`)
	err := tx.Exec(sql.String(), param.UpdateUserId, param.RoleId).Error
	if err != nil {
		tx.Rollback()
		log.Error("角色修改失败 ： ", err)
		return err
	}
	tx.Commit()
	return nil
}

// 角色分页查询
func RolePage(param params.RolePage) (resultList results.RoleInfoList, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(`select * from sys_role`)
	if param.SearchKey != "" {
		sql.WriteString(` where role_code like '%` + param.SearchKey + `%'`)
		sql.WriteString(` or role_name like '%` + param.SearchKey + `%'`)
	}
	selectSql := strings.Builder{}
	selectSql.WriteString(`select * from`)
	selectSql.WriteString(` ( `)
	selectSql.WriteString(sql.String())
	selectSql.WriteString(`) t order by role_id `)
	if param.PageParam != nil || param.PageParam.PageSize > 0 {
		selectSql.WriteString(param.PageParam.GetPgPageSql())
	}
	err = gormx.GormDB.Raw(selectSql.String()).Scan(&resultList).Error
	if err != nil {
		log.Error("角色分页查询失败 ： ", err)
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(`) t`)
	err = gormx.GormDB.Raw(countSql.String()).Scan(&total).Error
	if err != nil {
		log.Error("角色分页查询失败 ： ", err)
		return
	}
	return
}

// 角色信息查询
func RoleInfo(roleId int64) (result results.RoleInfo, err error) {
	err = gormx.GormDB.Model(&entity.SysRole{}).Where(`role_id = ?`, roleId).Scan(&result).Error
	return
}

// 用户角色
func RoleList() (resultList results.RoleSimpleList, err error) {
	err = gormx.GormDB.Model(&entity.SysRole{}).Select([]string{"role_id", "role_code", "role_name"}).Order("role_id").Scan(&resultList).Error
	return
}
