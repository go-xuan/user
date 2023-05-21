package entity

import "time"

// 用户鉴权表
type SysUserAuth struct {
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; primary_key; comment:用户ID;"`
	Password     string    `json:"password" gorm:"type:varchar(100); not null; comment:密码;"`
	Salt         string    `json:"salt" gorm:"type:varchar(100); not null; comment:密码盐;"`
	SessionTime  int64     `json:"sessionTime" gorm:"type:int; not null; comment:会话有效期(秒);"`
	ValidStart   time.Time `json:"validStart" gorm:"type:timestamp(0); not null; comment:账号有效期始;"`
	ValidEnd     time.Time `json:"validEnd" gorm:"type:timestamp(0); not null; comment:账号有效期止;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysUserAuth) TableName() string {
	return "sys_user_auth"
}
