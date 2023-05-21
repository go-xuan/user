package entity

import "time"

type SysGroup struct {
	GroupId      int64     `json:"groupId" gorm:"type:bigint; not null; primary_key; comment:群组ID;"`
	GroupCode    string    `json:"groupCode" gorm:"type:varchar(100); not null; comment:群组编码;"`
	GroupName    string    `json:"groupName" gorm:"type:varchar(100); not null; comment:群组名称;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysGroup) TableName() string {
	return "sys_group"
}
