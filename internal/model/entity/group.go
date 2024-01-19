package entity

import "time"

type Group struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:群组ID;"`
	Code         string    `json:"code" gorm:"type:varchar(100); not null; comment:群组编码;"`
	Name         string    `json:"name" gorm:"type:varchar(100); not null; comment:群组名称;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (Group) TableName() string {
	return "t_sys_group"
}

func (g Group) Comment() string {
	return "群组表"
}

func (g Group) InitData() interface{} {
	return nil
}

type GroupUser struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	GroupId      int64     `json:"groupId" gorm:"type:bigint; not null; comment:群组ID;"`
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; comment:用户ID;"`
	IsAdmin      bool      `json:"isAdmin" gorm:"type:bool; not null; comment:是否管理员;"`
	ValidStart   time.Time `json:"validStart" gorm:"type:timestamp(0); not null; comment:有效期始;"`
	ValidEnd     time.Time `json:"validEnd" gorm:"type:timestamp(0); not null; comment:有效期止;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (GroupUser) TableName() string {
	return "t_sys_group_user"
}

func (u GroupUser) Comment() string {
	return "群组成员表"
}

func (u GroupUser) InitData() interface{} {
	return nil
}

type GroupRole struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	GroupId      int64     `json:"groupId" gorm:"type:bigint; not null; comment:群组ID;"`
	RoleId       int64     `json:"roleId" gorm:"type:bigint; not null; comment:角色ID;"`
	ValidStart   time.Time `json:"validStart" gorm:"type:timestamp(0); not null; comment:有效期始;"`
	ValidEnd     time.Time `json:"validEnd" gorm:"type:timestamp(0); not null; comment:有效期止;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (GroupRole) TableName() string {
	return "t_sys_group_role"
}

func (r GroupRole) Comment() string {
	return "群组角色表"
}

func (r GroupRole) InitData() interface{} {
	return nil
}
