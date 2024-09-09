package dao

import (
	"github.com/go-xuan/quanx/server/gormx"

	"user/internal/model/entity"
)

// AddLog 新增日志
func AddLog(log *entity.Log) error {
	return gormx.GetDB().Create(log).Error
}
