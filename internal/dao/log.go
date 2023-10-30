package dao

import (
	"github.com/quanxiaoxuan/quanx/public/gormx"

	"quan-user/model/table"
)

// 更新用户身份信息
func SysLogCreate(syslog table.Log) (err error) {
	tx := gormx.CTL.DB.Begin()
	err = tx.Create(&syslog).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// 更新用户身份信息
func LogDelete(types []string) (err error) {
	tx := gormx.CTL.DB.Begin()
	err = tx.Delete(&table.Log{}, `type in ? and ip = ?`, types, "192.168.152.63").Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}
