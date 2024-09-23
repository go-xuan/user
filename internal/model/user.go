package model

import (
	"time"

	"github.com/go-xuan/quanx/common/modelx"
	"github.com/go-xuan/quanx/types/anyx"
	"github.com/go-xuan/quanx/types/timex"
	"github.com/go-xuan/quanx/utils/encryptx"
	"github.com/go-xuan/quanx/utils/randx"

	"user/internal/model/entity"
)

// User 用户信息
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
	Password    string     `json:"-" comment:"密码"`
	Salt        string     `json:"-" comment:"密码盐"`
	SessionTime int        `json:"sessionTime" comment:"会话有效期"`
	ValidStart  timex.Time `json:"validStart" comment:"账号有效期始"`
	ValidEnd    timex.Time `json:"validEnd" comment:"账号有效期止"`
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

// UserPage 用户分页参数
type UserPage struct {
	*modelx.Page
	Keyword string `json:"keyword" comment:"关键字"`
}

// Login 用户登录参数
type Login struct {
	Username string `json:"username" v:"required" comment:"用户名"`
	Password string `json:"password" comment:"密码"`
}

// LoginResult 用户登录返回结果
type LoginResult struct {
	User  *User  `json:"user"`
	Token string `json:"token" comment:"token"`
}

// UserSave 用户保存参数
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
	var salt = randx.UUID()
	var password = encryptx.PasswordSalt(encryptx.MD5(u.Password), salt)
	var validStart = anyx.If(u.ValidStart != "", timex.Parse(u.ValidStart), time.Now())
	var validEnd = anyx.If(u.ValidEnd != "", timex.Parse(u.ValidEnd), validStart.AddDate(1, 0, 0))
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
		Password:     password,
		Salt:         salt,
		SessionTime:  3600,
		ValidStart:   validStart,
		ValidEnd:     validEnd,
		CreateUserId: u.CurrUserId,
		UpdateUserId: u.CurrUserId,
		UpdateTime:   time.Now(),
	}
}

func (u *UserSave) UserUpdate() *entity.User {
	user := &entity.User{
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
	if u.Password != "" {
		var salt = randx.UUID()
		user.Salt = salt
		user.Password = encryptx.PasswordSalt(encryptx.MD5(u.Password), salt)
	}
	if u.ValidStart != "" {
		user.ValidStart = timex.Parse(u.ValidStart)
	}
	if u.ValidEnd != "" {
		user.ValidEnd = timex.Parse(u.ValidEnd)
	}
	return user
}
