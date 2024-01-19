package model

import (
	"github.com/go-xuan/quanx/utilx/anyx"
	"time"

	"github.com/go-xuan/quanx/importx/encryptx"
	"github.com/go-xuan/quanx/utilx/randx"
	"github.com/go-xuan/quanx/utilx/timex"

	"user/internal/model/entity"
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
	*entity.User     // 用户基本信息
	*entity.UserAuth // 用户鉴权信息
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

func (u *UserSave) UserCreate() *entity.User {
	return &entity.User{
		Id:           u.Id,
		Account:      u.Account,
		Name:         u.Name,
		Phone:        u.Phone,
		Gender:       u.Gender,
		Birthday:     u.Birthday,
		Email:        u.Email,
		Address:      u.Address,
		Remark:       u.Remark,
		CreateUserId: u.CurrUserId,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
}

func (u *UserSave) UserUpdate() *entity.User {
	return &entity.User{
		Id:           u.Id,
		Account:      u.Account,
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

func (u *UserSave) UserAuthCreate() (auth *entity.UserAuth) {
	auth = &entity.UserAuth{
		UserId:       u.Id,
		SessionTime:  3600,
		CreateUserId: u.CurrUserId,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
	var salt = randx.UUID()
	auth.Salt = salt
	auth.Password = encryptx.PasswordSalt(encryptx.MD5(u.Password), salt)
	var validStart = anyx.IfElse(u.ValidStart != "", timex.ToTime(u.ValidStart), time.Now())
	var validEnd = anyx.IfElse(u.ValidEnd != "", timex.ToTime(u.ValidEnd), validStart.AddDate(1, 0, 0))
	auth.ValidStart = validStart
	auth.ValidEnd = validEnd
	return
}

func (u *UserSave) UserAuthUpdate() (auth *entity.UserAuth) {
	auth = &entity.UserAuth{
		UserId:       u.Id,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
	if u.Password != "" {
		var salt = randx.UUID()
		var password = encryptx.PasswordSalt(encryptx.MD5(u.Password), salt)
		auth.Password = password
		auth.Salt = salt
	}
	auth.ValidStart = anyx.IfElse(u.ValidStart != "", timex.ToTime(u.ValidStart), time.Time{})
	auth.ValidEnd = anyx.IfElse(u.ValidEnd != "", timex.ToTime(u.ValidEnd), time.Time{})
	return
}
