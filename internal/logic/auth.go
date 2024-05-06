package logic

import (
	"context"
	"errors"
	"time"

	"github.com/go-xuan/quanx/core/ginx"
	"github.com/go-xuan/quanx/utils/encryptx"
	"github.com/go-xuan/quanx/utils/snowflakex"
	log "github.com/sirupsen/logrus"

	"user/internal/dao"
	"user/internal/model"
	"user/internal/model/entity"
)

// 用户登录
func UserLogin(ctx context.Context, param model.Login, loginIp string) (result *model.LoginResult, err error) {
	var user *model.User
	if user, err = dao.GetUserByName(param.Username); err != nil {
		return
	}
	var userAuth *entity.UserAuth
	if userAuth, err = dao.QueryUserAuth(user.Id); err != nil {
		return
	}
	var password string
	if password, err = encryptx.RSA().Decrypt(param.Password); err != nil {
		return
	}
	password = encryptx.MD5(password)
	// 校验密码
	if encryptx.PasswordSalt(password, userAuth.Salt) != userAuth.Password {
		err = errors.New("密码错误")
		return
	}
	var authUser = &ginx.User{
		Id:         user.Id,
		Account:    user.Account,
		Name:       user.Name,
		Phone:      user.Phone,
		Ip:         loginIp,
		ExpireTime: time.Now().Unix() + userAuth.SessionTime,
	}
	var token string
	if token, err = ginx.NewToken(authUser); err != nil {
		log.Error("生成token失败")
		return
	}
	// token存入redis
	ginx.AuthCache.Set(ctx, user.Account, token, time.Duration(userAuth.SessionTime)*time.Second)
	// 记录日志
	var sysLog = entity.Log{
		Id:           snowflakex.New().Int64(),
		Module:       "auth",
		Type:         "login",
		Content:      user.Name + "【" + user.Phone + "账号密码登录】",
		Ip:           loginIp,
		CreateUserId: userAuth.UserId,
	}
	if err = dao.SysLogCreate(sysLog); err != nil {
		log.Error("记录登录日志失败")
		return
	}
	result = &model.LoginResult{User: user, Token: token}
	return
}

// 用户登出
func UserLogout(ctx context.Context, user *ginx.User, ip string) (userId int64, err error) {
	var sysLog = entity.Log{
		Id:           snowflakex.New().Int64(),
		Module:       "auth",
		Type:         "logout",
		Content:      user.Name + "【登出】",
		Ip:           ip,
		CreateUserId: user.Id,
	}
	if err = dao.SysLogCreate(sysLog); err != nil {
		log.Error("记录日志失败")
		return
	}
	// 删除redis上用户token
	ginx.AuthCache.Del(ctx, user.Account)
	return
}
