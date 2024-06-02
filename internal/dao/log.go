package dao

import (
	"github.com/go-xuan/quanx/server/gormx"

	"user/internal/model/entity"
)

// 更新用户身份信息
func SysLogCreate(syslog entity.Log) error {
	return gormx.This().DB.Create(&syslog).Error
}

// 更新用户身份信息
func LogDelete(types []string) error {
	return gormx.This().DB.Delete(&entity.Log{}, `type in ? and ip = ?`, types, "192.168.152.63").Error
}
