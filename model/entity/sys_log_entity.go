package entity

import "time"

type SysLog struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:日志ID;"`
	Module       string    `json:"module" gorm:"type:varchar(100); comment:操作模块;"`
	Type         string    `json:"type" gorm:"type:varchar(100); comment:操作类型;"`
	Content      string    `json:"content" gorm:"type:varchar(1000); comment:操作内容;"`
	Ip           string    `json:"ip" gorm:"type:varchar(100); comment:当前登录IP;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
}

func (SysLog) TableName() string {
	return "sys_log"
}
