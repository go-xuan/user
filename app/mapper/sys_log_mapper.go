package mapper

import (
	"github.com/quanxiaoxuan/go-builder/gormx"
	log "github.com/sirupsen/logrus"

	"quan-admin/model/entity"
)

// 更新用户身份信息
func SysLogAdd(syslog entity.SysLog) (err error) {
	tx := gormx.CTL.DB.Begin()
	err = tx.Create(&syslog).Error
	if err != nil {
		tx.Rollback()
		log.Error("日志新增失败 ： ", err)
		return
	}
	tx.Commit()
	return
}

// 更新用户身份信息
func LogDelete(types []string) (err error) {
	tx := gormx.CTL.DB.Begin()
	err = tx.Delete(&entity.SysLog{}, `type in ? and ip = ?`, types, "192.168.152.63").Error
	if err != nil {
		tx.Rollback()
		log.Error("日志新增失败 ： ", err)
		return
	}
	tx.Commit()
	return
}
