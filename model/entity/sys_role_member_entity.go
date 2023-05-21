package entity

import "time"

// 角色所属用户
type SysRoleMemberList []*SysRoleMember
type SysRoleMember struct {
	Id           int64     `json:"id" gorm:"type:bigint; not null; primary_key; comment:主键ID;"`
	RoleId       int64     `json:"roleId" gorm:"type:bigint; not null; comment:角色ID;"`
	MemberId     int64     `json:"memberId" gorm:"type:bigint; not null; comment:成员ID;"`
	ValidStart   time.Time `json:"validStart" gorm:"type:timestamp(0); not null; comment:有效期始;"`
	ValidEnd     time.Time `json:"validEnd" gorm:"type:timestamp(0); not null; comment:有效期终;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysRoleMember) TableName() string {
	return "sys_role_member"
}
