package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"user/internal/model/entity"

	"github.com/go-xuan/quanx/os/errorx"
)

func NewTokenUser(user *entity.User) *TokenUser {
	return &TokenUser{
		Id:         user.Id,
		Account:    user.Account,
		Name:       user.Name,
		Phone:      user.Phone,
		ExpireTime: time.Now().Unix() + user.SessionTime,
	}
}

// TokenUser 用户token参数
type TokenUser struct {
	Id         int64  `json:"id"`         // 用户ID
	Account    string `json:"account"`    // 用户账号
	Name       string `json:"name"`       // 用户姓名
	Phone      string `json:"phone"`      // 登录手机
	Ip         string `json:"ip"`         // 登录IP
	Domain     string `json:"domain"`     // 域名
	ExpireTime int64  `json:"expireTime"` // 过期时间
}

func (user *TokenUser) Token(key any) (token string, err error) {

	if token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, user).SignedString(key); err != nil {
		return
	}
	return
}

func (user *TokenUser) Username() string {
	return user.Phone
}

func (user *TokenUser) UserId() int64 {
	return user.Id
}

func (user *TokenUser) Valid() error {
	if user.ExpireTime < time.Now().Unix() {
		return errorx.Errorf("token has expired")
	}
	return nil
}
