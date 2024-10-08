package entity

import "time"

// Role 角色
type Role struct {
	Id           int64     `json:"roleId" gorm:"type:bigint; primary_key; comment:角色ID;"`
	Code         string    `json:"roleCode" gorm:"type:varchar(100); not null; comment:角色编码;"`
	Name         string    `json:"roleName" gorm:"type:varchar(100); not null; comment:角色名称;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (Role) TableName() string {
	return "t_sys_role"
}

func (Role) TableComment() string {
	return "角色表"
}

func (Role) InitData() any {
	return nil
}

// RoleUser 角色所属用户
type RoleUser struct {
	Id           int64     `json:"id" gorm:"type:bigint; primary_key; comment:主键ID;"`
	RoleId       int64     `json:"roleId" gorm:"type:bigint; not null; comment:角色ID;"`
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; comment:成员ID;"`
	ValidStart   time.Time `json:"validStart" gorm:"type:timestamp(0); not null; comment:有效期始;"`
	ValidEnd     time.Time `json:"validEnd" gorm:"type:timestamp(0); not null; comment:有效期终;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (RoleUser) TableName() string {
	return "t_sys_role_user"
}

func (RoleUser) TableComment() string {
	return "角色成员表"
}

func (RoleUser) InitData() any {
	return nil
}
