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
func GroupCodeCount(groupCode string) (count int64, err error) {
	err = gormx.GormDB.Model(&entity.SysGroup{}).Where(`group_code = ? `, groupCode).Count(&count).Error
	return
}

// 群组新增
func GroupAdd(group entity.SysGroup) error {
	err := gormx.GormDB.Create(&group).Error
	if err != nil {
		log.Error("群组新增失败 ： ", err)
		return err
	}
	return nil
}

// 群组修改
func GroupUpdate(param params.GroupUpdate) error {
	tx := gormx.GormDB.Begin()
	sql := strings.Builder{}
	sql.WriteString(`update sys_group set update_time = now(), update_user_id = ? `)
	if param.GroupName != "" {
		sql.WriteString(`, group_name = '` + param.GroupName + `'`)
	}
	if param.Remark != "" {
		sql.WriteString(`, remark = '` + param.Remark + `'`)
	}
	sql.WriteString(` where group_id = ?`)
	err := tx.Exec(sql.String(), param.UpdateUserId, param.GroupId).Error
	if err != nil {
		tx.Rollback()
		log.Error("群组修改失败 ： ", err)
		return err
	}
	tx.Commit()
	return nil
}

// 群组删除
func GroupDelete(groupId int64) error {
	err := gormx.GormDB.Delete(&entity.SysGroup{}, groupId).Error
	if err != nil {
		log.Error("群组删除失败 ： ", err)
		return err
	}
	return nil
}

// 群组分页
func GroupPage(param params.GroupPage) (resultList results.GroupInfoList, total int64, err error) {
	sql := strings.Builder{}
	sql.WriteString(`select * from sys_group`)
	if param.SearchKey != "" {
		sql.WriteString(` where group_code like '%` + param.SearchKey + `%'`)
		sql.WriteString(` or group_name like '%` + param.SearchKey + `%'`)
	}
	selectSql := strings.Builder{}
	selectSql.WriteString(`select * from`)
	selectSql.WriteString(` ( `)
	selectSql.WriteString(sql.String())
	selectSql.WriteString(` ) t order by group_id `)
	if param.PageParam != nil || param.PageParam.PageSize > 0 {
		selectSql.WriteString(param.PageParam.GetPgPageSql())
	}
	err = gormx.GormDB.Raw(selectSql.String()).Scan(&resultList).Error
	if err != nil {
		log.Error("群组分页查询失败 ： ", err)
		return
	}
	countSql := strings.Builder{}
	countSql.WriteString(`select count(*) from`)
	countSql.WriteString(` ( `)
	countSql.WriteString(sql.String())
	countSql.WriteString(`) t`)
	err = gormx.GormDB.Raw(countSql.String()).Scan(&total).Error
	if err != nil {
		log.Error("群组分页查询失败 ： ", err)
		return
	}
	return
}

// 查询群组信息
func GroupInfo(groupId int64) (result results.GroupInfo, err error) {
	err = gormx.GormDB.Model(&entity.SysGroup{}).Where(`group_id = ?`, groupId).Scan(&result).Error
	return
}
