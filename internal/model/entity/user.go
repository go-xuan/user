package entity

import (
	"github.com/go-xuan/quanx/utils/idx"
	"github.com/go-xuan/quanx/utils/randx"
	"time"

	"github.com/go-xuan/quanx/utils/encryptx"
)

// User 用户表
type User struct {
	Id           int64     `json:"id" gorm:"type:bigint; primary_key; comment:用户ID;"`
	Name         string    `json:"name" gorm:"type:varchar(100); not null; comment:姓名;"`
	Account      string    `json:"account" gorm:"type:varchar(20); not null; comment:用户账号;"`
	Phone        string    `json:"phone" gorm:"type:varchar(100); not null; comment:手机;"`
	Gender       string    `json:"gender" gorm:"type:varchar(100); comment:性别;"`
	Birthday     string    `json:"birthday" gorm:"type:date; comment:生日;"`
	Email        string    `json:"email" gorm:"type:varchar(100); comment:邮箱;"`
	Address      string    `json:"address" gorm:"type:varchar(1000); comment:地址;"`
	Remark       string    `json:"remark" gorm:"type:varchar(1000); comment:备注;"`
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

func (*User) TableName() string {
	return "t_sys_user"
}

func (*User) TableComment() string {
	return "用户信息表"
}

func (*User) InitData() any {
	var data []*User
	salt := randx.UUID()
	password := encryptx.PasswordSalt(encryptx.MD5("Init@123"), salt)
	data = append(data, &User{
		Id:           idx.SnowFlake().Int64(),
		Name:         "Admin",
		Account:      "admin",
		Phone:        "15172364557",
		Gender:       "male",
		Birthday:     "2000-01-01",
		Email:        "quanchao1996@qq.com",
		Address:      randx.City(),
		Remark:       "系统管理员",
		Password:     password,
		Salt:         salt,
		SessionTime:  1800,
		ValidStart:   time.Now(),
		ValidEnd:     time.Now().AddDate(3, 0, 0),
		CreateUserId: 9999999999,
		UpdateUserId: 9999999999,
	})
	return data
}

func (a *User) CheckPassword(password string) bool {
	var err error
	if password, err = encryptx.RSA().Decrypt(password); err == nil {
		if encryptx.PasswordSalt(encryptx.MD5(password), a.Salt) == a.Password {
			return true
		}
	}
	return false
}
