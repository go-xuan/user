package entity

import "time"

type SysRole struct {
	RoleId       int64     `json:"roleId" gorm:"type:bigint; not null; primary_key; comment:角色ID;"`
	RoleCode     string    `json:"roleCode" gorm:"type:varchar(100); not null; comment:角色编码;"`
	RoleName     string    `json:"roleName" gorm:"type:varchar(100); not null; comment:角色名称;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
