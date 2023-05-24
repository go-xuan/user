package logic

import (
	"github.com/quanxiaoxuan/go-builder/snowflakex"

	"quan-admin/app/mapper"
	"quan-admin/model/entity"
)

// 新增日志
func LogAdd(sysLog entity.SysLog) error {
	sysLog.Id = snowflakex.NewSnow.NewId()
	return mapper.SysLogAdd(sysLog)
}

// 新增日志
func LogDelete(types []string) (err error) {
	return mapper.LogDelete(types)
}
