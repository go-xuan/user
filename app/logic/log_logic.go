package logic

import (
	"quan-admin/app/mapper"
	"quan-admin/conf"
	"quan-admin/model/entity"
)

// 新增日志
func LogAdd(sysLog entity.SysLog) error {
	sysLog.Id = conf.NewSnow.NewId()
	return mapper.SysLogAdd(sysLog)
}

// 新增日志
func LogDelete(types []string) (err error) {
	return mapper.LogDelete(types)
}
