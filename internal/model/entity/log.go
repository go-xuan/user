package entity

import "time"

// Log 操作日志
type Log struct {
	Id           int64     `json:"id" gorm:"type:bigint; primary_key; comment:日志ID"`
	Service      string    `json:"service" gorm:"type:varchar(100); comment:所属服务"`
	Type         string    `json:"type" gorm:"type:varchar(100); comment:日志类型"`
	Operation    string    `json:"operation" gorm:"type:varchar(100); comment:操作"`
	Url          string    `json:"url" gorm:"type:varchar(100); comment:URL"`
	Content      string    `json:"content" gorm:"type:text; comment:日志详情"`
	Ip           string    `json:"ip" gorm:"type:varchar(100); comment:操作IP"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; comment:创建人ID"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间"`
}

func (Log) TableName() string {
	return "t_sys_log"
}

func (Log) TableComment() string {
	return "日志表"
}

func (Log) InitData() any {
	return nil
}
