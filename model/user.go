package model

import (
	"time"

	"github.com/quanxiaoxuan/quanx/utils/encryptx"
	"github.com/quanxiaoxuan/quanx/utils/randx"
	"github.com/quanxiaoxuan/quanx/utils/timex"

	"quan-user/model/table"
)

// 用户信息
type User struct {
	Id          int64      `json:"id" comment:"用户ID"`
	Account     string     `json:"account" comment:"用户账号"`
	Name        string     `json:"name" comment:"姓名"`
	Phone       string     `json:"phone" comment:"手机"`
	Gender      string     `json:"gender" comment:"性别"`
	Birthday    timex.Date `json:"birthday" comment:"生日"`
	Email       string     `json:"email" comment:"邮箱"`
	Address     string     `json:"address" comment:"地址"`
	Remark      string     `json:"remark" comment:"备注"`
	SessionTime int64      `json:"sessionTime" comment:"会话有效期"`
	ValidStart  timex.Time `json:"validStart" comment:"账号有效期始"`
	ValidEnd    timex.Time `json:"validEnd" comment:"账号有效期止"`
}

// 用户明细
type UserDetail struct {
	*table.User     // 用户基本信息
	*table.UserAuth // 用户鉴权信息
}

// 用户分页参数
type UserPage struct {
	Page
}

// 用户登录参数
type Login struct {
	Username string `json:"username" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
}

// 用户登录返回结果
type LoginResult struct {
	User  *User  `json:"user"`
	Token string `json:"token" comment:"token"`
}

// 用户保存参数
type UserSave struct {
	Id         int64  `json:"id" comment:"用户ID"`
	Account    string `json:"account" comment:"账号"`
	Name       string `json:"name" comment:"姓名"`
	Password   string `json:"password" comment:"密码"`
	Phone      string `json:"phone" comment:"手机"`
	Gender     string `json:"gender" comment:"性别"`
	Birthday   string `json:"birthday" comment:"生日"`
	Email      string `json:"email" comment:"邮箱"`
	Address    string `json:"address" comment:"地址"`
	Remark     string `json:"remark" comment:"备注"`
	ValidStart string `json:"validStart" comment:"账号有效期始"`
	ValidEnd   string `json:"validEnd" comment:"账号有效期止"`
	CurrUserId int64  `json:"currUserId" comment:"当前用户ID"`
}

func (u *UserSave) User() *table.User {
	return &table.User{
		Id:           u.Id,
		Name:         u.Name,
		Phone:        u.Phone,
		Gender:       u.Gender,
		Birthday:     u.Birthday,
		Email:        u.Email,
		Address:      u.Address,
		Remark:       u.Remark,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
}

func (u *UserSave) UserAuth() *table.UserAuth {
	var auth = &table.UserAuth{
		UserId:       u.Id,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
	if u.Password != "" {
		passWord, salt := u.Password, randx.UUID()
		if len(passWord) < 32 {
			passWord = encryptx.MD5(passWord)
		}
		passWord = encryptx.PasswordSalt(passWord, salt)
		auth.Password = passWord
		auth.Salt = salt
	}
	if u.ValidStart != "" {
		auth.ValidStart = timex.ToTime(u.ValidStart)
	}
	if u.ValidEnd != "" {
		auth.ValidEnd = timex.ToTime(u.ValidEnd)
	}
	return auth
}
