package entity

import "time"

type SysUser struct {
	UserId       int64     `json:"userId" gorm:"type:bigint; not null; primary_key; comment:用户ID;"`
	UserName     string    `json:"userName" gorm:"type:varchar(100); not null; comment:姓名;"`
	Phone        string    `json:"phone" gorm:"type:varchar(100); not null; comment:手机;"`
	Gender       string    `json:"gender" gorm:"type:varchar(100); comment:性别;"`
	Birthday     string    `json:"birthday" gorm:"type:date; comment:生日;"`
	Email        string    `json:"email" gorm:"type:varchar(100); comment:邮箱;"`
	Address      string    `json:"address" gorm:"type:varchar(1000); comment:地址;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
	CreateUserId int64     `json:"createUserId" gorm:"type:bigint; not null; comment:创建人ID;"`
	CreateTime   time.Time `json:"createTime" gorm:"type:timestamp(0); default:now(); comment:创建时间;"`
	UpdateUserId int64     `json:"updateUserId" gorm:"type:bigint; not null; comment:更新人ID;"`
	UpdateTime   time.Time `json:"updateTime" gorm:"type:timestamp(0); default:now(); comment:更新时间;"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
